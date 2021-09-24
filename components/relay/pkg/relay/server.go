package relay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis"

	"github.com/pacedotdev/oto/otohttp"

	"github.com/gorilla/websocket"
	"github.com/thavlik/gcommz/components/relay/pkg/api"
	"github.com/thavlik/gcommz/components/relay/pkg/storage"
	"go.uber.org/zap"
)

type userInfo struct {
	id    string
	l     sync.Mutex
	conns []*websocket.Conn
}

type Server struct {
	usersL  sync.Mutex
	users   map[string]*userInfo
	redis   *redis.Client
	storage storage.Storage
	log     *zap.Logger
}

func NewServer(
	redis *redis.Client,
	log *zap.Logger,
) *Server {
	return &Server{
		users: make(map[string]*userInfo),
		redis: redis,
		log:   log,
	}
}

var errUserNotFound = fmt.Errorf("user not found")

func (s *Server) getUser(userID string) (*userInfo, error) {
	s.usersL.Lock()
	u, ok := s.users[userID]
	s.usersL.Unlock()
	if !ok {
		return nil, errUserNotFound
	}
	return u, nil
}

func (s *Server) Start(
	servicePort int,
	webSocketPort int,
) error {
	webSocketDone := make(chan error, 1)
	go func() {
		webSocketDone <- s.listenWebSocket(webSocketPort)
		close(webSocketDone)
	}()
	serviceDone := make(chan error, 1)
	go func() {
		serviceDone <- s.listenService(servicePort)
		close(serviceDone)
	}()
	redisDone := make(chan error, 1)
	go func() {
		redisDone <- s.subscribeRedis()
		close(redisDone)
	}()
	select {
	case err := <-webSocketDone:
		return fmt.Errorf("listen websocket: %v", err)
	case err := <-serviceDone:
		return fmt.Errorf("listen http: %v", err)
	case err := <-redisDone:
		return fmt.Errorf("redis subscription: %v", err)
	}
}

func (s *Server) listenWebSocket(port int) error {
	mux := http.NewServeMux()
	mux.Handle("/", s.handler())
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func (s *Server) subscribeRedis() error {
	ch := s.redis.Subscribe("_relay").Channel()
	for {
		msg, ok := <-ch
		if !ok {
			return fmt.Errorf("channel closed")
		}
		_ = msg
		// TODO: broadcast message to proper clients
	}
}

func (s *Server) listenService(port int) error {
	otoServer := otohttp.NewServer()
	api.RegisterRelay(otoServer, s)
	return (&http.Server{
		Handler:      otoServer,
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe()
}

func (s *Server) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgraded := false
		statusCode := http.StatusInternalServerError
		if err := func() error {
			s.log.Debug("Handling connection attempt")

			accessToken := r.Header.Get("Authorization")
			if accessToken == "" {
				statusCode = http.StatusBadRequest
				return fmt.Errorf("missing Authorization JWT in header")
			}

			// TODO: validate access token, get userID
			userID := "TODO"
			connLog := s.log.With(zap.String("userID", userID))

			connLog.Info("Upgrading connection")
			defer connLog.Info("Connection terminated")

			upgrader := websocket.Upgrader{}
			c, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				return fmt.Errorf("upgrade: %v", err)
			}
			defer c.Close()
			upgraded = true

			connLog.Debug("Upgraded connection")

			if err := s.ensureUserLoaded(userID); err != nil {
				return fmt.Errorf("ensureUserLoaded: %v", err)
			}

			user := s.addConn(userID, c)
			defer s.removeConn(user, c)

			for {
				mt, message, err := c.ReadMessage()
				if err != nil {
					return fmt.Errorf("ReadMessage: %v", err)
				}
				if mt == websocket.TextMessage {
					data := make(map[string]interface{})
					if err := json.Unmarshal(message, &data); err != nil {
						connLog.Error("unmarshal json message",
							zap.String("message", string(message)), zap.Error(err))
						continue
					}
					// TODO: handle sending messages to channels
				}
			}
		}(); err != nil {
			s.log.Error("WebSocket handler error", zap.Error(err))
			if !upgraded {
				w.WriteHeader(statusCode)
			}
			return
		}
	}
}

func (s *Server) addConn(userID string, c *websocket.Conn) *userInfo {
	defer numWebSocketConnections.Inc()
	s.usersL.Lock()
	user, ok := s.users[userID]
	if ok {
		s.usersL.Unlock() // Unlock immediately
		user.l.Lock()
		user.conns = append(user.conns, c)
		user.l.Unlock()
		return user
	} else {
		// Create a new user
		user := &userInfo{
			id:    userID,
			conns: []*websocket.Conn{c},
		}
		s.users[userID] = user
		s.usersL.Unlock()
		numUsers.Inc()
		return user
	}
}

func (s *Server) removeConn(user *userInfo, c *websocket.Conn) {
	defer numWebSocketConnections.Dec()
	user.l.Lock()
	defer user.l.Unlock()
	conns := user.conns
	connFound := false
	for j, conn := range conns {
		if conn == c {
			conns = append(conns[:j], conns[j+1:]...)
			connFound = true
			break
		}
	}
	if !connFound {
		// Sanity check
		panic(fmt.Sprintf("specific connection not found for user '%s'", user.id))
	}
	if len(conns) == 0 {
		s.usersL.Lock()
		delete(s.users, user.id)
		s.usersL.Unlock()
		numUsers.Dec()
	} else {
		user.conns = conns
	}
}

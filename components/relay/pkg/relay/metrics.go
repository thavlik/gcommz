package relay

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	numWebSocketConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "relay_websocket_connections_total",
		Help: "Total number of active WebSocket connections on this server",
	})

	numUsers = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "relay_users_total",
		Help: "Total number of active users on this server",
	})
)

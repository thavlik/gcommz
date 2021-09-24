package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
	"github.com/thavlik/gcommz/components/relay/pkg/storage"
)

func (s *Server) Join(ctx context.Context, req api.JoinRequest) (*api.JoinResponse, error) {
	p := s.redis.Pipeline()
	p.SAdd(rkUserChannels(req.UserID), req.ChannelID)
	p.SAdd(rkChannelUsers(req.ChannelID), req.UserID)
	if _, err := p.Exec(); err != nil {
		return nil, fmt.Errorf("redis: %v", err)
	}
	if err := s.storage.StoreSubscription(&storage.Subscription{
		UserID:    req.UserID,
		ChannelID: req.ChannelID,
	}); err != nil {
		return nil, fmt.Errorf("storage: %v", err)
	}
	return &api.JoinResponse{}, nil
}

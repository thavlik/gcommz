package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
	"github.com/thavlik/gcommz/components/relay/pkg/storage"
)

func (s *Server) Leave(ctx context.Context, req api.LeaveRequest) (*api.LeaveResponse, error) {
	p := s.redis.Pipeline()
	p.SRem(rkUserChannels(req.UserID), req.ChannelID)
	p.SRem(rkChannelUsers(req.ChannelID), req.UserID)
	if _, err := p.Exec(); err != nil {
		return nil, fmt.Errorf("redis: %v", err)
	}
	if err := s.storage.DeleteSubscription(&storage.Subscription{
		UserID:    req.UserID,
		ChannelID: req.ChannelID,
	}); err != nil {
		return nil, fmt.Errorf("storage: %v", err)
	}
	return &api.LeaveResponse{}, nil
}

package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
)

func (s *Server) Ban(ctx context.Context, req api.BanRequest) (*api.BanResponse, error) {
	p := s.redis.Pipeline()
	p.SAdd(rkChannelBanList(req.ChannelID), req.UserID)
	if _, err := p.Exec(); err != nil {
		return nil, fmt.Errorf("redis: %v", err)
	}
	return &api.BanResponse{}, nil
}

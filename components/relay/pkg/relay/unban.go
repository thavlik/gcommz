package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
)

func (s *Server) Unban(ctx context.Context, req api.UnbanRequest) (*api.UnbanResponse, error) {
	p := s.redis.Pipeline()
	p.SRem(rkChannelBanList(req.ChannelID), req.UserID)
	if _, err := p.Exec(); err != nil {
		return nil, fmt.Errorf("redis: %v", err)
	}
	return &api.UnbanResponse{}, nil
}

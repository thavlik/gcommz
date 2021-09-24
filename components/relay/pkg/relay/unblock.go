package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
)

func (s *Server) Unblock(ctx context.Context, req api.UnblockRequest) (*api.UnblockResponse, error) {
	p := s.redis.Pipeline()
	p.SRem(rkUserBlockList(req.UserID), req.UserToUnblockId)
	if _, err := p.Exec(); err != nil {
		return nil, fmt.Errorf("redis: %v", err)
	}
	return &api.UnblockResponse{}, nil
}

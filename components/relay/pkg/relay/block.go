package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
)

func (s *Server) Block(ctx context.Context, req api.BlockRequest) (*api.BlockResponse, error) {
	p := s.redis.Pipeline()
	p.SAdd(rkUserBlockList(req.UserID), req.UserToBlockId)
	if _, err := p.Exec(); err != nil {
		return nil, fmt.Errorf("redis: %v", err)
	}
	return &api.BlockResponse{}, nil
}

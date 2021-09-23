package relay

import (
	"context"
	"fmt"

	"github.com/thavlik/gcommz/components/relay/pkg/api"
)

func (s *Server) DeleteMessage(ctx context.Context, req api.DeleteMessageRequest) (*api.DeleteMessageResponse, error) {
	return nil, fmt.Errorf("unimplemented")
}

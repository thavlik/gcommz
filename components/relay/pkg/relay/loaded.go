package relay

import "fmt"

func (s *Server) ensureUserLoaded(userID string) error {
	blockIDs, err := s.storage.ListBlocks(userID)
	if err != nil {
		return fmt.Errorf("ListBlocks: %v", err)
	}
	channelIDs, err := s.storage.ListSubscriptions(userID)
	if err != nil {
		return fmt.Errorf("ListSubscriptions: %v", err)
	}
	p := s.redis.Pipeline()
	blockKey := rkUserBlockList(userID)
	for _, blockID := range blockIDs {
		p.SAdd(blockKey, blockID)
	}
	channelKey := rkUserChannels(userID)
	for _, channelID := range channelIDs {
		p.SAdd(channelKey, channelID)
	}
	p.Set(rkUserLoaded(userID), 1, 0)
	if _, err := p.Exec(); err != nil {
		return fmt.Errorf("redis: %v", err)
	}
	return nil
}

package mongodb_storage

import (
	"github.com/thavlik/gcommz/components/relay/pkg/storage"
)

type mongoDBStorage struct{}

func NewMongoDBStorage() storage.Storage {
	return &mongoDBStorage{}
}

func (s *mongoDBStorage) StoreChannel(*storage.Channel) error {
	return nil
}

func (s *mongoDBStorage) RetrieveChannel(id string) (*storage.Channel, error) {
	return nil, nil
}

func (s *mongoDBStorage) StoreSubscription(*storage.Subscription) error {
	return nil
}

func (s *mongoDBStorage) DeleteSubscription(*storage.Subscription) error {
	return nil
}

func (s *mongoDBStorage) ListSubscriptions(userID string) ([]string, error) {
	return nil, nil
}

func (s *mongoDBStorage) StoreBan(*storage.Ban) error {
	return nil
}

func (s *mongoDBStorage) DeleteBan(*storage.Ban) error {
	return nil
}

func (s *mongoDBStorage) ListBans(channelID string) ([]string, error) {
	return nil, nil
}

func (s *mongoDBStorage) StoreBlock(*storage.Block) error {
	return nil
}

func (s *mongoDBStorage) DeleteBlock(*storage.Block) error {
	return nil
}

func (s *mongoDBStorage) ListBlocks(userID string) ([]string, error) {
	return nil, nil
}

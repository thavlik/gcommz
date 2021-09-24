package storage

type Channel struct {
	ID        string
	CreatorID string
	Created   int64
}

type Subscription struct {
	UserID    string
	ChannelID string
}

type Ban struct {
	ChannelID string
	UserID    string
}

type Block struct {
	UserID        string
	UserToBlockID string
}

type Storage interface {
	StoreChannel(*Channel) error
	RetrieveChannel(id string) (*Channel, error)

	StoreSubscription(*Subscription) error
	DeleteSubscription(*Subscription) error
	ListSubscriptions(userID string) ([]string, error)

	StoreBan(*Ban) error
	DeleteBan(*Ban) error
	ListBans(channelID string) ([]string, error)

	StoreBlock(*Block) error
	DeleteBlock(*Block) error
	ListBlocks(userID string) ([]string, error)
}

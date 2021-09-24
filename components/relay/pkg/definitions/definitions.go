package definitions

type Relay interface {
	CreateChannel(CreateChannelRequest) CreateChannelResponse
	DeleteChannel(DeleteChannelRequest) DeleteChannelResponse
	Join(JoinRequest) JoinResponse
	Leave(LeaveRequest) LeaveResponse
	DeleteMessage(DeleteMessageRequest) DeleteMessageResponse
	Kick(KickRequest) KickResponse
	Ban(BanRequest) BanResponse
	Unban(UnbanRequest) UnbanResponse
	Block(BlockRequest) BlockResponse
	Unblock(UnblockRequest) UnblockResponse
}

type CreateChannelRequest struct {
	IsChannel bool `json:"isChannel"`
}

type CreateChannelResponse struct {
	ID string `json:"id"`
}

type DeleteChannelRequest struct {
	ID string `json:"id"`
}

type DeleteChannelResponse struct {
}

type JoinRequest struct {
	UserID    string `json:"userId"`
	ChannelID string `json:"channelId"`
}

type JoinResponse struct {
	ID string `json:"id"` // Subscription ID
}

type LeaveRequest struct {
	UserID    string `json:"userId"`
	ChannelID string `json:"channelId"`
}

type LeaveResponse struct {
}

type DeleteMessageRequest struct {
	ID string `json:"id"`
}

type DeleteMessageResponse struct {
}

type KickRequest struct {
	UserID string `json:"userId"`
}

type KickResponse struct {
}

type BanRequest struct {
	UserID    string `json:"userId"`
	ChannelID string `json:"channelId"`
}

type UnbanResponse struct {
}

type UnbanRequest struct {
	UserID    string `json:"userId"`
	ChannelID string `json:"channelId"`
}

type BanResponse struct {
}

type BlockRequest struct {
	UserID        string `json:"userId"`
	UserToBlockId string `json:"userToBlockId"`
}

type BlockResponse struct {
}

type UnblockRequest struct {
	UserID          string `json:"userId"`
	UserToUnblockId string `json:"userToUnblockId"`
}

type UnblockResponse struct {
}

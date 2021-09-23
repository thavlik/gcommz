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
	UserID string `json:"userId"`
	ChatID string `json:"chatId"`
}

type JoinResponse struct {
}

type LeaveRequest struct {
	UserID string `json:"userId"`
	ChatID string `json:"chatId"`
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
	UserID string `json:"userId"`
	ChatID string `json:"chatId"`
}

type UnbanResponse struct {
}

type UnbanRequest struct {
	UserID string `json:"userId"`
	ChatID string `json:"chatId"`
}

type BanResponse struct {
}

type BlockRequest struct {
	UserID string `json:"userId"`
}

type BlockResponse struct {
}

type UnblockRequest struct {
	UserID string `json:"userId"`
}

type UnblockResponse struct {
}

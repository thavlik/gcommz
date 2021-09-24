package relay

import (
	"fmt"
)

func rkUserLoaded(userID string) string {
	return fmt.Sprintf("u:%s:l", userID)
}

func rkUserChannels(userID string) string {
	return fmt.Sprintf("u:%s:ch", userID)
}

func rkChannelUsers(channelID string) string {
	return fmt.Sprintf("ch:%s:u", channelID)
}

func rkChannelBanList(channelID string) string {
	return fmt.Sprintf("ch:%s:bl", channelID)
}

func rkUserBlockList(userID string) string {
	return fmt.Sprintf("u:%s:bl", userID)
}

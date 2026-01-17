package chat

import (
	"github.com/zelenin/go-tdlib/client"
)

func GetChatHistory(cl client.Client, chatID int64, limit int32) []*client.Message {
	history, err := cl.GetChatHistory(&client.GetChatHistoryRequest{
		ChatId: chatID,
		Limit:  limit,
	})
	if err != nil {
		panic(err)
	}
	return history.Messages
}

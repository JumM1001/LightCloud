package chat

import (
	"github.com/zelenin/go-tdlib/client"
)

func GetMessageById(cl client.Client, chatID int64, messageID int64) *client.Message {
	message, err := cl.GetMessage(&client.GetMessageRequest{
		ChatId:    chatID,
		MessageId: messageID,
	})
	if err != nil {
		panic(err)
	}
	return message
}

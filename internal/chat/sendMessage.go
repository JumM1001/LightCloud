package chat

import (
	"github.com/zelenin/go-tdlib/client"
)

func SendMessage(cl client.Client, chatID int64, messageText string) {
	_, err := cl.SendMessage(&client.SendMessageRequest{
		ChatId: chatID,
		InputMessageContent: &client.InputMessageText{
			Text: &client.FormattedText{
				Text: messageText,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

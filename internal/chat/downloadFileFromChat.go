package chat

import (
	"fmt"
	"strings"

	"github.com/zelenin/go-tdlib/client"
)

func DownloadFileFromChat(cl client.Client, chatID int64, fileName string) {
	chatHistory := GetChatHistory(cl, chatID, 100)
	fmt.Println("Chat history length:", len(chatHistory), "messages", chatHistory)
	for _, message := range chatHistory {
		messageID := message.Id
		message := GetMessageById(cl, chatID, messageID)

		if message.Content != nil {
			if document, ok := message.Content.(*client.MessageDocument); ok {

				fmt.Println("Found document:", document.Caption.Text)

				if strings.Contains(document.Caption.Text, fileName) {
					fmt.Println("File check: ")
					file, err := cl.DownloadFile(&client.DownloadFileRequest{
						FileId:      document.Document.Thumbnail.File.Id,
						Priority:    1,
						Offset:      0,
						Limit:       0,
						Synchronous: true,
					})
					if err != nil {
						panic(err)
					}
					fmt.Println("File downloaded to:", file.Local.Path)
				} else {
					fmt.Println("File name does not match:", document.Caption.Text)
				}

			}
		}
	}
}

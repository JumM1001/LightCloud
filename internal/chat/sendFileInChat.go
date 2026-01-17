package chat

import (
	"strings"

	"github.com/zelenin/go-tdlib/client"
)

func SendFileInChat(cl client.Client, chatID int64, filePath string) {

	// Отправка файла в чат
	_, err := cl.SendMessage(&client.SendMessageRequest{
		ChatId: chatID,
		InputMessageContent: &client.InputMessageDocument{
			Document: &client.InputFileLocal{
				Path: filePath,
			},
			Caption: &client.FormattedText{
				Text: getFileNameFromPath(filePath),
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

// Извлекает имя файла из полного пути
func getFileNameFromPath(filePath string) string {
	filePathElement := make([]string, 0)
	if strings.Contains(filePath, "/") {
		filePathElement = strings.Split(filePath, "/")
	} else if strings.Contains(filePath, `\`) {
		filePathElement = strings.Split(filePath, `\`)
	}
	return filePathElement[len(filePathElement)-1]
}

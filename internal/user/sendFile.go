package user

import (
	"fmt"

	"github.com/zelenin/go-tdlib/client"
	"test.go/internal/chat"
	"test.go/internal/entities"
)

// Просим у пользователя путь к файлу и отправляем его в указанный чат
func SendFile(cl client.Client, myChat entities.ChatData) {
	var filePath string
	fmt.Print("Enter the file path to send: ")
	fmt.Scan(&filePath)

	chat.SendFileInChat(cl, myChat.ChatID, filePath)
}

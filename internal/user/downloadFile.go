package user

import (
	"fmt"

	"github.com/zelenin/go-tdlib/client"
	"test.go/internal/chat"
	"test.go/internal/entities"
)

func DownloadFile(cl client.Client, myChat entities.ChatData) {
	var fileName string
	fmt.Print("Enter the file name: ")
	fmt.Scan(&fileName)
	fmt.Println("Chat history length before download: ", len(chat.GetChatHistory(cl, myChat.ChatID, 100)))
	chat.DownloadFileFromChat(cl, myChat.ChatID, fileName)
	fmt.Println("Chat history length after download: ", len(chat.GetChatHistory(cl, myChat.ChatID, 100)), "\n")
	fmt.Printf("\ndownloadFile: file %s downloaded", fileName)
}

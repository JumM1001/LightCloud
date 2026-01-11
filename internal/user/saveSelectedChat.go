package user

import (
	"encoding/json"
	"fmt"
	"os"

	"test.go/internal/entities"
)

func SaveSelectedChat(chatID int64, chatTitle string, chatType string) {
	chat := entities.ChatData{
		ChatID:    chatID,
		ChatTitle: chatTitle,
		ChatType:  chatType,
	}

	jsonData, err := json.MarshalIndent(chat, "", "  ")
	if err != nil {
		fmt.Println("saveSelectedChat: Error marshaling JSON:", err)
		return
	}
	filepath := "../configs/chat.json"
	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		fmt.Println("saveSelectedChat: Error writing to file:", err)
		return
	}

	fmt.Println("Selected chat saved:", chat.ChatID, chat.ChatTitle, chat.ChatType)
}

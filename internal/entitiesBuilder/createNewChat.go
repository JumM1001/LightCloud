package entitiesBuilder

import (
	"encoding/json"
	"os"

	"github.com/zelenin/go-tdlib/client"
	"test.go/internal/entities"
	"test.go/internal/user"
)

func CreateChatEntity(tdClient client.Client) (entities.ChatData, error) {
	chatID, chatTitle, chatType := user.GetChatsToSelect(tdClient)
	user.SaveSelectedChat(chatID, chatTitle, chatType)
	chatFilePath := "../configs/chat.json"
	configChatFile, err := os.ReadFile(chatFilePath)
	if err != nil {
		return entities.ChatData{}, err
	}

	chat := entities.ChatData{}
	_ = json.Unmarshal(configChatFile, &chat)
	return chat, nil
}

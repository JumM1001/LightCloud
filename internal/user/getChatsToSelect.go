package user

import (
	"fmt"

	"github.com/zelenin/go-tdlib/client"
	"test.go/internal/chat"
)

func GetChatsToSelect(cl client.Client) (int64, string, string) {
	fmt.Println("Выберите чат для хранения файлов")
	chatIDs, chatTitles := chat.GetChats(cl, 20)
	chatsToSelectMap := make(map[int]int64)
	for i, chatID := range chatIDs {
		chatsToSelectMap[i+1] = chatID
		fmt.Printf("%d: %s\n", i+1, chatTitles[i])
	}
	fmt.Print("Введите номер чата: ")
	var choice int
	fmt.Scanln(&choice)

	selectedChatID := chatsToSelectMap[choice]
	fmt.Printf("Выбранный чат: %s, ID чата: %d\n", chatTitles[choice-1], selectedChatID)
	chatType, err := cl.GetChat(&client.GetChatRequest{
		ChatId: selectedChatID,
	})
	if err != nil {
		return 0, "", ""
	}

	return selectedChatID, chatTitles[choice-1], chatType.Type.ChatTypeType()
}

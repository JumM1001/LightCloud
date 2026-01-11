package chat

import (
	"log"

	"github.com/zelenin/go-tdlib/client"
)

func GetChats(cl client.Client, limit int) ([]int64, []string) {
	chatListType := client.ChatListMain{}
	chats, err := cl.GetChats(&client.GetChatsRequest{
		ChatList: &chatListType,
		Limit:    int32(limit),
	})

	if err != nil {
		panic(err)
	}
	chatIDs := make([]int64, 0)
	chatTitles := make([]string, 0)
	for _, chatID := range chats.ChatIds {
		chat, err := cl.GetChat(&client.GetChatRequest{
			ChatId: chatID,
		})
		if err != nil {
			log.Println("GetChat error:", err)
			continue
		}
		chatIDs = append(chatIDs, chatID)
		chatTitles = append(chatTitles, chat.Title)
	}
	return chatIDs, chatTitles
}

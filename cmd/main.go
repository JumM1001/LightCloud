package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	myclient "test.go/internal/client"
	"test.go/internal/entities"
	"test.go/internal/user"

	"github.com/zelenin/go-tdlib/client"
)

func main() {
	tdClient := myclient.CreateNewClient()
	me, err := tdClient.GetMe()
	if err != nil {
		log.Fatal("main: Error getting user info:", err)
	}
	log.Println("Logged as:", me.FirstName, me.LastName)

	err = runapp(tdClient)
	if err != nil {
		log.Fatal("runapp: Error running app:", err)
	}
	select {}

}

func runapp(tdClient client.Client) error {

	chatID, chatTitle, chatType := user.GetChatsToSelect(tdClient)
	user.SaveSelectedChat(chatID, chatTitle, chatType)
	chatFilePath := "../configs/chat.json"
	configChatFile, err := os.ReadFile(chatFilePath)
	if err != nil {
		return err
	}

	chat := entities.ChatData{}
	_ = json.Unmarshal(configChatFile, &chat)
	fmt.Println(chat)
	return nil
}

package main

import (
	"fmt"
	"log"

	"test.go/internal/chat"
	myclient "test.go/internal/client"
	"test.go/internal/entitiesBuilder"
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

	go runapp(tdClient)
	if err != nil {
		log.Fatal("runapp: Error running app:", err)
	}

	select {}
}

func runapp(tdClient client.Client) error {

	myChat, err := entitiesBuilder.CreateChatEntity(tdClient)
	if err != nil {
		return err
	}

	user.DownloadFile(tdClient, myChat)
	fmt.Println("Chat history length:", len(chat.GetChatHistory(tdClient, myChat.ChatID, 100)))
	return nil
}

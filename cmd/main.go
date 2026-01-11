package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	godotenv "github.com/joho/godotenv"
	"github.com/zelenin/go-tdlib/client"
	"test.go/internal/entities"
	api "test.go/internal/entities"
	"test.go/internal/user"
)

func main() {
	// Load .env file
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	api := api.APIData{}

	id, err := strconv.Atoi(os.Getenv("api_id"))
	if err != nil {
		log.Fatal("Error init api id")
	}

	// Set API ID and API HASH
	api.API_ID = int32(id)
	api.API_HASH = os.Getenv("api_hash")

	// TDLib parameters setup
	tdlibParameters := &client.SetTdlibParametersRequest{
		DatabaseDirectory:   filepath.Join("tdlib", "db"),
		FilesDirectory:      filepath.Join("tdlib", "files"),
		UseTestDc:           false,
		ApiId:               api.API_ID,   // ← ТВОЙ api_id
		ApiHash:             api.API_HASH, // ← ТВОЙ api_hash
		UseFileDatabase:     false,
		UseMessageDatabase:  false,
		UseChatInfoDatabase: false,
		SystemLanguageCode:  "en",
		DeviceModel:         "Windows PC",
		SystemVersion:       "Windows 11",
		ApplicationVersion:  "1.0",
	}

	// Create TDLib client and authorizer

	authorizer := client.ClientAuthorizer(tdlibParameters)
	go client.CliInteractor(authorizer)
	client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})

	tdClient, err := client.NewClient(authorizer)
	if err != nil {
		log.Fatal(err)
	}

	me, err := tdClient.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Logged as:", me.FirstName, me.LastName)

	_ = runapp(*tdClient)
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

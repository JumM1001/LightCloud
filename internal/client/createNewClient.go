package myclient

import (
	"log"
	"strconv"

	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/zelenin/go-tdlib/client"
	api "test.go/internal/entities"
)

func CreateNewClient() client.Client {
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
		ApiId:               api.API_ID,
		ApiHash:             api.API_HASH,
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
	return *tdClient
}

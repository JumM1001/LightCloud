package entities

type APIData struct {
	API_HASH string `envconfig:"api_hash"`
	API_ID   int32  `envconfig:"api_id"`
}

type ChatData struct {
	ChatID    int64  `json:"chat_id"`
	ChatTitle string `json:"chat_title"`
	ChatType  string `json:"chat_type"`
}

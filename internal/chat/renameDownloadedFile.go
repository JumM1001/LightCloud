package chat

import (
	"os"

	"github.com/zelenin/go-tdlib/client"
)

func RenameDownloadedFile(cl client.Client, fileName, newFileName string) {
	filePath := "../tdlib/db/thumbnails/" + fileName
	newFilePath := "../tdlib/db/thumbnails/" + newFileName

	err := os.Rename(filePath, newFilePath)
	if err != nil {
		panic(err)
	}

}

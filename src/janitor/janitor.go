package janitor

import (
	"log"
	"os"
)

func DeleteFile(filePath string) bool {
	if err := os.Remove(filePath); err != nil {
		log.Println("Failed to delete file: ", filePath, "\nbecause:", err)
		return false
	}
	return true
}

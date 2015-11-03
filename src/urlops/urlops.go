package urlops

import (
	"io"
	"log"
	"net/http"
	"os"
)

type DownloadCallBack func(error, bool, string)

func DownloadImage(downloadURL, destImagePath string, callback func(error, bool, string)) {
		//check if file already exist, if so just return true
	if _, err := os.Stat(destImagePath); os.IsExist(err) {
		log.Println("Image already downloaded", destImagePath)
		callback(nil, true, destImagePath)
		return
	}
	resp, err := http.Get(downloadURL)
	if err != nil {
		log.Println("Error downloading from URL: ", downloadURL, "\nbecause: ", err)
		callback(err, false, "")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Error closing download connection for url: ", downloadURL, "\nbecause: ", err)
		}
	}()
	file, err := os.Create(destImagePath)
	if err != nil {
		log.Println("Failed to create write file for download URL: ", downloadURL, "\nbecause: ", err)
		callback(err, false, "")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Failed to close opened file: ", destImagePath, "\nbecuase: ", err)
		}
	}()
	if bytesCopied, err := io.Copy(file, resp.Body); err != nil {
		log.Println("Failed to copy data for download url: ", downloadURL, " to ", destImagePath, "\nbecause: ", err)
		callback(err, false, "")
	} else {
		log.Println("Download successful from ", downloadURL, "to ", destImagePath, " ", bytesCopied, " bytes written.")
		callback(nil, true, destImagePath)
	}
}

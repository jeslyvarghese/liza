package urlops

import (
	"io"
	"log"
	"net/http"
	"os"
)

type DownloadCallBack func(err error, isSuccess bool, destImagePath string)

func DownloadImage(downloadURL, destImagePath string, callback DownloadCallBack) {
	resp, err := http.Get(downloadURL)
	if err != nil {
		log.Fatalln("Error downloading from URL: ", downloadURL, "\nbecause: ", err)
		callback(err, false)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalln("Error closing download connection for url: ", downloadURL, "\nbecause: ", err)
		}
	}()
	file, err := os.Create(destImagePath)
	if err != nil {
		log.Fatalln("Failed to create write file for download URL: ", downloadURL, "\nbecause: ", destImagePath)
		callback(err, false)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln("Failed to close opened file: ", destImagePath, "\nbecuase: ", err)
		}
	}()
	if bytesCopied, err := io.Copy(file, resp.Body); err != nil {
		log.Fatalln("Failed to copy data for download url: ", downloadURL, " to ", destImagePath, "\nbecause: ", err)
		callback(err, false, nil)
	} else {
		log.Println("Download successful from ", downloadURL, "to ", destImagePath, " ", bytesCopied, " bytes written.")
		callback(nil, true, destImagePath)
	}
}

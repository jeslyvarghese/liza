package engine

import (
	"github.com/jeslyvarghese/liza/src/janitor"
	"github.com/jeslyvarghese/liza/src/rackspace"
	"github.com/jeslyvarghese/liza/src/redis"
	"github.com/jeslyvarghese/liza/src/urlops"
	"github.com/jeslyvarghese/liza/src/vips"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func CheckHasImage(imageURL string) (string bool) {
	return redis.GetURL(imageURL)
}

func DownloadImage(imageURL string, callback urlops.DownloadCallBack) bool {
	u, err := url.Parse(imageURL)
	if err != nil {
		log.Fatalln("Unable to parse url: ", imageURL, "\nerror: ", err)
	}
	host := u.Host
	path := u.Path
	dirPath := "/tmp" + path[0:len(path)-len(filepath.Ext(path))]
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		log.Fatalln("Unable to create directories: ", dirPath, "\ncause:", err)
	}
	destImagePath := dirPath + filepath.Base(path)
	urlops.DownloadImage(imageURL, destImagePath, callback)
}

func ResizeImage(imagePath string, width, height int64) (string bool) {
	filePath = imagePath[0 : len(imagePath)-len(filepath.Base(imagePath))]
	if err := os.MkdirAll(filepath, 0777); err != nil {
		log.Fatalln("Could not create directories: ", filepath, "\nerror: ", err)
	}
	dstImagePath := filePath + fmt.Sprintf("%dx%d", width, height)
	success := vips.ResizeImage(imagePath, dstImagePath, width, height)
	if success {
		return dstImagePath, success
	} else {
		return "", success
	}
}

func UploadImage(imagePath string) {
	fileName := strings.Replace(imagePath[len("/tmp/"):len(imagePath)], "/", "", -1)
	containerName := "merlin"
	go func() {
		if isSuccess := rackspace.UploadImage(imagePath, fileName, containerName); isSuccess != true {
			rackspace.UploadCallback(nil, false)
		} else {
			rackspace.UploadCallback(nil, true)
		}

	}()
}

func AddImage(requestURL, imageURL string) bool {
	return redis.AddURL(requestURL, imageURL)
}

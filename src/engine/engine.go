package engine

import (
	"github.com/jeslyvarghese/liza/src/rackspace"
	"github.com/jeslyvarghese/liza/src/redis"
	"github.com/jeslyvarghese/liza/src/urlops"
	"github.com/jeslyvarghese/liza/src/vips"
	"math/rand"
	"time"
	"github.com/jeslyvarghese/liza/src/janitor"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"fmt"
	"math"
)

func CheckHasImage(imageURL string) (string, bool) {
	return redis.GetURL(imageURL)
}

func DownloadImage(imageURL string, callback func(error, bool, string)) bool {
	u, err := url.Parse(imageURL)
	if err != nil {
		log.Println("Unable to parse url: ", imageURL, "\nerror: ", err)
		return false
	}
	// host := u.Host
	path := u.Path
	// hl := int(math.Min(5., float64(len(host))))
	dirPath := "/tmp/" + RandomString(5) + path[0:len(path)-len(filepath.Base(path))]
	log.Println("Dir path for downloads: ", dirPath)
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		log.Println("Unable to create directories: ", dirPath, "\ncause:", err)
		return false
	}
	l := int(math.Min(7, float64(len(filepath.Base(path)))))
	destImagePath := dirPath + filepath.Base(path)[0:l] + filepath.Base(path)[len(filepath.Base(path))-l:len(filepath.Base(path))] + filepath.Ext(path)
	log.Println("DownloadImage path assigned:", destImagePath)
	log.Println("Path has length:", len(destImagePath))
	urlops.DownloadImage(imageURL, destImagePath, callback)
	return true
}

func ResizeImage(imagePath, imageURL string) (string, bool) {
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Println("Could not find downloaded image", imagePath)
		return "", false
	}
	parsedURL, _ := url.Parse(imageURL)
	width, _ := strconv.Atoi(parsedURL.Query().Get("width"))
	height, _ := strconv.Atoi(parsedURL.Query().Get("height"))
	log.Println("URL:", parsedURL, "width:", width, "height:", height)
	imageFilePath := imagePath[0 : len(imagePath)-len(filepath.Base(imagePath))]
	if err := os.MkdirAll(imageFilePath, 0777); err != nil {
		log.Fatalln("Could not create directories: ", imageFilePath, "\nerror: ", err)
	}
	dstImagePath := imageFilePath + fmt.Sprintf("%dx%d", width, height) + filepath.Ext(imagePath)
	log.Println("Writing to path:",dstImagePath)
	success := vips.ResizeImage(imagePath, dstImagePath, width, height)
	if success {
		janitor.DeleteFile(imagePath)
		return dstImagePath, success
	} else {
		return "", success
	}
}

func UploadImage(imagePath, imageURL string, callback rackspace.UploadCallback) {
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Println("Could not find resizedImage image", imagePath)
		callback(nil, false, "")
		return
	}
	u, _ := url.Parse(imageURL)
	fileName := RandomString(5)+"/"+u.Path[0:len(u.Path)-len(filepath.Ext(u.Path))]+"/"+filepath.Base(imagePath)
	log.Println("Rackspace filepath:", fileName)
	cdnURL := "https://03188cc7126169c646ce-4ec321cd871e45e74b11708f248e0363.ssl.cf1.rackcdn.com/"
	containerName := "merlin"
	go func() {
		if isSuccess := rackspace.UploadImage(imagePath, fileName, containerName); isSuccess != true {
			callback(nil, false, "")
		} else {
			callback(nil, true, cdnURL+fileName)
			janitor.DeleteFile(imagePath)
		}
	}()
}

func AddImage(requestURL, imageURL string) bool {
	return redis.AddURL(requestURL, imageURL)
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

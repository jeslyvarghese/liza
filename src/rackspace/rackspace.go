package rackspace

import (
	"github.com/ncw/swift"
	"io/ioutil"
	"log"
)

type UploadCallback func(err error, isSuccess bool, uploadImageURL string)

func UploadImage(fileURL, fileName, containerName string) bool {
	cDallas := swift.Connection{
		UserName: "teliportme",
		ApiKey:   "",
		AuthUrl:  "https://auth.api.rackspacecloud.com/v1.0",
		Region:   "DFW",
	}
	if err := cDallas.Authenticate(); err != nil {
		log.Println("Couldn't authenticate with rackspace Dallas")
		return false
	}
	var rspaceHeader swift.Headers
	fDallas, err := cDallas.ObjectCreate(containerName, fileName, false, "", "image/jpg", rspaceHeader)
	if err != nil {
		log.Println("Failed to create rackspace object ", err)
		return false
	}
	defer fDallas.Close()
	d, err := ioutil.ReadFile(fileURL)
	if err != nil {
		log.Println("Failed to read source file ", err)
		return false
	}
	log.Println("uploading path:", fileURL)
	if _, err := fDallas.Write(d); err != nil {
		log.Println("Failed to write file to rackspace Dallas server", err)
		return false
	}
	return true
}

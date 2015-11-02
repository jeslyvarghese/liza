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
		ApiKey:   "f54540fe1c443d41a33c6f301123795b",
		AuthUrl:  "https://auth.api.rackspacecloud.com/v1.0",
		Region:   "DFW",
	}
	cHKNG := swift.Connection{
		UserName: "teliportme",
		ApiKey:   "f54540fe1c443d41a33c6f301123795b",
		AuthUrl:  "https://auth.api.rackspacecloud.com/v1.0",
		Region:   "HKG",
	}
	if err := cDallas.Authenticate(); err != nil {
		log.Println("Couldn't authenticate with rackspace Dallas")
		return false
	}
	if err := cHKNG.Authenticate(); err != nil {
		log.Println("Couldn't authenticate with rackspace Hong Kong")
		return false
	}
	var rspaceHeader swift.Headers
	fDallas, err := cDallas.ObjectCreate(containerName, fileName, false, "", "image/jpg", rspaceHeader)
	if err != nil {
		log.Println("Failed to create rackspace object ", err)
		return false
	}
	defer fDallas.Close()
	fHKNG, err := cHKNG.ObjectCreate(containerName, fileName, false, "", "image/jpg", rspaceHeader)
	if err != nil {
		log.Println("Failed to create rackspace object ", err)
		return false
	}
	defer fHKNG.Close()
	d, err := ioutil.ReadFile(fileURL)
	if err != nil {
		log.Println("Failed to read source file ", err)
		return false
	}
	if _, err := fDallas.Write(d); err != nil {
		log.Println("Failed to write file to rackspace Dallas server", err)
		return false
	}
	// if _, err := fHKNG.Write(d); err != nil {
	// 	log.Println("Failed to write file to rackspace HongKong server", err)
	// 	return false
	// }
	return true
}

package vips

import (
	"github.com/daddye/vips"
	"io/ioutil"
	"log"
	"os"
)

func ResizeImage(srcImagePath, dstImagePath string, width, height int) bool {
	//check if file already exist, if so just return true
	if _, err := os.Stat(dstImagePath); os.IsExist(err) {
		return true
	}
	options := vips.Options{
		Width:   width,
		Height:  height,
		Crop:    false,
		Enlarge: false,
		Extend:  vips.EXTEND_WHITE,
		Interpolator: vips.NOHALO,
		Gravity:      vips.CENTRE,
		Quality:      90,
	}
	f, _ := os.Open(srcImagePath)
	inBuf, _ := ioutil.ReadAll(f)
	buf, err := vips.Resize(inBuf, options)
	if err != nil {
		log.Println("VIPS Error:", err)
		return false
	}
	of, err := os.Create(dstImagePath)
	if err != nil {
		log.Println("Failed to open file to write to:", dstImagePath, " because: ", err)
		return false
	}
	defer func() {
		if err := of.Close(); err != nil {
			log.Println("Failed to close file: ", dstImagePath, " because: ", err)
		}
	}()
	if _, err := of.Write(buf); err != nil {
		log.Println("Failed to write to file: ", dstImagePath, "because: ", err)
		return false
	}
	return true
}

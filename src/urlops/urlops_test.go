package urlops

import (
	"os"
	"testing"
)

const (
	DownloadURL string = "http://de59658a8604eeb307ec-0d35c4f15040cfced3f623ba9067988e.r54.cf1.rackcdn.com/photos/2500125/4de27ff0a2093e8ac1a3068d64c7b262.jpg"
	DestPath    string = "/tmp/traxex.jpg"
)

func TestDownloadImage(t *testing.T) {
	DownloadImage(DownloadURL, DestPath, func(err error, isSuccess bool) {
		if isSuccess == false {
			t.Fail()
		} else if _, err := os.Stat(DestPath); os.IsNotExist(err) {
			t.Fail()
		}
	})
}

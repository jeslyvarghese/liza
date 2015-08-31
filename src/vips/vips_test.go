package vips

import (
	"os"
	"testing"
)

func TestResizeImage(t *testing.T) {
	testDir := "/Users/jeslyvarghese/go/src/github.com/jesly.varghese/liza/" //runPath[0 : len(runPath)-len("/src/vips/vips_test.go")]
	srcImagePath := testDir + "/tests/test_resize.jpg"
	dstImagePath := testDir + "/tests/test_resize_300x400.jpg"
	if ResizeImage(srcImagePath, dstImagePath, 400, 300) != true {
		t.Fail()
	} else if _, err := os.Stat(dstImagePath); os.IsNotExist(err) {
		t.Fail()
	}
}

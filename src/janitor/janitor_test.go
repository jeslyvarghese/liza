package janitor

import (
	"os"
	"testing"
)

const (
	TestPath string = "/tmp/test1.dat"
)

func TestDeleteFile(t *testing.T) {
	if file, err := os.Create(TestPath); err != nil {
		defer file.Close()
		t.Fail()
	} else {
		defer file.Close()
	}
	if DeleteFile(TestPath) == false {
		t.Fail()
	} else if _, err := os.Stat(TestPath); !os.IsNotExist(err) {
		t.Fail()
	}
}

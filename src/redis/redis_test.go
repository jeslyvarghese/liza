package redis

import (
	"strings"
	"testing"
)

const (
	JobURL   string = "http://de59658a8604eeb307ec-0d35c4f15040cfced3f623ba9067988e.r54.cf1.rackcdn.com/photos/2500125/4de27ff0a2093e8ac1a3068d64c7b262.jpg"
	LocalURL string = "/tmp/de59658a8604eeb307ec-0d35c4f15040cfced3f623ba9067988e.r54.cf1.rackcdn.comphotos/2500125/300x400/4de27ff0a2093e8ac1a3068d64c7b262.jpg"
)

func TestQueueJob(t *testing.T) {
	if false == QueueJob(JobURL) {
		t.Fail()
	}
}

func TestGetJob(t *testing.T) {
	if jobURL, success := GetJob(); success == false {
		t.Fail()
	} else {
		if strings.EqualFold(jobURL, JobURL) == false {
			t.Fail()
		}
	}
}

func TestAddURL(t *testing.T) {
	if false == AddURL(JobURL, LocalURL) {
		t.Fail()
	}
}

func TestGetURL(t *testing.T) {
	if localURL, success := GetURL(JobURL); success == false {
		t.Fail()
	} else if strings.EqualFold(localURL, LocalURL) == false {
		t.Fail()
	}
}

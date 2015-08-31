package collector

import (
	"github.com/jeslyvarghese/liza/src/worker/request"
	"log"
)

var WorkQueue = make(chan Request, 100)

func Collector(url string, width, height int64) {
	request := Request{URL: url, Width: width, Height: height}
	WorkQueue <- request
}

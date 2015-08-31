package collector

import (
	"github.com/jeslyvarghese/liza/src/toiler/request"
)

var WorkQueue = make(chan request.Request, 100)

func Collector(url string, width, height int64) {
	req := request.Request{URL: url, Width: width, Height: height}
	WorkQueue <- req
}

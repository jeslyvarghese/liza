package dispatcher

import (
	"github.com/jeslyvarghese/liza/src/toiler/worker"
	"github.com/jeslyvarghese/liza/src/toiler/request"
)

var WorkerQueue chan chan request.Request

func StartDispatcher(nworkers int) {
	WorkerQueue = make(chan chan request.Request, nworkers)
	for
}

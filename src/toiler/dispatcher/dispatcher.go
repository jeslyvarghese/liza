package dispatcher

import (
	"github.com/jeslyvarghese/liza/src/worker/worker"
	"github.com/jeslyvarghese/liza/src/worker/request"
)

var WorkerQueue chan chan request.Request

func StartDispatcher(nworkers int) {
	WorkerQueue = make(chan chan request.Request, nworkers)
	for
}

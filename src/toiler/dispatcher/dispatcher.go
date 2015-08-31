package dispatcher

import (
	"github.com/jeslyvarghese/liza/src/toiler/request"
	"github.com/jeslyvarghese/liza/src/toiler/worker"
)

var WorkerQueue chan chan request.Request

func StartDispatcher(nworkers int) {
	WorkerQueue = make(chan chan request.Request, nworkers)
	for i := 0; i < nworkers; i++ {
		workers := worker.New()
	}
}

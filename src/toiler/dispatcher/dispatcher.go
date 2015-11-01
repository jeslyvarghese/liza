package dispatcher

import (
	"github.com/jeslyvarghese/liza/src/toiler/collector"
	"github.com/jeslyvarghese/liza/src/toiler/request"
	"github.com/jeslyvarghese/liza/src/toiler/worker"
)

var WorkerQueue chan chan request.Request

func StartDispatcher(nworkers int, p worker.WorkerProcedure) {
	WorkerQueue = make(chan chan request.Request, nworkers)
	for i := 0; i < nworkers; i++ {
		worker := worker.New(i, WorkerQueue)
		worker.Start(p)
	}
	go func() {
		for {
			select {
			case work := <-collector.WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}

package worker

import (
	"github.com/jeslyvarghese/liza/src/worker/request"
	"github.com/jeslyvarghese/liza/src/worker/worker"
)

func NewWorker(id int, workerQueue chan chan request.Request) Worker {

}

type Worker struct {
	ID          int
	Work        chan request.Request
	WorkerQueue chan chan request.Request
	QuitChan    chan bool
}

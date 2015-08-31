package workers

import (
	"github.com/jeslyvarghese/liza/src/worker/request"
)

type WorkerProcedure func(string, int64, int64)

func NewWorker(id int, workerQueue chan chan request.Request) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan request.Request),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
	return worker
}

type Worker struct {
	ID          int
	Work        chan request.Request
	WorkerQueue chan chan request.Request
	QuitChan    chan bool
}

func (w Worker) Start(p WorkerProcedure) {
	go func() {
		for {
			w.WorkerQueue <- w.Work
			select {
			case work := <-w.Work:
				p(work.URL, work.Width, work.Height)
			case <-w.QuitChan:

				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

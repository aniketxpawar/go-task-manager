package manager

import (
	"fmt"
	"time"
)

type Worker struct {
	ID       int
	TaskChan chan Task
	quit     chan bool
}

func NewWorker(id int) *Worker {
	return &Worker{
		ID:       id,
		TaskChan: make(chan Task),
		quit:     make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case task := <-w.TaskChan:
				fmt.Printf("Worker %d processing Task: %s\n",w.ID,task.Payload)
				time.Sleep(2 * time.Second)
				fmt.Printf("Worker %d completed Task: %s\n",w.ID,task.Payload)
			case <- w.quit:
				fmt.Printf("Worker %d Stopping...\n",w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop(){
	w.quit <- true
}
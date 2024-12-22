package manager

type Dispatcher struct {
	Workers   []*Worker
	TaskQueue *TaskQueue
}

func NewDispatcher(workerCount int) *Dispatcher {
	workers := make([]*Worker, workerCount)
	for i := 0; i < workerCount; i++ {
		workers[i] = NewWorker(i + 1)
	}

	return &Dispatcher{
		Workers:   workers,
		TaskQueue: NewTaskQueue(),
	}
}

func (d *Dispatcher) Start() {
	for _, worker := range d.Workers {
		worker.Start()
	}
	go func() {
		for {
			task := d.TaskQueue.Dequeue()
			if task != nil {
				for _, worker := range d.Workers {
					select {
					case worker.TaskChan <- *task:
						break
					default:
						continue
					}
				}
			}
		}

	}()
}

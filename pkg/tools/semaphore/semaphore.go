package semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type DefaultSemaphore struct {
	Semaphore chan struct{}
}

func NewDefaultSemaphore(maxParallelRequests int) *DefaultSemaphore {
	return &DefaultSemaphore{
		Semaphore: make(chan struct{}, maxParallelRequests),
	}
}

func (s *DefaultSemaphore) Acquire() {
	s.Semaphore <- struct{}{}
}

func (s *DefaultSemaphore) Release() {
	<-s.Semaphore
}

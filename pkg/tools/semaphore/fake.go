package semaphore

type FakeSemaphore struct {
	AcquireCalled bool
	ReleaseCalled bool
}

func NewFakeSemaphore() *FakeSemaphore {
	return &FakeSemaphore{}
}

func (s *FakeSemaphore) Acquire() {
	s.AcquireCalled = true
}

func (s *FakeSemaphore) Release() {
	s.ReleaseCalled = true
}

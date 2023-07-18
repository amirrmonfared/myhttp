package hasher

import "io"

type FakeHasher struct {
	HashValue string
	HashError error
}

func (h *FakeHasher) Hash(body io.Reader) (string, error) {
	return h.HashValue, h.HashError
}

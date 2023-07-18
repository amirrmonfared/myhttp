package hasher

import (
	"crypto/md5"
	"fmt"
	"io"
)

type Hasher interface {
	Hash(io.Reader) (string, error)
}

type URLHasher struct{}

func NewURLHasher() *URLHasher {
	return &URLHasher{}
}

func (h *URLHasher) Hash(body io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, body); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

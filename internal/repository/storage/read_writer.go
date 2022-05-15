package storage

import (
	"io/ioutil"

	"github.com/yeldisbayev/crypto-cli/internal/repository"
)

type readWriter struct {
}

func NewReadWriter() repository.ReadWriter {
	return readWriter{}

}

func (rw readWriter) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)

}

func (rw readWriter) Write(path string, payload []byte) error {
	return ioutil.WriteFile(path, payload, 0777)

}

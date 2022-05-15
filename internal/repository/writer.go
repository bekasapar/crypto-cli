package repository

type Writer interface {
	Write(path string, payload []byte) error
}

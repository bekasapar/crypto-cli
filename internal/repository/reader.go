package repository

type Reader interface {
	Read(path string) ([]byte, error)
}

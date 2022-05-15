package usecase

type CryptoUsecase interface {
	Encrypt(path string) error
	Decrypt(path string) error
}

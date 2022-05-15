package delivery

import "github.com/yeldisbayev/crypto-cli/internal/usecase"

type handler struct {
	cryptoUsecase usecase.CryptoUsecase
}

func NewHandler(cryptoUsecase usecase.CryptoUsecase) handler {
	return handler{
		cryptoUsecase,
	}

}

func (h handler) Encrypt(path string) error {
	return h.cryptoUsecase.Encrypt(path)

}

func (h handler) Decrypt(path string) error {
	return h.cryptoUsecase.Decrypt(path)

}

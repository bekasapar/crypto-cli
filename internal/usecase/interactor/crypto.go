package interactor

import (
	"strings"

	"github.com/yeldisbayev/crypto-cli/internal/repository"
	"github.com/yeldisbayev/crypto-cli/internal/usecase"
	"github.com/yeldisbayev/crypto-cli/pkg/crypto"
)

type cryptoInteractor struct {
	extension string
	crypto    crypto.Crypto
	storage   repository.ReadWriter
}

func NewCryptoUsecase(
	extension string,
	crypto crypto.Crypto,
	storage repository.ReadWriter,
) usecase.CryptoUsecase {
	return cryptoInteractor{
		extension,
		crypto,
		storage,
	}

}

func (c cryptoInteractor) Encrypt(path string) error {
	file, err := c.storage.Read(path)

	if err != nil {
		return err
	}

	ciphertext, err := c.crypto.Encrypt(file)

	if err != nil {
		return err
	}

	path = path + c.extension

	return c.storage.Write(path, ciphertext)

}

func (c cryptoInteractor) Decrypt(path string) error {
	file, err := c.storage.Read(path)

	if err != nil {
		return err
	}

	plaintext, err := c.crypto.Decrypt(file)

	if err != nil {
		return err
	}

	path = strings.ReplaceAll(path, c.extension, "")

	return c.storage.Write(path, plaintext)

}

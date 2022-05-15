package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type Crypto interface {
	Encrypt(payload []byte) ([]byte, error)
	Decrypt(payload []byte) ([]byte, error)
}

type crypto struct {
	secret []byte
}

func NewCrypto(secret []byte) Crypto {
	return crypto{
		secret,
	}

}

func (c crypto) Encrypt(payload []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.secret)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, payload, nil)

	return ciphertext, nil

}

func (c crypto) Decrypt(payload []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.secret)
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Panic(err)
	}

	nonce := payload[:gcm.NonceSize()]
	payload = payload[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, payload, nil)

	if err != nil {
		return nil, err
	}

	return plaintext, nil

}

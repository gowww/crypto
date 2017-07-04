// Package crypto provides encryption and hashing utilities.
package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
)

// An Encrypter provides encryption and hashing methods.
type Encrypter interface {
	Encrypt([]byte) ([]byte, error)
	EncryptBase64([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	DecryptBase64([]byte) ([]byte, error)
	HashHS256([]byte) ([]byte, error)
}

// NewEncrypter returns a new AES-128 encrypter for secret key.
// The key must be 32 bytes long.
func NewEncrypter(key []byte) (Encrypter, error) {
	if len(key) != 32 {
		panic("encrypt: encrypter key must be 32 bytes long")
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return &encrypter{
		key:  key,
		aead: gcm,
	}, nil
}

type encrypter struct {
	key  []byte
	aead cipher.AEAD
}

// Encrypt encrypts a plain text.
func (e *encrypter) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, e.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return e.aead.Seal(nonce, nonce, plaintext, nil), nil
}

// EncryptBase64 encrypts a plain text and encodes it in base64.
func (e *encrypter) EncryptBase64(plaintext []byte) ([]byte, error) {
	b, err := e.Encrypt(plaintext)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	enc := base64.NewEncoder(base64.StdEncoding, buf)
	_, err = enc.Write(b)
	if err != nil {
		return nil, err
	}
	enc.Close()
	return buf.Bytes(), nil
}

// Decrypt decrypts a ciphered text.
func (e *encrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := e.aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("encrypt: ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return e.aead.Open(nil, nonce, ciphertext, nil)
}

// DecryptBase64 decodes a base64 ciphered text and decrypts it.
func (e *encrypter) DecryptBase64(ciphertext []byte) ([]byte, error) {
	b, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, bytes.NewReader(ciphertext)))
	if err != nil {
		return nil, err
	}
	return e.Decrypt(b)
}

// HashHS256 returns the HMAC with SHA256 of a plain text.
func (e *encrypter) HashHS256(plaintext []byte) ([]byte, error) {
	h := hmac.New(sha256.New, e.key)
	_, err := h.Write(plaintext)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

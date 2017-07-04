package crypto

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	testEncrypterKey       = []byte("secret-key-secret-key-secret-key")
	testEncrypterPlaintext = []byte("foobar")
)

func TestEncrypter(t *testing.T) {
	e, err := NewEncrypter(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	b, err := e.Encrypt([]byte(testEncrypterPlaintext))
	if err != nil {
		panic(err)
	}
	b, err = e.Decrypt(b)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(testEncrypterPlaintext, b) {
		t.Errorf("encode/decode: want %s, got %s", testEncrypterPlaintext, b)
	}
}

func TestEncrypterBase64(t *testing.T) {
	e, err := NewEncrypter(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	b, err := e.EncryptBase64(testEncrypterPlaintext)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	b, err = e.DecryptBase64(b)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(testEncrypterPlaintext, b) {
		t.Errorf("encode/decode base64: want %s, got %s", testEncrypterPlaintext, b)
	}
}

func TestHashHS256(t *testing.T) {
	e, err := NewEncrypter(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	_, err = e.HashHS256(testEncrypterPlaintext)
	if err != nil {
		panic(err)
	}
}

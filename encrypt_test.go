package encrypt

import (
	"testing"
)

var (
	testEncrypterKey       = "secret-key-secret-key-secret-key"
	testEncrypterPlaintext = "foobar"
)

func TestEncrypter(t *testing.T) {
	e, err := New(testEncrypterKey)
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
	if testEncrypterPlaintext != string(b) {
		t.Errorf("encode/decode: want %s, got %s", testEncrypterPlaintext, b)
	}
}

func TestEncrypterString(t *testing.T) {
	e, err := New(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	s, err := e.EncryptString(testEncrypterPlaintext)
	if err != nil {
		panic(err)
	}
	s, err = e.DecryptString(s)
	if err != nil {
		panic(err)
	}
	if testEncrypterPlaintext != s {
		t.Errorf("encode/decode string: want %s, got %s", testEncrypterPlaintext, s)
	}
}

func TestBase64Encrypter(t *testing.T) {
	e, err := New(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	s, err := e.EncryptBase64(testEncrypterPlaintext)
	if err != nil {
		panic(err)
	}
	s, err = e.DecryptBase64(s)
	if err != nil {
		panic(err)
	}
	if testEncrypterPlaintext != s {
		t.Errorf("encode/decode base64: want %s, got %s", testEncrypterPlaintext, s)
	}
}

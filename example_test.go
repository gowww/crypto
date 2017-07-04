package crypto_test

import (
	"bytes"
	"fmt"

	"github.com/gowww/crypto"
)

func ExampleNewEncrypter() {
	data := []byte("data to encrypt")
	encrypter, _ := crypto.NewEncrypter([]byte("secret-key-secret-key-secret-key"))
	encData, _ := encrypter.Encrypt(data)
	data, _ = encrypter.Decrypt(encData)
	fmt.Println(bytes.Equal(data, encData))
}

package encrypt_test

import (
	"fmt"

	"github.com/gowww/encrypt"
)

func Example() {
	encrypter, _ := encrypt.New("secret-key-secret-key-secret-key")
	data, _ := encrypter.EncryptString("data to encrypt")
	data, _ = encrypter.DecryptString(data)
	fmt.Println(data)
}

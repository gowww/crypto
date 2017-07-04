# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) encrypt [![GoDoc](https://godoc.org/github.com/gowww/crypto?status.svg)](https://godoc.org/github.com/gowww/crypto) [![Build](https://travis-ci.org/gowww/crypto.svg?branch=master)](https://travis-ci.org/gowww/crypto) [![Coverage](https://coveralls.io/repos/github/gowww/crypto/badge.svg?branch=master)](https://coveralls.io/github/gowww/crypto?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/crypto)](https://goreportcard.com/report/github.com/gowww/crypto)

Package [crypto](https://godoc.org/github.com/gowww/crypto) provides encryption and hashing utilities.

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/crypto
	```

2. Import it in your code:

	```Go
	import "github.com/gowww/crypto"
	```

## Usage

Use [NewEncrypter](https://godoc.org/github.com/gowww/crypto#NewEncrypter) with a 32 bytes long secret key to make a new [Encrypter](https://godoc.org/github.com/gowww/crypto#Encrypter):

```Go
encrypter, _ := crypto.NewEncrypter([]byte("secret-key-secret-key-secret-key"))
```

Use [Encrypter.Encrypt](https://godoc.org/github.com/gowww/crypto#Encrypter.Encrypt) or [Encrypter.EncryptBase64](https://godoc.org/github.com/gowww/crypto#Encrypter.EncryptBase64) to encrypt a value:

```Go
encryptedData, _ := encrypter.Encrypt([]byte("data to encrypt"))
```

Use [Encrypter.Decrypt](https://godoc.org/github.com/gowww/crypto#Encrypter.Decrypt) or [Encrypter.DecryptBase64](https://godoc.org/github.com/gowww/crypto#Encrypter.DecryptBase64) to decrypt a value:

```Go
decryptedData, _ := encrypter.Decrypt(encryptedData)
```

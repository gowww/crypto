# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) encrypt [![GoDoc](https://godoc.org/github.com/gowww/encrypt?status.svg)](https://godoc.org/github.com/gowww/encrypt) [![Build](https://travis-ci.org/gowww/encrypt.svg?branch=master)](https://travis-ci.org/gowww/encrypt) [![Coverage](https://coveralls.io/repos/github/gowww/encrypt/badge.svg?branch=master)](https://coveralls.io/github/gowww/encrypt?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/encrypt)](https://goreportcard.com/report/github.com/gowww/encrypt)

Package [encrypt](https://godoc.org/github.com/gowww/encrypt) provides encryption utilities.

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/encrypt
	```

2. Import it in your code:

	```Go
	import "github.com/gowww/encrypt"
	```

## Usage

Use [New](https://godoc.org/github.com/gowww/encrypt#New) with a 32 bytes long secret key to make a new encrypter:

```Go
encrypter, _ := encrypt.New("secret-key-secret-key-secret-key")
```

Use [Encrypter.Encrypt](https://godoc.org/github.com/gowww/encrypt#Encrypter.Encrypt) or [Encrypter.EncryptString](https://godoc.org/github.com/gowww/encrypt#Encrypter.EncryptString) to encrypt a value:

```Go
encryptedData, _ := encrypter.EncryptString("data to encrypt")
```

Use [Encrypter.Decrypt](https://godoc.org/github.com/gowww/encrypt#Encrypter.Decrypt) or [Encrypter.DecryptString](https://godoc.org/github.com/gowww/encrypt#Encrypter.DecryptString) to decrypt a value:

```Go
decryptedData, _ := encrypter.DecryptString("data to encrypt")
```

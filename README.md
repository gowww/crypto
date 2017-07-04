# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) encrypt [![GoDoc](https://godoc.org/github.com/gowww/crypto?status.svg)](https://godoc.org/github.com/gowww/crypto) [![Build](https://travis-ci.org/gowww/crypto.svg?branch=master)](https://travis-ci.org/gowww/crypto) [![Coverage](https://coveralls.io/repos/github/gowww/crypto/badge.svg?branch=master)](https://coveralls.io/github/gowww/crypto?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/crypto)](https://goreportcard.com/report/github.com/gowww/crypto)

Package [crypto](https://godoc.org/github.com/gowww/crypto) provides encryption utilities.

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

Use [NewEncoder](https://godoc.org/github.com/gowww/crypto#NewEncoder) with a 32 bytes long secret key to make a new encoder:

```Go
encoder, _ := crypto.NewEncoder("secret-key-secret-key-secret-key")
```

Use [Encoder.Encrypt](https://godoc.org/github.com/gowww/crypto#Encoder.Encrypt) or [Encoder.EncryptString](https://godoc.org/github.com/gowww/crypto#Encoder.EncryptString) to encrypt a value:

```Go
encryptedData, _ := encoder.EncryptString("data to encrypt")
```

Use [Encoder.Decrypt](https://godoc.org/github.com/gowww/crypto#Encoder.Decrypt) or [Encoder.DecryptString](https://godoc.org/github.com/gowww/crypto#Encoder.DecryptString) to decrypt a value:

```Go
decryptedData, _ := encoder.DecryptString("data to encrypt")
```

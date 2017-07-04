// Package encrypt provides encryption utilities.
package encrypt

import "crypto/rand"

// Random sources.
const (
	SourceNum           = "0123456789"
	SourceAlphaLower    = "abcdefghijklmnopqrstuvwxyz"
	SourceAlphaUpper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SourceAlphaNumLower = "abcdefghijklmnopqrstuvwxyz0123456789"
	SourceAlphaNumUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SourceAlpha         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SourceAlphaNum      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SourceAll           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=~`[]{}|\\:;\"'<>,.?/"
)

// Random returns random bytes.
func Random(n int) ([]byte, error) {
	var bb = make([]byte, n)
	_, err := rand.Read(bb)
	if err != nil {
		panic(err)
	}
	return bb, nil
}

// RandomSourced returns a random string.
func RandomSourced(src string, n int) (string, error) {
	bb, err := Random(n)
	if err != nil {
		return "", err
	}
	for i, b := range bb {
		bb[i] = src[b%byte(len(src))]
	}
	return string(bb), nil
}

// RandomKey returns a random 32 bytes alphanumeric key.
func RandomKey() string {
	k, err := RandomSourced(SourceAlphaNum, 32)
	if err != nil {
		panic(err)
	}
	return k
}

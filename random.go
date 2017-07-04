package crypto

import "crypto/rand"

// Random sources.
var (
	SourceNum           = []byte("0123456789")
	SourceAlphaLower    = []byte("abcdefghijklmnopqrstuvwxyz")
	SourceAlphaUpper    = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	SourceAlphaNumLower = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	SourceAlphaNumUpper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	SourceAlpha         = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	SourceAlphaNum      = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	SourceAll           = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=~`[]{}|\\:;\"'<>,.?/")
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

// RandomSourced returns a random string only containing bytes from a source string.
func RandomSourced(src []byte, n int) ([]byte, error) {
	bb, err := Random(n)
	if err != nil {
		return nil, err
	}
	for i, b := range bb {
		bb[i] = src[b%byte(len(src))]
	}
	return bb, nil
}

// RandomKey returns a random 32 bytes alphanumeric key.
func RandomKey() string {
	k, err := RandomSourced(SourceAlphaNum, 32)
	if err != nil {
		panic(err)
	}
	return string(k)
}

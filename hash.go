package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// HashMD5 returns the MD5 checksum of source.
func HashMD5(src io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// HashFileMD5 returns the MD5 checksum of a file.
func HashFileMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return HashMD5(f)
}

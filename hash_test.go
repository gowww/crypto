package crypto

import (
	"os"
	"testing"
)

var testsHashFile = []struct {
	path string
	want string
}{
	{"LICENSE", "edbc4f9728a8e311b55e081b27e3caff"},
	{"unknown", ""},
}

func TestHashFileMD5(t *testing.T) {
	for _, tt := range testsHashFile {
		got, err := HashFileMD5(tt.path)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}
		if got != tt.want {
			t.Errorf("%q: want %q, got %q", tt.path, tt.want, got)
		}
	}
}

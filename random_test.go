package encrypt

import (
	"testing"
)

func TestRandomKey(t *testing.T) {
	k := RandomKey()
	if len(k) != 32 {
		t.Errorf("random key length: want 32, got %d", len(k))
	}
}

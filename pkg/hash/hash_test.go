package hash

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	as := assert.New(t)

	h := Hasher{}
	data := []byte("hello")
	got := h.encode(data)
	want := sha256.Sum256(data)
	as.Equal(want, got)
}

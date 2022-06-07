package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodecEncode(t *testing.T) {
	as := assert.New(t)

	c := Codec{}

	got := c.encode([]byte("hello"))
	want := encoding.EncodeToString([]byte("hello"))

	as.Equal(want, got)
}

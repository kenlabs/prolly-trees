package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	as := assert.New(t)

	k := StringKind
	got := k.String()
	want := "string"

	as.Equal(got, want)
}

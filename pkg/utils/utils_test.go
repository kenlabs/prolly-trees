package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	as := assert.New(t)

	a, b := 1, 2
	got := Min(a, b)
	want := 1

	as.Equal(want, got)
}

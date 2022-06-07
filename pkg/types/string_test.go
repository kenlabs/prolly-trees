package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Equal(t *testing.T) {
	as := assert.New(t)

	s1 := String{
		s: "hello",
	}

	s2 := String{
		s: "hello",
	}

	as.Equal(s1, s2)
}

func TestString_Less(t *testing.T) {
	as := assert.New(t)

	s1 := String{
		s: "abcd",
	}

	s2 := String{
		s: "abce",
	}

	got := s1.Less(s2)

	as.Equal(true, got)
}

func TestString_KindOf(t *testing.T) {
	as := assert.New(t)

	s := String{s: "hello"}
	got := s.KindOf()

	as.Equal(StringKind, got)
}

func TestString_ValueOf(t *testing.T) {
	as := assert.New(t)

	s := String{s: "hello"}
	got := s.ValueOf()

	as.Equal("hello", got)
}

package types

import "github.com/kenlab/prolly-trees/go/utils"

type String struct {
	s string
}

func (s String) Equal(other Value) bool {
	if other.KindOf() != StringKind {
		return false
	}
	t := other.(String)
	if len(s.s) != len(t.s) {
		return false
	}
	for i := 0; i < len(s.s); i++ {
		if s.s[i] != t.s[i] {
			return false
		}
	}
	return true
}

// Less return true if this string is less than the other string lexicographically
func (s String) Less(other Value) bool {
	if other.KindOf() != StringKind {
		return false
	}
	t := other.(String)
	p, q := s.s, t.s
	for i := 0; i < utils.Min(len(p), len(q)); i++ {
		if p[i] < q[i] {
			return true
		}
		if p[i] > q[i] {
			return false
		}
	}
	return len(p) < len(q)
}

func (s String) KindOf() Kind {
	return StringKind
}

func (s String) ValueOf() interface{} {
	return s.s
}

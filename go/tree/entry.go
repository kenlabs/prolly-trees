package tree

import (
	"github.com/kenlab/prolly-trees/go/codec"
	"github.com/kenlab/prolly-trees/go/hash"
	"github.com/kenlab/prolly-trees/go/types"
)

type Entry struct {
	Key     types.Value
	Address string
	Codec   codec.Codec
	Hasher  hash.Hasher
}

type Entries []Entry

// implement the sort interface

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	return e[i].Key.Less(e[j].Key)
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

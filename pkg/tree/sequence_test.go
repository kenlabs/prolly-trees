package tree

import (
	"github.com/kenlab/prolly-trees/pkg/codec"
	"github.com/kenlab/prolly-trees/pkg/hash"
	"github.com/kenlab/prolly-trees/pkg/types"
	"github.com/kenlab/prolly-trees/pkg/utils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSequence_Find(t *testing.T) {
	as := assert.New(t)
	e := Entry{
		Key:     types.NewValue("hello"),
		Address: "",
		Codec:   codec.Codec{},
		Hasher:  hash.Hasher{},
	}
	foo := NewSequence(Entries{e}, false)
	res := foo.Find(types.NewValue("hello"))
	as.Equal(e, res)
}

func TestSequence_MultiFind(t *testing.T) {
	as := assert.New(t)

	strs := []string{"go", "java", "python", "cpp", "rust"}
	var es Entries
	for _, v := range strs {
		es = append(es, Entry{
			Key:     types.NewValue(v),
			Address: "",
			Codec:   codec.Codec{},
			Hasher:  hash.Hasher{},
		})
	}
	foo := NewSequence(es, true)

	var vs types.Values
	for _, v := range strs[1:3] {
		vs = append(vs, types.NewValue(v))
	}

	got := foo.MultiFind(vs)
	want := es[1:3]

	as.Equal(len(want), len(got))

	for i := 0; i < utils.Min(len(want), len(got)); i++ {
		as.Equal(got[i], want[i])
	}

}

func TestSequence_RangeFind(t *testing.T) {
	as := assert.New(t)

	var es Entries
	for i := 0; i < 10; i++ {
		x := strconv.Itoa(i)
		es = append(es, Entry{
			Key:     types.NewValue(x),
			Address: "",
			Codec:   codec.Codec{},
			Hasher:  hash.Hasher{},
		})
	}
	foo := NewSequence(es, false)
	got := foo.RangeFind(types.NewValue(strconv.Itoa(3)), types.NewValue(strconv.Itoa(7)))
	want := es[3:8]

	as.Equal(len(want), len(got))

	for i := 0; i < utils.Min(len(want), len(got)); i++ {
		as.Equal(want[i], got[i])
	}
}

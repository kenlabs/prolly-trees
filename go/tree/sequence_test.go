package tree

import (
	"fmt"
	"github.com/kenlab/prolly-trees/go/codec"
	"github.com/kenlab/prolly-trees/go/hash"
	"github.com/kenlab/prolly-trees/go/types"
	"strconv"
	"testing"
)

func TestSequence_Find(t *testing.T) {
	e1 := Entry{
		Key:     types.NewValue("hello"),
		Address: "",
		Codec:   codec.Codec{},
		Hasher:  hash.Hasher{},
	}
	foo := NewSequence(Entries{e1}, true)
	fmt.Printf("%#v\n", foo.Find(types.NewValue("hello")))
}

func TestSequence_MultiFind(t *testing.T) {
	e1 := Entry{
		Key:     types.NewValue("go"),
		Address: "",
		Codec:   codec.Codec{},
		Hasher:  hash.Hasher{},
	}
	e2 := Entry{
		Key:     types.NewValue("java"),
		Address: "",
		Codec:   codec.Codec{},
		Hasher:  hash.Hasher{},
	}
	e3 := Entry{
		Key:     types.NewValue("python"),
		Address: "",
		Codec:   codec.Codec{},
		Hasher:  hash.Hasher{},
	}
	foo := NewSequence(Entries{e1, e2, e3}, true)
	res := foo.MultiFind(types.Values{
		types.NewValue("java"),
		types.NewValue("go"),
	})
	fmt.Printf("%#v\n", res)
}

func TestSequence_RangeFind(t *testing.T) {
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
	fmt.Printf("%#v\n", foo.RangeFind(types.NewValue(strconv.Itoa(3)), types.NewValue(strconv.Itoa(7))))
}

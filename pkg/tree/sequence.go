package tree

import (
	"github.com/kenlab/prolly-trees/pkg/types"
)

type Sequence struct {
	EntrySet Entries
	Closed   bool
	Begin    types.Value
}

func NewSequence(entries Entries, closed bool) Sequence {
	return Sequence{
		EntrySet: entries,
		Closed:   closed,
		Begin:    entries[0].Key,
	}
}

// Find return entry that's key is equal to parameter
func (s Sequence) Find(key types.Value) Entry {
	for _, v := range s.EntrySet {
		if v.Key.Equal(key) {
			return v
		}
	}
	return Entry{}
}

// MultiFind return the result of find all entries whose key is in keys parameter
func (s Sequence) MultiFind(keys types.Values) Entries {

	var res Entries

	vi := make(map[types.Value]bool)

	for _, v := range keys {
		vi[v] = true
	}

	for _, v := range s.EntrySet {
		if vi[v.Key] {
			res = append(res, v)
		}
	}

	return res
}

// RangeFind return the entries whose key is in [start, end]
func (s Sequence) RangeFind(start, end types.Value) Entries {
	res := s.EntrySet

	first, last := 0, -1
	f1, f2 := true, true

	for i := 0; i < s.EntrySet.Len(); i++ {
		entry := s.EntrySet[i]
		if f1 {
			if !entry.Key.Less(start) {
				first = i
				f1 = false
			}
		}
		if f2 {
			if !end.Less(entry.Key) {
				last = i
			} else {
				f2 = false
				break
			}
		}
	}

	if (!f1) && (!f2) {
		return res[first : last+1]
	}
	return nil
}

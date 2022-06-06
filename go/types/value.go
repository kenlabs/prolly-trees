package types

type Value interface {
	Equal(other Value) bool

	Less(other Value) bool

	KindOf() Kind

	ValueOf() interface{}
}

type Values []Value

// implement the sort interface

func (vl Values) Len() int {
	return len(vl)
}

func (vl Values) Less(a, b int) bool {
	return vl[a].Less(vl[b])
}

func (vl Values) Swap(a, b int) {
	vl[a], vl[b] = vl[b], vl[a]
}

func NewValue(a interface{}) Value {
	switch a.(type) {
	case string:
		return String{
			s: a.(string),
		}
	}
	return nil
}

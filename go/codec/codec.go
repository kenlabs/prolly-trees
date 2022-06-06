package codec

import "encoding/base32"

var encoding = base32.NewEncoding("kenlababcdefghijklmnopqrstuvwxyz")

type Codec struct {
}

func (c Codec) encode(src []byte) string {
	return encoding.EncodeToString(src)
}

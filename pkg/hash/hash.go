package hash

import "crypto/sha256"

type Hasher struct {
}

func (h Hasher) encode(src []byte) [32]byte {
	return sha256.Sum256(src)
}

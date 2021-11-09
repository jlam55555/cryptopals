package set1

import (
	"encoding/hex"
)

func RepeatingKeyXor(s, key string) string {
	raw := []byte(s)
	for i := range raw {
		raw[i] ^= key[i%len(key)]
	}
	return hex.EncodeToString(raw)
}

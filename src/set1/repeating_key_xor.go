package set1

import (
	"encoding/hex"
)

func RepeatingKeyXor(s, cipher string) string {
	raw := []byte(s)
	for i := range raw {
		raw[i] ^= cipher[i%len(cipher)]
	}
	return hex.EncodeToString(raw)
}

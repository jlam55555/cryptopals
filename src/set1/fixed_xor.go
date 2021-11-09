package set1

import (
	"encoding/hex"
	"log"
)

func FixedXor(s1, s2 string) string {
	s1_raw, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal("invalid hex string s1")
	}

	s2_raw, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal("invalid hex string s2")
	}

	for i, c := range s2_raw {
		s1_raw[i] ^= c
	}

	return hex.EncodeToString(s1_raw)
}

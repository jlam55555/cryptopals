package set1

import (
	"encoding/hex"
	"log"
)

// returns a likelihood score of being generated by ECB
// by counting equivalent blocks
func AesEcbDetect(ciphertext string) int {
	// assume AES-128 block size
	bs := 16

	raw, err := hex.DecodeString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	seen, score := make(map[string]int), 0

	for start, end := 0, bs; start < len(raw); start, end = start+bs, end+bs {
		block := string(raw[start:end])
		seen[block]++
		if seen[block] > 1 {
			score++
		}
	}

	return score
}
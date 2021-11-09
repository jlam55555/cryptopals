package set1

import (
	"encoding/hex"
	"log"
)

// countSetBits counts the number of set bits in
// a byte using Kernighan's algorithm
func countSetBits(b byte) int {
	count := 0
	for b > 0 {
		b &= b - 1
		count++
	}
	return count
}

// hamming calculates the Hamming distance between two
// strings (number of differing bits)
func hamming(s1, s2 []byte) int {
	if len(s1) != len(s2) {
		log.Fatal("hamming: strings of different length")
	}

	distance := 0
	for i, c := range s2 {
		distance += countSetBits(s1[i] ^ c)
	}

	return distance
}

func BreakRepeatingKeyXor(ciphertext []byte) (string, string) {
	keysizeMin, keysizeMax := 2, 40
	ksCandidates, minEditDists := []int{1e9, 1e9, 1e9}, []float64{1e9, 1e9, 1e9}

	for ks := keysizeMin; ks <= keysizeMax; ks++ {
		blocks := [][]byte{ciphertext[:ks], ciphertext[ks : 2*ks], ciphertext[2*ks : 3*ks], ciphertext[3*ks : 4*ks]}
		editDist := (hamming(blocks[0], blocks[1]) +
			hamming(blocks[1], blocks[2]) +
			hamming(blocks[2], blocks[3])) / 3.
		normEditDist := float64(editDist) / float64(ks)

		if normEditDist < minEditDists[0] {
			// shift over old candidates
			minEditDists[1], minEditDists[2] = minEditDists[0], minEditDists[1]
			ksCandidates[1], ksCandidates[2] = ksCandidates[1], ksCandidates[2]
			// update new candidate
			minEditDists[0], ksCandidates[0] = normEditDist, ks
		} else if normEditDist < minEditDists[1] {
			minEditDists[2] = minEditDists[1]
			ksCandidates[2] = ksCandidates[1]
			minEditDists[1], ksCandidates[1] = normEditDist, ks
		} else if normEditDist < minEditDists[2] {
			minEditDists[2], ksCandidates[2] = normEditDist, ks
		}
	}

	bestScore, bestKey, bestText := 1e9, "", ""

	// try each of the candidate keysizes
	for _, ks := range ksCandidates {
		// fewer than 3 key size candidates were found
		if ks == 1e9 {
			break
		}

		key := make([]byte, ks)

		// truncate to multiple of ks for simplicity
		// transpose the blocks
		for i := 0; i < ks; i++ {
			col := make([]byte, len(ciphertext)/ks)
			for j := 0; j < len(ciphertext)/ks; j++ {
				col[j] = ciphertext[ks*j+i]
			}

			// solve for best letter
			_, _, keyChar := SingleByteXorCipher(hex.EncodeToString(col))
			key[i] = keyChar
		}

		plaintext := string(RepeatingKeyXor(ciphertext, key))
		if score := scoreString(plaintext); score < bestScore {
			bestScore = score
			bestKey = string(key)
			bestText = plaintext
		}
	}

	return bestText, bestKey
}

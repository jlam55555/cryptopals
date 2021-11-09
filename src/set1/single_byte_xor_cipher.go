package set1

import (
	"encoding/hex"
	"log"
	"math"
)

// relative frequencies of letters in the English alphabet,
// according to https://en.wikipedia.org/wiki/Letter_frequency
var freqs = [26]float64{
	8.2, 1.5, 2.8, 4.3, 13,
	2.2, 2, 6.1, 7, 0.15,
	0.77, 4, 2.4, 6.7, 7.5,
	1.9, 0.095, 6, 6.3, 9.1,
	2.8, 0.98, 2.4, 0.15, 2,
	0.074,
}

// scoreString uses a chi-square similarity metric on the frequencies
// of alphabetic characters, as suggested here:
// https://crypto.stackexchange.com/a/30259; before I tried
// a simpler "matched filter" approach (simply multiplying
// counts by relative frequencies) but this didn't work as well;
// lower score is better; an arbitrary penalty is added for non-
// alphabetic characters and non-printable characters as well
func scoreString(s string) float64 {
	score := 0.

	// used for calculating relative frequency of alphabetic
	// characters within the string
	var hist [26]int
	totalCount := 0

	for _, c := range s {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			// add score for letters
			// score += freqs[(byte(c) & ^byte(0x20))-'A']
			totalCount++
			hist[(byte(c) & ^byte(0x20))-'A']++
		} else if c != '\n' && (c < 32 || c > 126) {
			// for non-printable characters, apply penalty
			score += 1000
		} else if c != '\n' && c != ' ' {
			// for other characters, apply less-strict penalty
			score += 500
		}
	}

	// chi-square similarity metric
	for i, count := range hist {
		score += math.Pow((float64(count)/float64(totalCount))*100-freqs[i],
			2) / freqs[i]
	}

	return score
}

func SingleByteXorCipher(s string) (string, float64, byte) {
	// convert to raw bytes
	raw, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	minScore, bestString, bestKey := 1e9, "", 0

	// try each xor
	for k := 0; k <= 255; k++ {
		// copy raw bytes (so we don't affect the original)
		tmp := make([]byte, len(raw))
		copy(tmp, raw)

		for j := range raw {
			tmp[j] ^= byte(k)
		}

		// perform matched-filter score
		if score := scoreString(string(tmp)); score < minScore {
			minScore = score
			bestString = string(tmp)
			bestKey = k
		}
	}

	return bestString, minScore, byte(bestKey)
}

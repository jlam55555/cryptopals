package set1

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestChallenge1(t *testing.T) {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	out := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	if res := HexToBase64(s); res != out {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

func TestChallenge2(t *testing.T) {
	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"
	out := "746865206b696420646f6e277420706c6179"

	if res := FixedXor(s1, s2); res != out {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

func TestChallenge3(t *testing.T) {
	s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	out := "Cooking MC's like a pound of bacon"

	if res, _ := SingleByteXorCipher(s); res != out {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

func TestChallenge4(t *testing.T) {
	path := "../../res/4.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res, minScore := "", 1e9
	for scanner.Scan() {
		if str, score := SingleByteXorCipher(scanner.Text()); score < minScore {
			res, minScore = str, score
		}
	}

	out := "Now that the party is jumping\n"
	if res != out {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

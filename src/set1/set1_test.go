package set1

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func TestChallenge5(t *testing.T) {
	s := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	cipher := "ICE"
	out := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	if res := RepeatingKeyXor(s, cipher); res != out {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

func TestChallenge7(t *testing.T) {
	// read input
	inPath := "../../res/7.txt"
	file, err := os.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s := ""
	for scanner.Scan() {
		s += scanner.Text()
	}

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	// read expected output (including newlines)
	outPath := "../../res/7_out.txt"
	outb, err := ioutil.ReadFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	key := []byte("YELLOW SUBMARINE")
	out := string(outb)

	// use HasPrefix to ignore padding at end
	if res := AesDecode(b, key); strings.HasPrefix(res, out) {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

func TestChallenge8(t *testing.T) {
	path := "../../res/8.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res, maxScore := "", -1
	for scanner.Scan() {
		if score := AesEcbDetect(scanner.Text()); score > maxScore {
			maxScore, res = score, scanner.Text()
		}
	}

	out := "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a"

	if res != out {
		t.Errorf("test case failed: got %v, expected %v", res, out)
	}
}

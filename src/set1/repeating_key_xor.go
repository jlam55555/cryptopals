package set1

func RepeatingKeyXor(s, key []byte) []byte {
	b := make([]byte, len(s))
	for i := range s {
		b[i] = s[i] ^ key[i%len(key)]
	}
	return b
}

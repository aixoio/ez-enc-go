package main

import "crypto/sha256"

func sha256_bit_hash_from_string(str string) []byte {
	hash_32 := sha256.Sum256([]byte(str))
	hash := hash_32[:]
	return hash
}

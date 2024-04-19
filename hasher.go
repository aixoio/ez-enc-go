package main

import (
	"crypto/md5"
	"crypto/sha256"
)

func sha256_bit_hash_from_string(str string) []byte {
	hash_32 := sha256.Sum256([]byte(str))
	hash := hash_32[:]
	return hash
}

func md5_128_bit_hash_from_string(str string) []byte {
	hash_16 := md5.Sum([]byte(str))
	hash := hash_16[:]
	return hash
}

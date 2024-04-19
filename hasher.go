package main

import (
	"crypto/md5"
	"crypto/sha256"
	"io"

	"github.com/cxmcc/tiger"
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

func tiger_192_bit_hash_from_string(str string) []byte {
	hasher_24 := tiger.New()
	io.WriteString(hasher_24, str)
	return hasher_24.Sum(nil)
}

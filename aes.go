package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func aes_enc(dat []byte, key []byte) []byte {

	aes, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return []byte{}
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return []byte{}
	}

	cipherbytes := gcm.Seal(nonce, nonce, dat, nil)

	return cipherbytes
}

func aes_dec(dat []byte, key []byte) []byte {

	aes, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return []byte{}
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherbytes := dat[:nonceSize], dat[nonceSize:]

	bytes, err := gcm.Open(nil, nonce, cipherbytes, nil)
	if err != nil {
		return []byte{}
	}

	return bytes
}

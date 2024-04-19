package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func ask_mode() {
	fmt.Println(color.HiWhiteString("Welcome to ez-enc-go where you we can meet all of your AES file encryption needs"))

	fmt.Print(color.HiWhiteString("Encrpyt or Decrypt (e or d) "), color.HiCyanString("? "))
	enc_or_dec := ask()

	if strings.Compare(enc_or_dec, "d") == 0 {
		dec_ask_mode()
		return
	}

	fmt.Print(color.HiWhiteString("Name of the file to encrypt "), color.HiCyanString("? "))
	file_name := ask()

	file_dat, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(color.HiRedString("Error", err.Error()))
		return
	}

	fmt.Print(color.HiWhiteString("What should your key password be "), color.HiCyanString("? "))
	key_pwd := ask()

	fmt.Print(color.HiWhiteString("AES mode (128, 192, 256) "), color.HiCyanString("? "))
	mode := ask()

	var key []byte

	if strings.Compare(mode, "128") == 0 {
		fmt.Println(color.HiWhiteString("Using AES 128-bit GCM MD5 hash"))
		key = md5_128_bit_hash_from_string(key_pwd)
	} else if strings.Compare(mode, "192") == 0 {
		fmt.Println(color.HiWhiteString("Using AES 192-bit GCM Tiger hash"))
		key = tiger_192_bit_hash_from_string(key_pwd)
	} else if strings.Compare(mode, "256") == 0 {
		fmt.Println(color.HiWhiteString("Using AES 256-bit GCM SHA-256 hash"))
		key = tiger_192_bit_hash_from_string(key_pwd)
	} else {
		fmt.Println(color.HiRedString("Invalid AES mode"))
		return
	}

	fmt.Println(color.HiYellowString("Encrpyting file"))

	enc_dat := aes_enc(file_dat, key)

	fmt.Println(color.HiYellowString("Encrypted file"))
	fmt.Println(color.HiYellowString(fmt.Sprintf("Saving file to %s", file_name)))

	err = os.WriteFile(file_name, enc_dat, 0644)
	if err != nil {
		fmt.Println(color.HiRedString("Error", err.Error()))
		return
	}

	fmt.Println(color.HiYellowString(fmt.Sprintf("Saved file to %s", file_name)))

}

func dec_ask_mode() {
	fmt.Print(color.HiWhiteString("Name of the file to decrypt "), color.HiCyanString("? "))
	file_name := ask()

	file_dat, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(color.HiRedString("Error", err.Error()))
		return
	}

	fmt.Print(color.HiWhiteString("What is your key password "), color.HiCyanString("? "))
	key_pwd := ask()

	fmt.Print(color.HiWhiteString("AES mode (128, 192, 256) "), color.HiCyanString("? "))
	mode := ask()

	var key []byte

	if strings.Compare(mode, "128") == 0 {
		fmt.Println(color.HiWhiteString("Using AES 128-bit GCM MD5 hash"))
		key = md5_128_bit_hash_from_string(key_pwd)
	} else if strings.Compare(mode, "192") == 0 {
		fmt.Println(color.HiWhiteString("Using AES 192-bit GCM Tiger hash"))
		key = tiger_192_bit_hash_from_string(key_pwd)
	} else if strings.Compare(mode, "256") == 0 {
		fmt.Println(color.HiWhiteString("Using AES 256-bit GCM SHA-256 hash"))
		key = tiger_192_bit_hash_from_string(key_pwd)
	} else {
		fmt.Println(color.HiRedString("Invalid AES mode"))
		return
	}

	fmt.Println(color.HiYellowString("Decrpyting file"))

	dec_dat := aes_dec(file_dat, key)

	if len(dec_dat) == 0 {
		fmt.Println(color.HiRedString("Bad password"))
		return
	}

	fmt.Println(color.HiYellowString("Decrypted file"))
	fmt.Println(color.HiYellowString(fmt.Sprintf("Saving file to %s", file_name)))

	err = os.WriteFile(file_name, dec_dat, 0644)
	if err != nil {
		fmt.Println(color.HiRedString("Error", err.Error()))
		return
	}

	fmt.Println(color.HiYellowString(fmt.Sprintf("Saved file to %s", file_name)))
}

func ask() string {
	reader := bufio.NewReader(os.Stdin)

	txt, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	txt = strings.TrimSpace(txt)

	return txt
}

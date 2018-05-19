package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	MODE_ENCRYPT = iota
	MODE_DECRYPT
)

func main() {
	fileNamePtr := flag.String(
		"file",
		"",
		"name of file to encrypt or decrypt")

	passwordPtr := flag.String(
		"pass",
		"",
		"password to encypt or decrypt file with")

	encryptModePtr := flag.Bool(
		"encrypt",
		false,
		"pass -encrypt to encrypt file")

	decryptModePtr := flag.Bool(
		"decrypt",
		false,
		"pass -decrypt to decrypt file")

	flag.Parse()

	if *fileNamePtr == "" || *passwordPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *encryptModePtr == false && *decryptModePtr == false {
		flag.PrintDefaults()
		fmt.Println("\nYou must pass -encrypt or -decrypt")
		os.Exit(1)
	}

	if *encryptModePtr == true && *decryptModePtr == true {
		flag.PrintDefaults()
		fmt.Println("\nYou must pass either -encrypt or -decrypt")
		os.Exit(1)
	}

	mode := MODE_ENCRYPT
	if *decryptModePtr == true {
		mode = MODE_DECRYPT
	}

	input, err := ioutil.ReadFile(*fileNamePtr)

	if err != nil {
		fmt.Printf("Could not read file: '%s'\n", *fileNamePtr)
		os.Exit(1)
	}

	var output []byte

	if mode == MODE_ENCRYPT {
		output = encrypt(input, *passwordPtr)
	} else if mode == MODE_DECRYPT {
		output = decrypt(input, *passwordPtr)
	}

	fmt.Printf("%s", string(output))
}

package main

import (
	"fmt"
	"os"
)

const (
	GENERATE = "--generate"
	G        = "-g"
	SIGN     = "--sign"
	S        = "-s"
	VERIFY   = "--verify"
	V        = "-v"
	HELP     = "--help"
	H        = "-h"
)

func handleCommand(command string, args ...string) {
	if command == GENERATE || command == G {
		generate()
	} else if command == SIGN || command == S {
		usageMessage(len(args), 2)
		sign(args...)
	} else if command == VERIFY || command == V {
		usageMessage(len(args), 3)
		verify(args...)
	} else if command == HELP || command == H {
		availableCommands()
	} else {
		panic("")
	}

}

func availableCommands() {
	generate := fmt.Sprintf("%s, %s: %s", GENERATE, G, ". --generate")
	sign := fmt.Sprintf("%s,     %s: %s", SIGN, S, ". --sign private_key file_to_sign")
	verify := fmt.Sprintf("%s,   %s: %s", VERIFY, V, ". --verify public_key signature file_to_sign")
	help := fmt.Sprintf("%s,     %s: %s", HELP, H, ". --help")
	fmt.Printf("Available commands are:\n%s\n%s\n%s\n%s\n", generate, sign, verify, help)
}

func generate() {
	privateKey, publicKey := generateKeyPair()
	dumpToFile(privateKey, "private_key.pem")
	dumpToFile(publicKey, "public_key.pem")
	fmt.Printf("Your new key pair is generated and exported successfully as follows:\n%s\n%s\n", privateKey, publicKey)
}

func sign(args ...string) {
	privateKey := readFromFile(args[0])
	data := readFromFile(args[1])
	signature := generateSignature(privateKey, data)
	dumpToFile(signature, "signature.pem")
	fmt.Printf("Signature is generated and exported successfully as follows:\n%s\n", signature)
}

func verify(args ...string) {
	publicKey := readFromFile(args[0])
	signature := readFromFile(args[1])
	data := readFromFile(args[2])
	if verifySignature(data, signature, publicKey) {
		fmt.Printf("Signature validation -> OK\n")
	} else {
		fmt.Printf("Signature validation -> FAILURE\n")
	}
}

func usageMessage(argsSize int, wantedSize int) {
	if argsSize != wantedSize {
		fmt.Printf("Wrong usage, please run this program with --help command to see correct usages.\n")
		os.Exit(1)
	}
}

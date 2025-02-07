// Copyright (c) 2025 Berk Kirtay

package main

import (
	"fmt"
	"os"
)

const (
	GENERATE = "--generate"
	G        = "-g"
	ALL      = "--all"
	A        = "-a"
	SIGN     = "--sign"
	S        = "-s"
	VERIFY   = "--verify"
	V        = "-v"
	HELP     = "--help"
	H        = "-h"
)

func main() {
	if len(os.Args) < 2 {
		usageMessage(-1, 0)
	}
	handleCommand(os.Args[1], os.Args[2:]...)
}

func handleCommand(command string, args ...string) {
	if command == GENERATE || command == G {
		generate()
	} else if command == ALL || command == A {
		createPayloadFromFiles(args...)
	} else if command == SIGN || command == S {
		usageMessage(len(args), 2)
		sign(args...)
	} else if command == VERIFY || command == V {
		usageMessage(len(args), 3)
		verify(args...)
	} else if command == HELP || command == H {
		availableCommands()
	} else {
		usageMessage(-1, 0)
	}
}

func availableCommands() {
	generate := fmt.Sprintf("%s, %s: %s", GENERATE, G, ". --generate")
	all := fmt.Sprintf("%s,      %s: %s", ALL, A, ". --all payload_file")
	sign := fmt.Sprintf("%s,     %s: %s", SIGN, S, ". --sign private_key file_to_sign")
	verify := fmt.Sprintf("%s,   %s: %s", VERIFY, V, ". --verify public_key signature file_to_sign")
	help := fmt.Sprintf("%s,     %s: %s", HELP, H, ". --help")
	fmt.Printf("Available commands are:\n%s\n%s\n%s\n%s\n%s\n", generate, all, sign, verify, help)
}

func generate() {
	privateKey, publicKey := generateKeyPair()
	dumpToFile(privateKey, "private_key.pem")
	dumpToFile(publicKey, "public_key.pem")
	fmt.Printf(
		"Your new key pair is generated and exported successfully as follows:\n%s\n%s\n",
		privateKey,
		publicKey)
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

func createPayloadFromFiles(args ...string) {
	payload, count := readAll()
	var fileName string = "payload.txt"
	if len(args) > 0 {
		fileName = args[0]
	}
	dumpToFile(payload, fileName)
	fmt.Printf("%d files are merged into a payload file %s.\n", count, fileName)
}

func usageMessage(argsSize int, wantedSize int) {
	if argsSize != wantedSize {
		fmt.Printf("Wrong usage, please run this program with --help command to see correct usages.\n")
		os.Exit(1)
	}
}

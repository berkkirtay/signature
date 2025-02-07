// Copyright (c) 2025 Berk Kirtay

package main

import (
	"os"
	"path/filepath"
)

func dumpToFile(data string, fileName string) {
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func readFromFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func readAll() (string, int) {
	var payload string = ""
	var count int = 0
	err := filepath.Walk("./", func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			return nil
		}
		payload += readFromFile(filePath)
		count++
		return nil
	})

	if err != nil {
		panic(err)
	}
	return payload, count
}

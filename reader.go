package main

import (
	"io"
	"log"
	"os"
)

func reader(path string) []byte {

	var err error

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	return data

}

package main

import (
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("/home/bridgeit/example.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

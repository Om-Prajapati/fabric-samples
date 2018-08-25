package main

import (
	"log"
	"os"
)

func main() {
	oldName := "/home/bridgeit/example.txt"
	newName := "/home/bridgeit/testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

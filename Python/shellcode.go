package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed torrc
var content embed.FS

func main() {
	fileContent, err := readTextFileFromEmbed()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileContent)
}

func readTextFileFromEmbed() (string, error) {
	fileContent, err := content.ReadFile("torrc")
	if err != nil {
		return "", fmt.Errorf("could not read textfile.txt: %v", err)
	}

	return string(fileContent), nil
}


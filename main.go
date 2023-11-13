package main

import (
	"log"

	mdutils "swapnildev.md-utils/pkg"
)

func main() {

	mdFile := mdutils.New("demoFile")
	if mdFile.Err != nil {
		log.Println("an error occurred while creating the file")
	}

	defer mdFile.Close()

	mdFile.Append("Welcome to MD-Utils package", mdutils.H1)
	mdFile.Append("Welcome to MD-Utils package", mdutils.H2)
	mdFile.Append("Welcome to MD-Utils package", mdutils.H3)
	mdFile.Append("Welcome to MD-Utils package", mdutils.Bold)
	mdFile.Append("Welcome to MD-Utils package", mdutils.Italics)
	mdFile.Append("Welcome to MD-Utils package", mdutils.BlockQuote)
}

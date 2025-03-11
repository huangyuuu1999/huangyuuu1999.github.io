package main

import (
	"fmt"
	"log"
	"github.com/barasher/go-exiftool"
)

func main() {
	et, err := exiftool.NewExiftool()
	if err != nil {
		log.Fatalf("Error initializing Exiftool: %v", err)
	}
	defer et.Close()

	metadata := et.ExtractMetadata("c.jpeg")
	if err != nil {
		log.Fatalf("Error extracting metadata: %v", err)
	}

	for _, md := range metadata {
		for key, val := range md.Fields {
			fmt.Printf("%s: %v\n", key, val)
		}
	}
}
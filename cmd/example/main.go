package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ianbibby/link"
)

type Link struct {
	Href string
	Text string
}

func main() {
	const (
		defaultFilePath = "ex1.html"
	)
	var (
		filePath string
	)

	flag.StringVar(&filePath, "file", defaultFilePath, "Path to file for parsing")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	links, err := link.Parse(f)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}

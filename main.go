package main

import (
	"flag"
	"fmt"

	"github.com/tars47/go-html-link-parser/link"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "path", "ex1.html", "a html file path")
}

func main() {

	flag.Parse()

	links := link.Parse(fileName)

	for _, v := range links {
		fmt.Printf("%+v\n", v)
	}

}

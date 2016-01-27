package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatalln(err)
	}
}

// read this:
// https://en.wikipedia.org/wiki/Pipeline_(computing)
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
	err = tpl.Execute(os.Stdout, []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"})
	if err != nil {
		log.Fatalln(err)
	}
}

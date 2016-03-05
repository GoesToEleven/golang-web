package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type model struct {
	state bool
	pictures []string
}

func main() {
	m := model{}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(bs))
	os.Stdout.Write(bs)
}
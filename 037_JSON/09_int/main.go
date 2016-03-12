package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	var data int
	rcvd := `42`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

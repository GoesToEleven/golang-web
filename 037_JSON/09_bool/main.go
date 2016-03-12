package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	var data bool
	rcvd := `true`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

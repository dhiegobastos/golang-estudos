package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com.br")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Header)
}

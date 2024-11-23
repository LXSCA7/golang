package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Insira url: ")
	var inputUrl string
	_, err := fmt.Scan(&inputUrl)
	if err != nil {
		os.Exit(1)
	}
	apiUrl := "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=" + inputUrl
	response, err := http.Get(apiUrl)

	if err != nil {
		os.Exit(1)
	}

	responseData, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("qr-code.png", responseData, 0644)
}

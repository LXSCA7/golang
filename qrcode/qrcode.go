package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func sanitizeFileName(input string) string {
	// Remove caracteres indesejados da URL
	input = strings.ReplaceAll(input, "https://", "")
	input = strings.ReplaceAll(input, "http://", "")
	input = strings.ReplaceAll(input, "/", "-")
	input = strings.ReplaceAll(input, ":", "-")
	return strings.TrimSpace(input)
}

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
	fileName := "qr-code-[" + inputUrl + "]"
	clearFileName := sanitizeFileName(fileName) + ".png"
	os.WriteFile(clearFileName, responseData, 0644)
	fmt.Println("Arquivo salvo. Nome:", clearFileName)
}

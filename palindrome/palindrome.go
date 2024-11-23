package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkPalindrome(arr string) bool {
	arr = strings.Replace(arr, " ", "", -1)
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

func read(msg string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(msg)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err, "\n Please, try again.")
		return read(msg)
	}

	return line
}

func main() {
	word := read("insert word: ")
	if checkPalindrome(word) {
		fmt.Printf("\"%s\" is a palindrome!", word)
		return
	}
	fmt.Printf("%s is not a palindrome!", word)
}

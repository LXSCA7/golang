package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := read("insert sentence: ")
	input = strings.ReplaceAll(input, " ", "")
	if checkPalindrome(input, 0, len(input)-1) {
		fmt.Println(input, "is a palindrome!")
	} else {
		fmt.Println(input, "is not a palindrome!")
	}
}

func checkPalindrome(arr string, start int, end int) bool {
	if start >= end {
		return true
	}

	if arr[start] != arr[end] {
		return false
	}

	return checkPalindrome(arr, start+1, end-1)
}

func read(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	input, err := reader.ReadString('\n')
	if err != nil {
		print("invalid input. try again.")
		return read(msg)
	}

	return strings.TrimSpace(input)
}

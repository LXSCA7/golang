package main

import (
	"fmt"
	"math/rand"
)

func randomNumber(max int) int {
	return rand.Intn(max + 1)
}

func guessNumber() int {
	var numero int
	fmt.Print("Tente adivinhar o numero: ")
	fmt.Scanln(&numero)
	return numero
}

func verifyNumber(number int, random int) {
	if number > random {
		fmt.Println("O número aleatório é menor que", number)
	}
	if number < random {
		fmt.Println("O número aleatório é maior que", number)
	}
}

func main() {
	random := randomNumber(100)
	tentativas := 0
	var numero int
	numero = guessNumber()
	verifyNumber(numero, random)
	tentativas++
	for numero != random {
		numero = guessNumber()
		tentativas++
		verifyNumber(numero, random)
	}

	fmt.Printf("Acertou em %d tentativas papai!\n", tentativas)
}

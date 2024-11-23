package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createTable(char string) [3][3]string {
	var tab [3][3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if char == "TUTORIAL" {
				tab[i][j] = strconv.Itoa(j + 1)
			} else {
				tab[i][j] = char
			}
		}
	}
	return tab
}

func printTable(table [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Print(strconv.Itoa(i+1), " | ")
		for j := 0; j < 3; j++ {
			if j == 0 {
				fmt.Print(" ")
			}
			fmt.Print(table[i][j])
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		if i < 2 {
			fmt.Print("\n")
			fmt.Println("  | ---+---+---")
		}
	}
	fmt.Print("\n\n")
}

func readPos(msg string) []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Erro. tente novamente")
		return readPos(msg)
	}

	returnString := strings.Split(strings.TrimSpace(input), " ")
	if !verifyInput(returnString) {
		return readPos(msg)
	}

	return returnString
}

func verifyInput(arr []string) bool {
	if len(arr) > 2 {
		fmt.Println("Erro. Use: LINHA COLUNA. Ex: 1 3")
		return false
	}

	line, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("error: ", err)
	}

	if !(line <= 3 && line >= 1) {
		fmt.Println("Número deve estar entre 1 e 3")
		return false
	}
	block, err := strconv.Atoi(arr[1])
	if !(block <= 3 && block >= 1) {
		fmt.Println("Número deve estar entre 1 e 3")
		return false
	}

	return true
}

func tutorial() {
	fmt.Println("A tabela abaixo mostra as linhas e posições de cada bloco. Sendo LINHA | COLUNA.")
	fmt.Print("Para escolher a posição para jogar use \"LINHA COLUNA\". Separado por espaço.\n\n")
	table := createTable("TUTORIAL")
	printTable(table)
}

func playerSymbol(num int) string {
	if num%2 == 0 {
		return "X"
	}

	return "O"
}

func fillBlock(table [3][3]string, pos []string, player int) [3][3]string {
	line, err := strconv.Atoi(pos[0])
	if err != nil {
		fmt.Println("erro", err)
	}
	block, err := strconv.Atoi(pos[1])
	if err != nil {
		fmt.Println("erro", err)
	}
	if table[line-1][block-1] == " " {
		table[line-1][block-1] = playerSymbol(player)
	}
	return table
}

func verifyWinner(tab [3][3]string) bool {
	for i := 0; i < 3; i++ {
		if tab[i][0] != " " && tab[i][0] == tab[i][1] && tab[i][1] == tab[i][2] {
			return true
		}

		if tab[0][i] != " " && tab[0][i] == tab[1][i] && tab[1][i] == tab[2][i] {
			return true
		}
	}

	if tab[0][0] != " " && tab[0][0] == tab[1][1] && tab[1][1] == tab[2][2] {
		return true
	}

	if tab[0][2] != " " && tab[0][2] == tab[1][1] && tab[1][1] == tab[2][0] {
		return true
	}

	return false
}

func main() {
	tab := createTable(" ")
	// printTable(tab)
	tutorial()
	maxJogadas := 9
	for i := 0; i < maxJogadas; i++ {
		pos := readPos("Insira a jogada: ")
		tab = fillBlock(tab, pos, i)
		printTable(tab)
		if verifyWinner(tab) {
			fmt.Printf("O jogador '%s' venceu!\n", playerSymbol(i))
			break
		}
	}
	fmt.Println("Fim do jogo!")
}

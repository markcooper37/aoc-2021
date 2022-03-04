package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	cardRows      [][]int
	cardColumns   [][]int
	numbersCalled map[int]bool
	winningNumber int
	winningIndex  int
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(numbers []int, cards []Card) int {
	cards = findWins(numbers, cards)
	return sumUncalledElements(cards[0]) * cards[0].winningNumber
}

func part2(numbers []int, cards []Card) int {
	cards = findWins(numbers, cards)
	return sumUncalledElements(cards[len(cards)-1]) * cards[len(cards)-1].winningNumber
}

func findWins(numbers []int, cards []Card) []Card {
	for index, card := range cards {
		for numIndex, number := range numbers {
			if _, exists := card.numbersCalled[number]; !exists {
				continue
			}
			cards[index].numbersCalled[number] = true
			if checkWin(card) {
				cards[index].winningNumber = number
				cards[index].winningIndex = numIndex + 1
				break
			}
		}
	}
	sort.Slice(cards, func(i, j int) bool { return cards[i].winningIndex < cards[j].winningIndex })
	return cards
}

func checkWin(card Card) bool {
	for j := 0; j <= 4; j++ {
		rowCalledCount := 0
		columnCalledCount := 0
		for k := 0; k <= 4; k++ {
			if card.numbersCalled[card.cardRows[j][k]] {
				rowCalledCount++
			}
			if card.numbersCalled[card.cardColumns[j][k]] {
				columnCalledCount++
			}
		}
		if rowCalledCount == 5 || columnCalledCount == 5 {
			return true
		}
	}
	return false
}

func sumUncalledElements(card Card) int {
	sum := 0
	for number, called := range card.numbersCalled {
		if !called {
			sum += number
		}
	}
	return sum
}

func convertStrToNum(stringNumbers []string) []int {
	numbers := []int{}
	for _, stringNumber := range stringNumbers {
		number, err := strconv.Atoi(stringNumber)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)
	}
	return numbers
}

func readData(fileName string) ([]int, []Card) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbers := convertStrToNum(strings.Split(scanner.Text(), ","))
	cards := []Card{}
	card := Card{cardRows: [][]int{}, cardColumns: [][]int{}, numbersCalled: map[int]bool{}, winningNumber: -1, winningIndex: -1}
	for scanner.Scan() {
		inputRow := convertStrToNum(strings.Fields(scanner.Text()))
		if len(inputRow) != 0 {
			row := []int{}
			for _, number := range inputRow {
				row = append(row, number)
				card.numbersCalled[number] = false
			}
			card.cardRows = append(card.cardRows, row)
		}
		if len(card.cardRows) == 5 {
			for i := 0; i <= 4; i++ {
				column := []int{}
				for j := 0; j <= 4; j++ {
					column = append(column, card.cardRows[j][i])
				}
				card.cardColumns = append(card.cardColumns, column)
			}
			cards = append(cards, card)
			card = Card{cardRows: [][]int{}, cardColumns: [][]int{}, numbersCalled: map[int]bool{}, winningNumber: -1, winningIndex: -1}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers, cards
}

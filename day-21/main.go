package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Player struct {
	points   int
	position int
	wins     *int
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(startPositions []int) int {
	player1 := Player{points: 0, position: startPositions[0]}
	player2 := Player{points: 0, position: startPositions[1]}
	diceNumber := 1
	turns := 0
	losingPoints := 0
	for {
		turns++
		rollTotal := 0
		for i := 1; i <= 3; i++ {
			rollTotal += diceNumber
			diceNumber += 1
			if diceNumber > 100 {
				diceNumber -= 100
			}
		}
		if turns%2 != 0 {
			player1.position += rollTotal
			if player1.position > 10 {
				player1.position = player1.position % 10
			}
			if player1.position == 0 {
				player1.position = 10
			}
			player1.points += player1.position
			if player1.points >= 1000 {
				losingPoints = player2.points
				break
			}
		} else {
			player2.position += rollTotal
			if player2.position > 10 {
				player2.position = player2.position % 10
			}
			if player2.position == 0 {
				player2.position = 10
			}
			player2.points += player2.position
			if player2.points >= 1000 {
				losingPoints = player1.points
				break
			}
		}
	}
	return losingPoints * turns * 3
}

func part2(startPositions []int) int {
	player1Wins := 0
	player2Wins := 0
	player1 := Player{points: 0, position: startPositions[0], wins: &player1Wins}
	player2 := Player{points: 0, position: startPositions[1], wins: &player2Wins}
	turn := 0
	takeTurn(player1, player2, 3, turn, 1)
	takeTurn(player1, player2, 4, turn, 3)
	takeTurn(player1, player2, 5, turn, 6)
	takeTurn(player1, player2, 6, turn, 7)
	takeTurn(player1, player2, 7, turn, 6)
	takeTurn(player1, player2, 8, turn, 3)
	takeTurn(player1, player2, 9, turn, 1)
	if player1Wins >= player2Wins {
		return player1Wins
	}
	return player2Wins
}

func takeTurn(player1, player2 Player, roll, turn, total int) {
	turn++
	if turn%2 != 0 {
		player1.position += roll
		if player1.position > 10 {
			player1.position = player1.position % 10
		}
		if player1.position == 0 {
			player1.position = 10
		}
		player1.points += player1.position
		if player1.points >= 21 {
			*player1.wins += total
			return
		}
	} else {
		player2.position += roll
		if player2.position > 10 {
			player2.position = player2.position % 10
		}
		if player2.position == 0 {
			player2.position = 10
		}
		player2.points += player2.position
		if player2.points >= 21 {
			*player2.wins += total
			return
		}
	}
	takeTurn(player1, player2, 3, turn, total)
	takeTurn(player1, player2, 4, turn, 3*total)
	takeTurn(player1, player2, 5, turn, 6*total)
	takeTurn(player1, player2, 6, turn, 7*total)
	takeTurn(player1, player2, 7, turn, 6*total)
	takeTurn(player1, player2, 8, turn, 3*total)
	takeTurn(player1, player2, 9, turn, total)
}

func readData(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	scanner := bufio.NewScanner(file)
	startPositions := []int{}
	for scanner.Scan() {
		stringData := strings.FieldsFunc(scanner.Text(), splitCriteria)
		position, err := strconv.Atoi(stringData[1])
		if err != nil {
			log.Fatal(err)
		}
		startPositions = append(startPositions, position)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return startPositions
}

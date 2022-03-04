package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(bracketsList []string) int {
	finalScore := 0
	scoreMap := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	bracketPairs := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	for _, brackets := range bracketsList {
		openBrackets := []rune{}
		for _, bracket := range brackets {
			if bracket == '(' || bracket == '[' || bracket == '{' || bracket == '<' {
				openBrackets = append([]rune{bracket}, openBrackets...)
			} else if len(openBrackets) == 0 || bracketPairs[bracket] != openBrackets[0] {
				finalScore += scoreMap[bracket]
				break
			} else {
				openBrackets = openBrackets[1:]
			}
		}
	}
	return finalScore
}

func part2(bracketsList []string) int {
	scoresToCalculate := [][]rune{}
	bracketPairs := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	for _, brackets := range bracketsList {
		openBrackets := []rune{}
		corrupt := false
		for _, bracket := range brackets {
			if bracket == '(' || bracket == '[' || bracket == '{' || bracket == '<' {
				openBrackets = append([]rune{bracket}, openBrackets...)
			} else if len(openBrackets) == 0 || bracketPairs[bracket] != openBrackets[0] {
				corrupt = true
				break
			} else {
				openBrackets = openBrackets[1:]
			}
		}
		if corrupt {
			continue
		}
		scoresToCalculate = append(scoresToCalculate, openBrackets)
	}
	scoreMap := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	calculatedScores := []int{}
	for _, scoreToCalculate := range scoresToCalculate {
		score := 0
		for _, bracket := range scoreToCalculate {
			score *= 5
			score += scoreMap[bracket]
		}
		calculatedScores = append(calculatedScores, score)
	}
	sort.Slice(calculatedScores, func(i, j int) bool { return calculatedScores[i] > calculatedScores[j] })
	return calculatedScores[(len(calculatedScores)-1)/2]
}

func readData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(data [][]string) int {
	count := 0
	for i := range data {
		for j := 10; j < 14; j++ {
			strLen := len(data[i][j])
			if strLen == 2 || strLen == 3 || strLen == 4 || strLen == 7 {
				count++
			}
		}
	}
	return count
}

func part2(data [][]string) int {
	total := 0
	for i := range data {
		numberLetterReps := [10]string{}
		twoThreeFive := []string{}
		zeroSixNine := []string{}
		for j := 0; j < 10; j++ {
			switch len(data[i][j]) {
			case 2:
				numberLetterReps[1] = data[i][j]
			case 3:
				numberLetterReps[7] = data[i][j]
			case 4:
				numberLetterReps[4] = data[i][j]
			case 5:
				twoThreeFive = append(twoThreeFive, data[i][j])
			case 6:
				zeroSixNine = append(zeroSixNine, data[i][j])
			case 7:
				numberLetterReps[8] = data[i][j]
			}
		}
		for _, number := range zeroSixNine {
			if containsString(number, numberLetterReps[4]) {
				numberLetterReps[9] = number
			} else if containsString(number, numberLetterReps[1]) {
				numberLetterReps[0] = number
			} else {
				numberLetterReps[6] = number
			}
		}
		for _, number := range twoThreeFive {
			if containsString(number, numberLetterReps[1]) {
				numberLetterReps[3] = number
			} else if containsString(numberLetterReps[9], number) {
				numberLetterReps[5] = number
			} else {
				numberLetterReps[2] = number
			}
		}
		numberToAdd := 0
		for j := 10; j < 14; j++ {
			for index, value := range numberLetterReps {
				if data[i][j] == value {
					numberToAdd += index * int(math.Pow(10, float64(13-j)))
					break
				}
			}
		}
		total += numberToAdd
	}
	return total
}

func containsString(bigString, smallString string) bool {
	if len(smallString) > len(bigString) {
		return false
	}
	for _, char1 := range smallString {
		missingChar := true
		for _, char2 := range bigString {
			if char1 == char2 {
				missingChar = false
			}
		}
		if missingChar {
			return false
		}
	}
	return true
}

func readData(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	scanner := bufio.NewScanner(file)
	data := [][]string{}
	for scanner.Scan() {
		newStrings := strings.FieldsFunc(scanner.Text(), splitCriteria)
		sortedStrings := []string{}
		for _, letters := range newStrings {
			splitLetters := strings.Split(letters, "")
			sort.Strings(splitLetters)
			sortedLetters := strings.Join(splitLetters, "")
			sortedStrings = append(sortedStrings, sortedLetters)
		}
		data = append(data, sortedStrings)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

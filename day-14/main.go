package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(template []string, insertionRules map[string]string) int {
	return elementDifference(template, insertionRules, 10)
}

func part2(template []string, insertionRules map[string]string) int {
	return elementDifference(template, insertionRules, 40)
}

func elementDifference(template []string, insertionRules map[string]string, steps int) int {
	pairCounts := map[string]int{}
	for j := 0; j < len(template)-1; j++ {
		pair := fmt.Sprintf("%s%s", template[j], template[j+1])
		pairCounts[pair]++
	}
	for i := 0; i < steps; i++ {
		newPairCounts := map[string]int{}
		for key, value := range pairCounts {
			firstNewPair := fmt.Sprintf("%s%s", string(key[0]), insertionRules[key])
			secondNewPair := fmt.Sprintf("%s%s", insertionRules[key], string(key[1]))
			newPairCounts[firstNewPair] += value
			newPairCounts[secondNewPair] += value
		}
		pairCounts = newPairCounts
	}
	elementCountsMap := map[string]int{}
	for key, value := range pairCounts {
		firstElement := string(key[0])
		secondElement := string(key[1])
		elementCountsMap[firstElement] += value
		elementCountsMap[secondElement] += value
	}
	elementCountsMap[template[0]]++
	elementCountsMap[template[len(template)-1]]++
	for element, value := range elementCountsMap {
		elementCountsMap[element] = value / 2
	}
	elementCounts := []int{}
	for _, value := range elementCountsMap {
		elementCounts = append(elementCounts, value)
	}
	sort.Ints(elementCounts)
	return elementCounts[len(elementCounts)-1] - elementCounts[0]
}

func readData(fileName string) ([]string, map[string]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := strings.Split(scanner.Text(), "")
	scanner.Scan()
	insertionRules := map[string]string{}
	for scanner.Scan() {
		insertionRule := strings.FieldsFunc(scanner.Text(), splitCriteria)
		insertionRules[insertionRule[0]] = insertionRule[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return template, insertionRules
}

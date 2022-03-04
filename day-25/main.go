package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
}

func part1(cucumbers [][]string) int {
	solutionFound := false
	steps := 0
	for !solutionFound {
		solutionFound = true
		steps++
		for rowIndex, row := range cucumbers {
			firstColumnEmpty := false
			if row[0] == "." {
				firstColumnEmpty = true
			}
			movedBefore := false
			for columnIndex, space := range row {
				if movedBefore {
					movedBefore = false
					continue
				}
				if space == ">" {
					if columnIndex < len(row)-1 {
						if row[columnIndex+1] == "." {
							cucumbers[rowIndex][columnIndex] = "."
							cucumbers[rowIndex][columnIndex+1] = ">"
							movedBefore = true
							solutionFound = false
						}
					} else {
						if firstColumnEmpty {
							cucumbers[rowIndex][columnIndex] = "."
							cucumbers[rowIndex][0] = ">"
							solutionFound = false
						}
					}
				}
			}
		}
		for i := 0; i < len(cucumbers[0]); i++ {
			firstRowEmpty := false
			if cucumbers[0][i] == "." {
				firstRowEmpty = true
			}
			movedBefore := false
			for j := 0; j < len(cucumbers); j++ {
				if movedBefore {
					movedBefore = false
					continue
				}
				if cucumbers[j][i] == "v" {
					if j < len(cucumbers)-1 {
						if cucumbers[j+1][i] == "." {
							cucumbers[j][i] = "."
							cucumbers[j+1][i] = "v"
							movedBefore = true
							solutionFound = false
						}
					} else {
						if firstRowEmpty {
							cucumbers[j][i] = "."
							cucumbers[0][i] = "v"
							solutionFound = false
						}
					}
				}
			}
		}
	}
	return steps
}

// the second part of day 25 does not require any code

func readData(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cucumbers := [][]string{}
	for scanner.Scan() {
		cucumberRow := strings.Split(scanner.Text(), "")
		cucumbers = append(cucumbers, cucumberRow)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cucumbers
}

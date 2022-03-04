package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(cavernRisks [][]int) int {
	totalRisks := totalRisks(cavernRisks)
	totalRisks = checkBackwards(cavernRisks, totalRisks)
	return totalRisks[len(totalRisks)-1][len(totalRisks)-1]
}

func part2(cavernRisks [][]int) int {
	for i := 1; i < 5; i++ {
		for j := 0; j < len(cavernRisks); j++ {
			for k := 0; k < len(cavernRisks); k++ {
				if cavernRisks[j][k]+i > 9 {
					cavernRisks[j] = append(cavernRisks[j], (cavernRisks[j][k]+i)%9)
				} else {
					cavernRisks[j] = append(cavernRisks[j], cavernRisks[j][k]+i)
				}
			}
		}
	}
	numberOfRows := len(cavernRisks)
	for i := 1; i < 5; i++ {
		newCavernRisks := [][]int{}
		for j := 0; j < numberOfRows; j++ {
			row := []int{}
			for k := 0; k < len(cavernRisks[j]); k++ {
				if cavernRisks[j][k]+i > 9 {
					row = append(row, (cavernRisks[j][k]+i)%9)
				} else {
					row = append(row, cavernRisks[j][k]+i)
				}
			}
			newCavernRisks = append(newCavernRisks, row)
		}
		cavernRisks = append(cavernRisks, newCavernRisks...)
	}
	totalRisks := totalRisks(cavernRisks)
	totalRisks = checkBackwards(cavernRisks, totalRisks)
	return totalRisks[len(totalRisks)-1][len(totalRisks)-1]
}

func totalRisks(cavernRisks [][]int) [][]int {
	totalRisks := [][]int{}
	for i := 0; i < len(cavernRisks); i++ {
		row := []int{}
		for j := 0; j < len(cavernRisks); j++ {
			row = append(row, 20*len(cavernRisks))
		}
		totalRisks = append(totalRisks, row)
	}
	totalRisks[0][0] = 0
	for i := 1; i < len(cavernRisks); i++ {
		totalRisks[0][i] = cavernRisks[0][i] + totalRisks[0][i-1]
		totalRisks[i][0] = cavernRisks[i][0] + totalRisks[i-1][0]
		for j := 1; j < i; j++ {
			totalRisks[j][i-j] = cavernRisks[j][i-j] + int(math.Min(float64(totalRisks[j-1][i-j]), float64(totalRisks[j][i-j-1])))
		}
	}
	for i := 1; i < len(cavernRisks); i++ {
		for j := i; j <= len(cavernRisks)-1; j++ {
			totalRisks[j][len(cavernRisks)+i-j-1] = cavernRisks[j][len(cavernRisks)+i-j-1] + int(math.Min(float64(totalRisks[j-1][len(cavernRisks)+i-j-1]), float64(totalRisks[j][len(cavernRisks)+i-j-2])))
		}
	}
	return totalRisks
}

func checkBackwards(cavernRisks, totalRisks[][]int) [][]int {
	nothingToUpdate := false
	for !nothingToUpdate {
		noUpdatesSoFar := true
		for rowIndex, row := range totalRisks {
			for columnIndex := range row {
				valuesToCheck := []int{}
				if rowIndex > 0 {
					valuesToCheck = append(valuesToCheck, totalRisks[rowIndex-1][columnIndex])
				}
				if rowIndex < len(totalRisks)-1 {
					valuesToCheck = append(valuesToCheck, totalRisks[rowIndex+1][columnIndex])
				}
				if columnIndex > 0 {
					valuesToCheck = append(valuesToCheck, totalRisks[rowIndex][columnIndex-1])
				}
				if columnIndex < len(row)-1 {
					valuesToCheck = append(valuesToCheck, totalRisks[rowIndex][columnIndex+1])
				}
				sort.Ints(valuesToCheck)
				if valuesToCheck[0] + cavernRisks[rowIndex][columnIndex] < totalRisks[rowIndex][columnIndex] {
					totalRisks[rowIndex][columnIndex] = valuesToCheck[0] + cavernRisks[rowIndex][columnIndex]
					noUpdatesSoFar = false
				}
			}
		}
		if noUpdatesSoFar {
			nothingToUpdate = true
		}
	}
	return totalRisks
}

func readData(fileName string) [][]int {
	cavernRisks := [][]int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riskLevels := []int{}
		stringData := strings.Split(scanner.Text(), "")
		for _, string := range stringData {
			riskLevel, err := strconv.Atoi(string)
			if err != nil {
				log.Fatal(err)
			}
			riskLevels = append(riskLevels, riskLevel)
		}
		cavernRisks = append(cavernRisks, riskLevels)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cavernRisks
}

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

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(tubes [][]int) int {
	riskLevel := 0
	for rowIndex, row := range tubes {
		for columnIndex, value := range row {
			if rowIndex > 0 {
				if value >= tubes[rowIndex-1][columnIndex] {
					continue
				}
			}
			if rowIndex < len(tubes)-1 {
				if value >= tubes[rowIndex+1][columnIndex] {
					continue
				}
			}
			if columnIndex > 0 {
				if value >= tubes[rowIndex][columnIndex-1] {
					continue
				}
			}
			if columnIndex < len(row)-1 {
				if value >= tubes[rowIndex][columnIndex+1] {
					continue
				}
			}
			riskLevel += value + 1
		}
	}
	return riskLevel
}

func part2(tubes [][]int) int {
	checked := map[[2]int]bool{}
	for rowIndex, row := range tubes {
		for columnIndex := range row {
			coordinate := [2]int{rowIndex, columnIndex}
			checked[coordinate] = false
		}
	}
	basins := []int{}
	for rowIndex, row := range tubes {
		for columnIndex := range row {
			if tubes[rowIndex][columnIndex] == 9 {
				continue
			}
			coordinateToCheck := [2]int{rowIndex, columnIndex}
			if !checked[coordinateToCheck] {
				basins = append(basins, 1)
				basins, checked = checkAdjacent(tubes, rowIndex, columnIndex, checked, basins)
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })
	return basins[0] * basins[1] * basins[2]
}

func checkAdjacent(tubes [][]int, rowIndex, columnIndex int, checked map[[2]int]bool, basins []int) ([]int, map[[2]int]bool) {
	coordinateChecked := [2]int{rowIndex, columnIndex}
	checked[coordinateChecked] = true
	if tubes[rowIndex][columnIndex] == 9 {
		return basins, checked
	}
	if rowIndex > 0 {
		coordinateToCheck := [2]int{rowIndex-1, columnIndex}
		if !checked[coordinateToCheck] && tubes[rowIndex-1][columnIndex] < 9 {
			basins[len(basins)-1] += 1
			basins, checked = checkAdjacent(tubes, rowIndex-1, columnIndex, checked, basins)
		}
	}
	if rowIndex < len(tubes)-1 {
		coordinateToCheck := [2]int{rowIndex+1, columnIndex}
		if !checked[coordinateToCheck] && tubes[rowIndex+1][columnIndex] < 9 {
			basins[len(basins)-1] += 1
			basins, checked = checkAdjacent(tubes, rowIndex+1, columnIndex, checked, basins)
		}
	}
	if columnIndex > 0 {
		coordinateToCheck := [2]int{rowIndex, columnIndex-1}
		if !checked[coordinateToCheck] && tubes[rowIndex][columnIndex-1] < 9 {
			basins[len(basins)-1] += 1
			basins, checked = checkAdjacent(tubes, rowIndex, columnIndex-1, checked, basins)
		}
	}
	if columnIndex < len(tubes[0])-1 {
		coordinateToCheck := [2]int{rowIndex, columnIndex+1}
		if !checked[coordinateToCheck] && tubes[rowIndex][columnIndex+1] < 9 {
			basins[len(basins)-1] += 1
			basins, checked = checkAdjacent(tubes, rowIndex, columnIndex+1, checked, basins)
		}
	}
	return basins, checked
}

func readData(fileName string) [][]int {
	data := [][]int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringData := strings.Split(scanner.Text(), "")
		numbers := []int{}
		for _, value := range stringData {
			number, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)
		}
		data = append(data, numbers)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

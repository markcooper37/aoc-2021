package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Octopuses struct {
	octopuses  [][]int
	flashed    map[[2]int]bool
	flashCount int
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(octopuses Octopuses) int {
	for i := 1; i <= 100; i++ {
		for j, row := range octopuses.octopuses {
			for k := range row {
				octopuses.octopuses[j][k]++
				coordinates := [2]int{j, k}
				octopuses.flashed[coordinates] = false
			}
		}
		for j, row := range octopuses.octopuses {
			for k := range row {
				octopuses.flash(j, k)
			}
		}
		for j, row := range octopuses.octopuses {
			for k := range row {
				if octopuses.octopuses[j][k] > 9 {
					octopuses.octopuses[j][k] = 0
				}
			}
		}
	}
	return octopuses.flashCount
}

func part2(octopuses Octopuses) int {
	iteration := 1
	for {
		for j, row := range octopuses.octopuses {
			for k := range row {
				octopuses.octopuses[j][k]++
				coordinates := [2]int{j, k}
				octopuses.flashed[coordinates] = false
			}
		}
		for j, row := range octopuses.octopuses {
			for k := range row {
				octopuses.flash(j, k)
			}
		}
		for j, row := range octopuses.octopuses {
			for k := range row {
				if octopuses.octopuses[j][k] > 9 {
					octopuses.octopuses[j][k] = 0
				}
			}
		}
		allFlashed := true
		for j, row := range octopuses.octopuses {
			for k := range row {
				coordinates := [2]int{j, k}
				if !octopuses.flashed[coordinates] {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			break
		}
		iteration++
	}
	return iteration
}

func (o *Octopuses) flash(row, column int) {
	coordinates := [2]int{row, column}
	if o.flashed[coordinates] || o.octopuses[row][column] < 10 {
		return
	}
	if row < 0 || row > 9 || column < 0 || column > 9 {
		return
	}
	o.flashed[coordinates] = true
	o.flashCount++
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i == row && j == column {
				continue
			}
			if i < 0 || i > 9 || j < 0 || j > 9 {
				continue
			}
			o.octopuses[i][j]++
		}
	}
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i == row && j == column {
				continue
			}
			if i < 0 || i > 9 || j < 0 || j > 9 {
				continue
			}
			o.flash(i, j)
		}
	}
}

func readData(fileName string) Octopuses {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := Octopuses{octopuses: [][]int{}, flashed: map[[2]int]bool{}, flashCount: 0}
	rowIndex := 0
	for scanner.Scan() {
		stringData := strings.Split(scanner.Text(), "")
		numbers := []int{}
		for columnIndex, value := range stringData {
			number, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)
			coordinates := [2]int{rowIndex, columnIndex}
			data.flashed[coordinates] = false
		}
		data.octopuses = append(data.octopuses, numbers)
		rowIndex++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

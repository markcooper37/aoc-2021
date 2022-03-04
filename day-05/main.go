package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"strconv"
	"math"
)

func main() {
	vents := readData("input.txt")
	fmt.Println(part1(vents))
	fmt.Println(part2(vents))
}

func part1(vents [][]int) int {
	grid := [1000][1000]int{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}
	for _, vent := range vents {
		if vent[0] == vent[2] {
			for y := int(math.Min(float64(vent[1]), float64(vent[3]))); y <= int(math.Max(float64(vent[1]), float64(vent[3]))); y++ {
				grid[vent[0]][y]++
			}
		} else if vent[1] == vent[3] {
			for x := int(math.Min(float64(vent[0]), float64(vent[2]))); x <= int(math.Max(float64(vent[0]), float64(vent[2]))); x++ {
				grid[x][vent[1]]++
			}
		}
	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func part2(vents [][]int) int {
	grid := [1000][1000]int{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}
	for _, vent := range vents {
		if vent[0] == vent[2] {
			for y := int(math.Min(float64(vent[1]), float64(vent[3]))); y <= int(math.Max(float64(vent[1]), float64(vent[3]))); y++ {
				grid[vent[0]][y]++
			}
		} else if vent[1] == vent[3] {
			for x := int(math.Min(float64(vent[0]), float64(vent[2]))); x <= int(math.Max(float64(vent[0]), float64(vent[2]))); x++ {
				grid[x][vent[1]]++
			}
		} else {
			switch {
			case vent[0] <= vent[2] && vent[1] <= vent[3]:
				for i := 0; i <= vent[2] - vent[0]; i++ {
					grid[vent[0] + i][vent[1] + i]++
				}
			case vent[0] <= vent[2] && vent[1] > vent[3]:
				for i := 0; i <= vent[2] - vent[0]; i++ {
					grid[vent[0] + i][vent[1] - i]++
				}
			case vent[0] > vent[2] && vent[1] <= vent[3]:
				for i := 0; i <= vent[0] - vent[2]; i++ {
					grid[vent[0] - i][vent[1] + i]++
				}
			case vent[0] > vent[2] && vent[1] > vent[3]:
				for i := 0; i <= vent[0] - vent[2]; i++ {
					grid[vent[0] - i][vent[1] - i]++
				}
			}
		}
	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func readData(fileName string) [][]int {
	data := [][]int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringCoordinates := strings.FieldsFunc(line, splitCriteria)
		coordinates := []int{}
		for _, coordinate := range stringCoordinates {
			intCoordinate, err := strconv.Atoi(coordinate)
			if err != nil {
				log.Fatal(err)
			}
			coordinates = append(coordinates, intCoordinate)
		}
		data = append(data, coordinates)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(enhancement []string, input [][]string) int {
	return countLightPixels(enhancement, input, 2)
}

func part2(enhancement []string, input [][]string) int {
	return countLightPixels(enhancement, input, 50)
}

func countLightPixels(enhancement []string, input [][]string, numberOfApplications int) int {
	pixelToInfinity := "."
	for i := 1; i <= numberOfApplications; i++ {
		extraRow := []string{}
		for j := 0; j < len(input[0])+4; j++ {
			extraRow = append(extraRow, pixelToInfinity)
		}
		updatedInput := [][]string{}
		updatedInput = append(updatedInput, extraRow, extraRow)
		for _, row := range input {
			hashSlice := []string{pixelToInfinity, pixelToInfinity}
			row = append(hashSlice, row...)
			row = append(row, hashSlice...)
			updatedInput = append(updatedInput, row)
		}
		updatedInput = append(updatedInput, extraRow, extraRow)
		newInput := [][]string{}
		for j := 1; j < len(updatedInput)-1; j++ {
			newInput = append(newInput, []string{})
			for k := 1; k < len(updatedInput)-1; k++ {
				binaryValue := ""
				for l := -1; l <= 1; l++ {
					for m := -1; m <= 1; m++ {
						if updatedInput[j+l][k+m] == "#" {
							binaryValue = fmt.Sprintf("%s%s", binaryValue, "1")
						} else {
							binaryValue = fmt.Sprintf("%s%s", binaryValue, "0")
						}
					}
				}
				index, err := strconv.ParseInt(binaryValue, 2, 64)
				if err != nil {
					log.Fatal(err)
				}
				pixel := enhancement[int(index)]
				newInput[j-1] = append(newInput[j-1], pixel)
			}
		}
		input = newInput
		newPixelToInfinity := ""
		for j := 1; j <= 9; j++ {
			if pixelToInfinity == "#" {
				newPixelToInfinity = fmt.Sprintf("%s%s", newPixelToInfinity, "1")
			} else {
				newPixelToInfinity = fmt.Sprintf("%s%s", newPixelToInfinity, "0")
			}
		}
		index, err := strconv.ParseInt(newPixelToInfinity, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		pixelToInfinity = enhancement[index]
	}
	lightCount := 0
	for _, row := range input {
		for _, pixel := range row {
			if pixel == "#" {
				lightCount++
			}
		}
	}
	return lightCount
}

func readData(fileName string) ([]string, [][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	enhancement := strings.Split(scanner.Text(), "")
	scanner.Scan()
	input := [][]string{}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		input = append(input, data)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return enhancement, input
}

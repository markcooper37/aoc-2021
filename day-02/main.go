package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	data := readData("input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data [][]string) int {
	horizontalPosition := 0
	depth := 0
	for i := 0; i < len(data); i++ {
		distance, err := strconv.Atoi(data[i][1])
		if err != nil {
			log.Fatal(err)
		}
		position := data[i][0]
		switch position {
		case "forward":
			horizontalPosition += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		}
	}
	return depth * horizontalPosition
}

func part2(data [][]string) int {
	horizontalPosition := 0
	depth := 0
	aim := 0
	for i := 0; i < len(data); i++ {
		distance, err := strconv.Atoi(data[i][1])
		if err != nil {
			log.Fatal(err)
		}
		position := data[i][0]
		switch position {
		case "forward":
			horizontalPosition += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}
	return depth * horizontalPosition
}

func readData(fileName string) [][]string {
	data := [][]string{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := strings.Split(scanner.Text(), " ")
		data = append(data, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

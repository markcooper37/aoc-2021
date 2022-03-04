package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func main() {
	depths := readDepths("input.txt")
	fmt.Println(part1(depths))
	fmt.Println(part2(depths))
}

func part1(depths []int) int {
	return countIncrease(depths, 1)
}

func part2(depths []int) int {
	return countIncrease(depths, 3)
}

func countIncrease(depths []int, window int) int {
	count := 0
	for i := 0; i < len(depths) - window; i++ {
		if depths[i] < depths[i + window] {
			count++
		}
	}
	return count
}

func readDepths(fileName string) []int {
	depths := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return depths
}
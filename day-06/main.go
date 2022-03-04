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

func part1(allFish [9]int) int {
	return numberOfFish(allFish, 80)
}

func part2(allFish [9]int) int {
	return numberOfFish(allFish, 256)
}

func numberOfFish(allFish [9]int, numberOfDays int) int {
	fishTotal := 0
	for timer := 0; timer <= 8; timer++ {
		fishOnEachDay := []int{1}
		for i := 1; i <= numberOfDays; i++ {
			newFish := 0
			if i == timer+1 {
				newFish++
			}
			if i > 9 {
				newFish += fishOnEachDay[i-9] - fishOnEachDay[i-10]
			}
			if i > 7 {
				newFish += fishOnEachDay[i-7] - fishOnEachDay[i-8]
			}
			fishOnEachDay = append(fishOnEachDay, fishOnEachDay[i-1]+newFish)
		}
		fishTotal += allFish[timer] * fishOnEachDay[numberOfDays]
	}
	return fishTotal
}

func readData(fileName string) [9]int {
	data := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringData := strings.Split(scanner.Text(), ",")
		for _, value := range stringData {
			i, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			data[i]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

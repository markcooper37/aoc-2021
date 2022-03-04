package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Count struct {
	count int
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(cavePairs [][]string, availableCaves map[string]bool) int {
	count := Count{count: 0}
	count.findPaths("start", cavePairs, availableCaves)
	return count.count
}

func part2(cavePairs [][]string, availableCaves map[string]bool) int {
	visitedSmallCaves := map[string]int{}
	for key := range availableCaves {
		if strings.ToLower(key) == key && key != "start" && key != "end" {
			visitedSmallCaves[key] = 0
		}
	}
	count := Count{count: 0}
	count.findPathsPart2("start", cavePairs, availableCaves, visitedSmallCaves)
	return count.count
}

func (c *Count) findPaths(cave string, cavePairs [][]string, availableCaves map[string]bool) {
	if cave == "end" {
		c.count++
		return
	}
	for _, cavePair := range cavePairs {
		if cavePair[0] == cave && availableCaves[cavePair[1]] {
			newAvailableCaves := map[string]bool{}
			for key, value := range availableCaves {
				newAvailableCaves[key] = value
			}
			if strings.ToLower(cavePair[1]) == cavePair[1] {
				newAvailableCaves[cavePair[1]] = false
			}
			c.findPaths(cavePair[1], cavePairs, newAvailableCaves)
		}
		if cavePair[1] == cave && availableCaves[cavePair[0]] {
			newAvailableCaves := map[string]bool{}
			for key, value := range availableCaves {
				newAvailableCaves[key] = value
			}
			if strings.ToLower(cavePair[0]) == cavePair[0] {
				newAvailableCaves[cavePair[0]] = false
			}
			c.findPaths(cavePair[0], cavePairs, newAvailableCaves)
		}
	}
}

func (c *Count) findPathsPart2(cave string, cavePairs [][]string, availableCaves map[string]bool, visitedSmallCaves map[string]int) {
	if cave == "end" {
		c.count++
		return
	}
	smallCaveTwice := true
	for _, value := range visitedSmallCaves {
		if value >= 2 {
			smallCaveTwice = false
		}
	}
	for key, value := range visitedSmallCaves {
		if value >= 1 && !smallCaveTwice {
			availableCaves[key] = false
		}
	}
	for _, cavePair := range cavePairs {
		if cavePair[0] == cave && availableCaves[cavePair[1]] {
			newAvailableCaves := map[string]bool{}
			for key, value := range availableCaves {
				newAvailableCaves[key] = value
			}
			newVisitedSmallCaves := map[string]int{}
			for key, value := range visitedSmallCaves {
				newVisitedSmallCaves[key] = value
			}
			if strings.ToLower(cavePair[1]) == cavePair[1] && cavePair[1] != "end" {
				newVisitedSmallCaves[cavePair[1]]++
			}
			c.findPathsPart2(cavePair[1], cavePairs, newAvailableCaves, newVisitedSmallCaves)
		}
		if cavePair[1] == cave && availableCaves[cavePair[0]] {
			newAvailableCaves := map[string]bool{}
			for key, value := range availableCaves {
				newAvailableCaves[key] = value
			}
			newVisitedSmallCaves := map[string]int{}
			for key, value := range visitedSmallCaves {
				newVisitedSmallCaves[key] = value
			}
			if strings.ToLower(cavePair[0]) == cavePair[0] && cavePair[0] != "end" {
				newVisitedSmallCaves[cavePair[0]]++
			}
			c.findPathsPart2(cavePair[0], cavePairs, newAvailableCaves, newVisitedSmallCaves)
		}
	}
}

func readData(fileName string) ([][]string, map[string]bool) {
	cavePairs := [][]string{}
	availableCaves := map[string]bool{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringData := strings.Split(scanner.Text(), "-")
		cavePairs = append(cavePairs, stringData)
		availableCaves[stringData[0]] = true
		availableCaves[stringData[1]] = true
	}
	availableCaves["start"] = false
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cavePairs, availableCaves
}

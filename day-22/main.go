package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"math"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(switchRanges [][]int, onOrOff []string) int {
	switchedOn := map[[3]int]bool{}
	for i := 0; i < len(switchRanges); i++ {
		for x := int(math.Max(float64(switchRanges[i][0]), -50)); x <= int(math.Min(float64(switchRanges[i][1]), 50)); x++ {
			for y := int(math.Max(float64(switchRanges[i][2]), -50)); y <= int(math.Min(float64(switchRanges[i][3]), 50)); y++ {
				for z := int(math.Max(float64(switchRanges[i][4]), -50)); z <= int(math.Min(float64(switchRanges[i][5]), 50)); z++ {
					cube := [3]int{x, y, z}
					if onOrOff[i] == "on" {
						switchedOn[cube] = true
					} else {
						switchedOn[cube] = false
					}
				}
			}
		}
	}
	onCount := 0
	for _, value := range switchedOn {
		if value {
			onCount++
		}
	}
	return onCount
}

func part2(switchRanges [][]int, onOrOff []string) int {
	onCuboids := [][]int{}
	for index, switchRange := range switchRanges {
		newOnCuboids := [][]int{}
		for _, cuboid :=  range onCuboids {
			newCuboids := breakUpCuboid(switchRange, cuboid)
			newOnCuboids = append(newOnCuboids, newCuboids...)
		}
		onCuboids = newOnCuboids
		if onOrOff[index] == "on" {
			onCuboids = append(onCuboids, switchRange)
		}
	}
	onTotal := 0
	for _, cuboid := range onCuboids {
		onTotal += countCubes(cuboid)
	}
	return onTotal
}

func breakUpCuboid(new, old []int) ([][]int) {
	if new[0] > old[1] || new[1] < old[0] || new[2] > old[3] || new[3] < old[2] || new[4] > old[5] || new[5] < old[4] {
		return [][]int{old}
	} else if new[0] <= old[0] && new[1] >= old[1] && new[2] <= old[2] && new[3] >= old[3] && new[4] <= old[4] && new[5] >= old[5] {
		return [][]int{}
	} else {
		newCuboids := [][]int{}
		lowerX := old[0]
		upperX := old[1]
		lowerY := old[2]
		upperY := old[3]
		if old[0] < new[0] {
			newCuboids = append(newCuboids, []int{old[0], new[0]-1, old[2], old[3], old[4], old[5]})
			lowerX = new[0]
		}
		if old[1] > new[1] {
			newCuboids = append(newCuboids, []int{new[1]+1, old[1], old[2], old[3], old[4], old[5]})
			upperX = new[1]
		}
		if old[2] < new[2] {
			newCuboids = append(newCuboids, []int{lowerX, upperX, old[2], new[2]-1, old[4], old[5]})
			lowerY = new[2]
		}
		if old[3] > new[3] {
			newCuboids = append(newCuboids, []int{lowerX, upperX, new[3]+1, old[3], old[4], old[5]})
			upperY = new[3]
		}
		if old[4] < new[4] {
			newCuboids = append(newCuboids, []int{lowerX, upperX, lowerY, upperY, old[4], new[4]-1})
		}
		if old[5] > new[5] {
			newCuboids = append(newCuboids, []int{lowerX, upperX, lowerY, upperY, new[5]+1, old[5]})
		}
		return newCuboids
	}
}

func countCubes(cuboid []int) int {
	return (cuboid[1] - cuboid[0] + 1) * (cuboid[3] - cuboid[2] + 1) * (cuboid[5] - cuboid[4] + 1)
}

func readData(fileName string) ([][]int, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !unicode.IsNumber(c) && c != '-'
	}

	scanner := bufio.NewScanner(file)
	onOrOff := []string{}
	switchRanges := [][]int{}
	for scanner.Scan() {
		stringData := strings.Split(scanner.Text(), " ")
		onOrOff = append(onOrOff, stringData[0])
		switchRangeStrings := strings.FieldsFunc(stringData[1], splitCriteria)
		switchRange := []int{}
		for i := 0; i < len(switchRangeStrings); i++ {
			coordinate, err := strconv.Atoi(switchRangeStrings[i])
			if err != nil {
				log.Fatal(err)
			}
			switchRange = append(switchRange, coordinate)
		}
		switchRanges = append(switchRanges, switchRange)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return switchRanges, onOrOff
}

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

func part1(fuels []int) int {
	return totalFuelCost(fuels, true)
}

func part2(fuels []int) int {
	return totalFuelCost(fuels, false)
}

func totalFuelCost(fuels []int, constantRate bool) int {
	sort.Slice(fuels, func(i, j int) bool { return fuels[i] > fuels[j] })
	totalFuelCosts := []int{}
	for i := 0; i <= fuels[0]; i++ {
		fuelSum := 0
		for _, fuel := range fuels {
			posChange := int(math.Abs(float64(fuel - i)))
			fuelSum += fuelCost(posChange, constantRate)
		}
		totalFuelCosts = append(totalFuelCosts, fuelSum)
	}
	sort.Slice(totalFuelCosts, func(i, j int) bool { return totalFuelCosts[i] < totalFuelCosts[j] })
	return totalFuelCosts[0]
}

func fuelCost(posChange int, constantRate bool) int {
	if constantRate {
		return posChange
	}
	return posChange * (posChange + 1) / 2
}

func readData(fileName string) []int {
	fuels := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stringFuels := strings.Split(scanner.Text(), ",")
	for _, fuel := range stringFuels {
		i, err := strconv.Atoi(fuel)
		if err != nil {
			log.Fatal(err)
		}
		fuels = append(fuels, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fuels
}

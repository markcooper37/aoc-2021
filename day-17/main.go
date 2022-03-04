package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(xRange, yRange []int) int {
	xMin := int(math.Min(float64(xRange[0]), 0))
	xMax := int(math.Max(float64(xRange[1]), 0))
	yMin := int(-math.Abs(math.Max(math.Abs(float64(yRange[0])), (math.Abs(float64(yRange[1]))))))
	yMax := int(math.Abs(math.Max(math.Abs(float64(yRange[0])), (math.Abs(float64(yRange[1]))))))
	yHighest := yMin
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			xPosition := 0
			yPosition := 0 
			xVelocity := i
			yVelocity := j
			maxHeightReached := 0
			for k := 0; k < 3 * yMax + 1; k++ {
				xPosition += xVelocity
				yPosition += yVelocity
				if xVelocity > 0 {
					xVelocity--
				} else if xVelocity < 0 {
					xVelocity++
				}
				yVelocity--
				if yPosition > maxHeightReached {
					maxHeightReached = yPosition
				}
				if xPosition >= xRange[0] && xPosition <= xRange[1] && yPosition >= yRange[0] && yPosition <= yRange[1] {
					if maxHeightReached > yHighest {
						yHighest = maxHeightReached
					}
					break
				}
			} 
		}
	}
	return yHighest
}

func part2(xRange, yRange []int) int {
	xMin := int(math.Min(float64(xRange[0]), 0))
	xMax := int(math.Max(float64(xRange[1]), 0))
	yMin := int(-math.Abs(math.Max(math.Abs(float64(yRange[0])), (math.Abs(float64(yRange[1]))))))
	yMax := int(math.Abs(math.Max(math.Abs(float64(yRange[0])), (math.Abs(float64(yRange[1]))))))
	possibleVelocities := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			xPosition := 0
			yPosition := 0 
			xVelocity := i
			yVelocity := j
			for k := 0; k < 3 * yMax + 1; k++ {
				xPosition += xVelocity
				yPosition += yVelocity
				if xVelocity > 0 {
					xVelocity--
				} else if xVelocity < 0 {
					xVelocity++
				}
				yVelocity--
				if xPosition >= xRange[0] && xPosition <= xRange[1] && yPosition >= yRange[0] && yPosition <= yRange[1] {
					possibleVelocities++
					break
				}
			} 
		}
	}
	return possibleVelocities
}

func readData(fileName string) ([]int, []int) {
	xRange := []int{}
	yRange := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !(unicode.IsLetter(c) || unicode.IsNumber(c) || c == '-')
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stringData := strings.FieldsFunc(scanner.Text(), splitCriteria)
	xFirst, err := strconv.Atoi(stringData[3])
	if err != nil {
		log.Fatal(err)
	}
	xSecond, err := strconv.Atoi(stringData[4])
	if err != nil {
		log.Fatal(err)
	}
	yFirst, err := strconv.Atoi(stringData[6])
	if err != nil {
		log.Fatal(err)
	}
	ySecond, err := strconv.Atoi(stringData[7])
	if err != nil {
		log.Fatal(err)
	}
	xRange = append(xRange, xFirst, xSecond)
	yRange = append(yRange, yFirst, ySecond)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return xRange, yRange
}

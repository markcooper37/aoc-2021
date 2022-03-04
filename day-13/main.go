package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Fold struct {
	foldDirection string
	position      int
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(points map[[2]int]bool, folds []Fold) int {
	fold := folds[0]
	newPoints := map[[2]int]bool{}
	if fold.foldDirection == "horizontal" {
		for point := range points {
			if point[0] < fold.position {
				newPoints[point] = true
			} else {
				point[0] = 2*fold.position - point[0]
				newPoints[point] = true
			}
		}
	} else {
		for point := range points {
			if point[1] < fold.position {
				newPoints[point] = true
			} else {
				point[1] = 2*fold.position - point[1]
				newPoints[point] = true
			}
		}
	}
	return len(newPoints)
}

func part2(points map[[2]int]bool, folds []Fold) string {
	var horizonalSize int
	var verticalSize int
	for _, fold := range folds {
		newPoints := map[[2]int]bool{}
		if fold.foldDirection == "horizontal" {
			horizonalSize = fold.position
			for point := range points {
				if point[0] < fold.position {
					newPoints[point] = true
				} else {
					point[0] = 2*fold.position - point[0]
					newPoints[point] = true
				}
			}
		} else {
			verticalSize = fold.position
			for point := range points {
				if point[1] < fold.position {
					newPoints[point] = true
				} else {
					point[1] = 2*fold.position - point[1]
					newPoints[point] = true
				}
			}
		}
		points = newPoints
	}
	code := ""
	for i := 0; i < verticalSize; i++ {
		for j := 0; j < horizonalSize; j++ {
			point := [2]int{j, i}
			if _, ok := points[point]; ok {
				code = fmt.Sprintf("%s%s", code, "#")
			} else {
				code = fmt.Sprintf("%s%s", code, ".")
			}
			if j == horizonalSize - 1 {
				code = fmt.Sprintf("%s%s", code, "\n")
			}
		}
	}
	return code
}

func readData(fileName string) (map[[2]int]bool, []Fold) {
	coordinates := map[[2]int]bool{}
	folds := []Fold{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	splitCriteria := func(c rune) bool {
		return !(unicode.IsLetter(c) || unicode.IsNumber(c))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringData := strings.FieldsFunc(scanner.Text(), splitCriteria)
		if len(stringData) == 0 {
			continue
		}
		if firstCoordinate, err := strconv.Atoi(stringData[0]); err == nil {
			secondCoordinate, err := strconv.Atoi(stringData[1])
			if err != nil {
				log.Fatal(err)
			}
			numbers := [2]int{firstCoordinate, secondCoordinate}
			coordinates[numbers] = true
		} else {
			fold := Fold{}
			if stringData[2] == "x" {
				fold.foldDirection = "horizontal"
			} else {
				fold.foldDirection = "vertical"
			}
			position, err := strconv.Atoi(stringData[3])
			if err != nil {
				log.Fatal(err)
			}
			fold.position = position
			folds = append(folds, fold)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return coordinates, folds
}

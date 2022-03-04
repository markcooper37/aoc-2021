package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// for each input, the instructions can be simplified to the following steps:
// 1. we firstly check if z mod 26 plus some number1 equals the input
// 2. then, we divide z by 1 or 26
// 3. if step 1 was true, we keep z the same; otherwise, we multiply z by 26 and add the input and some number2

// there are 7 inputs where we divide z by 26 and 7 where we divide z by 1 in step 2
// by inspection, whenever we divide by 1 in step 2, step 1 is always false - the input is too small
// for each of these inputs, we are multiplying z by 26 (and adding something afterwards)
// thus, we must not subsequently multiply by 26 for the inputs where we divide z by 26
// the number we add on after multiplying by 26 is always less than 26 and therefore disappears when we divide by 26

// each multiplication by 26 for one input must be 'undone' by a division by 26 for a later input
// as a result, we obtain conditions on pairs of inputs that must be satisfied to ensure these divisions happen
// for example, we may have a div z 1 for input 3 and a div z 26 for input 4
// after step 3, z is equal to input3 + number2 mod 26, so we need that input3 + number2 + number1 equals input4 mod 26
// (here, number 2 is related to input 3 and number1 is related to input 4 as detailed in the simplified input steps)
// these pairs can be found by considering a stack - division by 26 'pops' the last multiplication by 26
// we maximise/minimise based on these pairings

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(divisions, firstAdds, secondAdds []int) int {
	indexPairs := findPairs(divisions)
	solution := maximiseSolution(indexPairs, firstAdds, secondAdds)
	return solution
}

func part2(divisions, firstAdds, secondAdds []int) int {
	indexPairs := findPairs(divisions)
	solution := minimiseSolution(indexPairs, firstAdds, secondAdds)
	return solution
}

func findPairs(divisions []int) [][]int {
	indices := []int{}
	indexPairs := [][]int{}
	for index, division := range divisions {
		if division == 1 {
			indices = append(indices, index)
		} else {
			indexPair := []int{indices[len(indices)-1], index}
			indexPairs = append(indexPairs, indexPair)
			indices = indices[:len(indices)-1]
		}
	}
	return indexPairs
}

func maximiseSolution(indexPairs [][]int, firstAdds, secondAdds []int) int {
	solutionSlice := [14]int{}
	for _, indexPair := range indexPairs {
		solutionFound := false
		for i := 9; i >= 1; i-- {
			if solutionFound {
				break
			}
			for j := 9; j >= 1; j-- {
				if solutionFound {
					break
				}
				if i+secondAdds[indexPair[0]]-j+firstAdds[indexPair[1]]%26 == 0 {
					solutionSlice[indexPair[0]] = i
					solutionSlice[indexPair[1]] = j
					solutionFound = true
				}
			}
		}
	}
	solution := 0
	for index, number := range solutionSlice {
		solution += number * int(math.Pow(10, float64(13-index)))
	}
	return solution
}

func minimiseSolution(indexPairs [][]int, firstAdds, secondAdds []int) int {
	solutionSlice := [14]int{}
	for _, indexPair := range indexPairs {
		solutionFound := false
		for i := 1; i <= 9; i++ {
			if solutionFound {
				break
			}
			for j := 1; j <= 9; j++ {
				if solutionFound {
					break
				}
				if i+secondAdds[indexPair[0]]-j+firstAdds[indexPair[1]]%26 == 0 {
					solutionSlice[indexPair[0]] = i
					solutionSlice[indexPair[1]] = j
					solutionFound = true
				}
			}
		}
	}
	solution := 0
	for index, number := range solutionSlice {
		solution += number * int(math.Pow(10, float64(13-index)))
	}
	return solution
}

func readData(fileName string) ([]int, []int, []int) {
	divisions := []int{}
	firstAdds := []int{}
	secondAdds := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		if line == 5 {
			instruction := strings.Split(scanner.Text(), " ")
			division, err := strconv.Atoi(instruction[2])
			if err != nil {
				log.Fatal(err)
			}
			divisions = append(divisions, division)
		} else if line == 6 {
			instruction := strings.Split(scanner.Text(), " ")
			firstAdd, err := strconv.Atoi(instruction[2])
			if err != nil {
				log.Fatal(err)
			}
			firstAdds = append(firstAdds, firstAdd)
		} else if line == 16 {
			instruction := strings.Split(scanner.Text(), " ")
			secondAdd, err := strconv.Atoi(instruction[2])
			if err != nil {
				log.Fatal(err)
			}
			secondAdds = append(secondAdds, secondAdd)
		} else if line == 18 {
			line = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return divisions, firstAdds, secondAdds
}

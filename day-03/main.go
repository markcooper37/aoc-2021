package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data := readData("input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data []string) int {
	gamma := 0
	epsilon := 0
	for i := 0; i < len(data[0]); i++ {
		zeroCount := 0
		oneCount := 0
		for j := 0; j < len(data); j++ {
			if data[j][i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		if zeroCount > oneCount {
			epsilon += 1 << (len(data[0]) - i - 1)
		} else {
			gamma += 1 << (len(data[0]) - i - 1)
		}
	}
	return gamma * epsilon
}


func part2(data []string) int {
	oxygen := data
	for i := 0; i < len(data[0]); i++ {
		newOxygen := []string{}
		zeroCount := 0
		oneCount := 0
		for j := 0; j < len(oxygen); j++ {
			if oxygen[j][i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		if zeroCount > oneCount {
			for k := 0; k < len(oxygen); k++ {
				if oxygen[k][i] == '0' {
					newOxygen = append(newOxygen, oxygen[k])
				}
			}
		} else {
			for k := 0; k < len(oxygen); k++ {
				if oxygen[k][i] == '1' {
					newOxygen = append(newOxygen, oxygen[k])
				}
			}
		}
		oxygen = newOxygen
		if len(oxygen) == 1 {
			break
		}
	}
	co2 := data
	for i := 0; i < len(data[0]); i++ {
		newCO2 := []string{}
		zeroCount := 0
		oneCount := 0
		for j := 0; j < len(co2); j++ {
			if co2[j][i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		if oneCount < zeroCount {
			for k := 0; k < len(co2); k++ {
				if co2[k][i] == '1' {
					newCO2 = append(newCO2, co2[k])
				}
			}
		} else  {
			for k := 0; k < len(co2); k++ {
				if co2[k][i] == '0' {
					newCO2 = append(newCO2, co2[k])
				}
			}
		}
		co2 = newCO2
		if len(co2) == 1 {
			break
		}
	}
	oxygenDecimal, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2Decimal, err := strconv.ParseInt(co2[0], 2, 64) 
	if err != nil {
		log.Fatal(err)
	}
	return int(co2Decimal * oxygenDecimal)
}

func readData(fileName string) []string {
	data := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		data = append(data, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
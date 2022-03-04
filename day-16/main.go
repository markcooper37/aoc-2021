package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
)

type VersionTotal struct {
	total int
	index int
}

type Index int

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(binaryDigits []string) int {
	v := VersionTotal{total: 0, index: 0}
	v.versionTotal(binaryDigits)
	return v.total
}

func part2(binaryDigits []string) int {
	var index Index
	index = 0
	transmissionTotal := index.transmissionTotal(binaryDigits)
	return transmissionTotal
}

func (v *VersionTotal) versionTotal(binaryDigits []string) {
	version, err := strconv.ParseInt(fmt.Sprintf("%s%s%s", binaryDigits[v.index], binaryDigits[v.index+1], binaryDigits[v.index+2]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	v.total += int(version)
	v.index += 3
	typeID, err := strconv.ParseInt(fmt.Sprintf("%s%s%s", binaryDigits[v.index], binaryDigits[v.index+1], binaryDigits[v.index+2]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	v.index += 3
	if int(typeID) == 4 {
		zeroFound := false
		for !zeroFound {
			if binaryDigits[v.index] == "0" {
				zeroFound = true
			}
			v.index += 5
		}
	} else {
		if binaryDigits[v.index] == "1" {
			v.index++
			numSubPacketsStr := ""
			for j := 0; j < 11; j++ {
				numSubPacketsStr = fmt.Sprintf("%s%s", numSubPacketsStr, binaryDigits[v.index+j])
			}
			v.index += 11
			numSubPackets, err := strconv.ParseInt(numSubPacketsStr, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			for j := 0; j < int(numSubPackets); j++ {
				v.versionTotal(binaryDigits)
			}
		} else {
			v.index++
			lenSubPacketsStr := ""
			for j := 0; j < 15; j++ {
				lenSubPacketsStr = fmt.Sprintf("%s%s", lenSubPacketsStr, binaryDigits[v.index+j])
			}
			v.index += 15
			lenSubPackets, err := strconv.ParseInt(lenSubPacketsStr, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			currentIndex := v.index
			for v.index < currentIndex+int(lenSubPackets) {
				v.versionTotal(binaryDigits)
			}
		}
	}
}

func (i *Index) transmissionTotal(binaryDigits []string) int {
	*i += 3
	typeIDStr, err := strconv.ParseInt(fmt.Sprintf("%s%s%s", binaryDigits[int(*i)], binaryDigits[int(*i)+1], binaryDigits[int(*i)+2]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	*i += 3
	typeID := int(typeIDStr)
	if typeID == 4 {
		zeroFound := false
		literalValue := ""
		for !zeroFound {
			if binaryDigits[int(*i)] == "0" {
				zeroFound = true
			}
			literalPart := fmt.Sprintf("%s%s%s%s", binaryDigits[int(*i)+1], binaryDigits[int(*i)+2], binaryDigits[int(*i)+3], binaryDigits[int(*i)+4])
			literalValue = fmt.Sprintf("%s%s", literalValue, literalPart)
			*i += 5
		}
		literalValueInt, err := strconv.ParseInt(literalValue, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		return int(literalValueInt)
	} else {
		if binaryDigits[*i] == "1" {
			*i++
			numSubPacketsStr := ""
			for j := 0; j < 11; j++ {
				numSubPacketsStr = fmt.Sprintf("%s%s", numSubPacketsStr, binaryDigits[int(*i)+j])
			}
			*i += 11
			numSubPackets, err := strconv.ParseInt(numSubPacketsStr, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			scores := []int{}
			for j := 0; j < int(numSubPackets); j++ {
				scores = append(scores, i.transmissionTotal(binaryDigits))
			}
			return packetScore(scores, typeID)
		} else {
			*i++
			lenSubPacketsStr := ""
			for j := 0; j < 15; j++ {
				lenSubPacketsStr = fmt.Sprintf("%s%s", lenSubPacketsStr, binaryDigits[int(*i)+j])
			}
			*i += 15
			lenSubPackets, err := strconv.ParseInt(lenSubPacketsStr, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			currentIndex := int(*i)
			scores := []int{}
			for int(*i) < currentIndex+int(lenSubPackets) {
				scores = append(scores, i.transmissionTotal(binaryDigits))
			}
			return packetScore(scores, typeID)
		}
	}
}

func packetScore(subpacketScores []int, typeID int) int {
	switch typeID {
	case 0:
		total := 0
		for _, score := range subpacketScores {
			total += score
		}
		return total
	case 1:
		total := 1
		for _, score := range subpacketScores {
			total *= score
		}
		return total
	case 2:
		sort.Ints(subpacketScores)
		return subpacketScores[0]
	case 3:
		sort.Ints(subpacketScores)
		return subpacketScores[len(subpacketScores)-1]
	case 5:
		if subpacketScores[0] > subpacketScores[1] {
			return 1
		}
		return 0
	case 6:
		if subpacketScores[0] < subpacketScores[1] {
			return 1
		}
		return 0
	case 7:
		if subpacketScores[0] == subpacketScores[1] {
			return 1
		}
		return 0
	}
	return 0
}

func readData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	binaryDigits := []string{}
	stringData := strings.Split(scanner.Text(), "")
	for _, string := range stringData {
		binaryDigit, err := strconv.ParseInt(string, 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		splitDigits := strings.Split(fmt.Sprintf("%04s", strconv.FormatInt(binaryDigit, 2)), "")
		for _, digit := range splitDigits {
			binaryDigits = append(binaryDigits, digit)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return binaryDigits
}

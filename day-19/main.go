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

type OceanScanner struct {
	beacons     [][]int
	coordinates []int
	number      int
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(oceanScanners []OceanScanner) int {
	oceanScanners = findScanners(oceanScanners)
	beaconsFound := countBeacons(oceanScanners)
	return beaconsFound
}

func part2(oceanScanners []OceanScanner) int {
	oceanScanners = findScanners(oceanScanners)
	scannerCoordinates := [][3]int{}
	for _, oceanScanner := range oceanScanners {
		newCoordinate := [3]int{}
		for j := 0; j <= 2; j++ {
			newCoordinate[j] = oceanScanner.coordinates[j]
		}
		scannerCoordinates = append(scannerCoordinates, newCoordinate)
	}
	maxManhattan := maxManhattan(scannerCoordinates)
	return maxManhattan
}

// find scanner positions relative to first scanner in list
func findScanners(oceanScanners []OceanScanner) []OceanScanner {
	oceanScanners[0].coordinates = []int{0, 0, 0}
	coordinatesFound := map[int]bool{0: true}
	for len(coordinatesFound) < len(oceanScanners) {
		for scannerNumber := range coordinatesFound {
			for i := 0; i < len(oceanScanners); i++ {
				if i == scannerNumber || coordinatesFound[i] {
					continue
				}
				newScanner, successful := scannerOverlap(oceanScanners[scannerNumber], oceanScanners[i])
				if successful {
					oceanScanners[i] = newScanner
					coordinatesFound[i] = true
				}
			}
		}
	}
	return oceanScanners
}

// check if two scanners have enough common beacons and if so, determine the position of the second scanner
func scannerOverlap(scanner1, scanner2 OceanScanner) (OceanScanner, bool) {
	successful := false
	for i := 0; i <= 23; i++ {
		newBeacons := [][]int{}
		for _, beacon := range scanner2.beacons {
			newBeacon := reorient(beacon, i)
			newBeacons = append(newBeacons, newBeacon)
		}
		for j := 0; j < len(scanner1.beacons); j++ {
			for k := 0; k < len(newBeacons); k++ {
				equalBeacons := 0
				xCoordinate := scanner1.beacons[j][0] - newBeacons[k][0]
				yCoordinate := scanner1.beacons[j][1] - newBeacons[k][1]
				zCoordinate := scanner1.beacons[j][2] - newBeacons[k][2]
				coordinate := []int{xCoordinate + scanner1.coordinates[0], yCoordinate + scanner1.coordinates[1], zCoordinate + scanner1.coordinates[2]}
				newNewBeacons := [][]int{}
				for _, beacon := range newBeacons {
					newNewBeacon := []int{xCoordinate + beacon[0], yCoordinate + beacon[1], zCoordinate + beacon[2]}
					newNewBeacons = append(newNewBeacons, newNewBeacon)
				}
				for _, beacon1 := range scanner1.beacons {
					for _, beacon2 := range newNewBeacons {
						if beacon1[0] == beacon2[0] && beacon1[1] == beacon2[1] && beacon1[2] == beacon2[2] {
							equalBeacons++
						}
					}
				}
				if equalBeacons >= 12 {
					scanner2.coordinates = coordinate
					scanner2.beacons = newBeacons
					successful = true
				}
			}
			if successful {
				break
			}
		}
		if successful {
			break
		}
	}
	if successful {
		return scanner2, true
	}
	return scanner2, false
}

// give the scanner one of the 24 different possible orientations
func reorient(coordinates []int, orientation int) []int {
	if orientation <= 3 {
		switch orientation {
		case 0:
			return coordinates
		case 1:
			return rotateAWiseAboutY(coordinates)
		case 2:
			return rotateAWiseAboutY(rotateAWiseAboutY(coordinates))
		case 3:
			return rotateCWiseAboutY(coordinates)
		}
	} else if orientation > 3 && orientation <= 7 {
		coordinates = rotateAWiseAboutX(coordinates)
		switch orientation {
		case 4:
			return coordinates
		case 5:
			return rotateAWiseAboutZ(coordinates)
		case 6:
			return rotateAWiseAboutZ(rotateAWiseAboutZ(coordinates))
		case 7:
			return rotateCWiseAboutZ(coordinates)
		}
	} else if orientation > 7 && orientation <= 11 {
		coordinates = rotateCWiseAboutX(coordinates)
		switch orientation {
		case 8:
			return coordinates
		case 9:
			return rotateAWiseAboutZ(coordinates)
		case 10:
			return rotateAWiseAboutZ(rotateAWiseAboutZ(coordinates))
		case 11:
			return rotateCWiseAboutZ(coordinates)
		}
	} else if orientation > 11 && orientation <= 15 {
		coordinates = rotateAWiseAboutZ(coordinates)
		switch orientation {
		case 12:
			return coordinates
		case 13:
			return rotateAWiseAboutX(coordinates)
		case 14:
			return rotateAWiseAboutX(rotateAWiseAboutX(coordinates))
		case 15:
			return rotateCWiseAboutX(coordinates)
		}
	} else if orientation > 15 && orientation <= 19 {
		coordinates = rotateCWiseAboutZ(coordinates)
		switch orientation {
		case 16:
			return coordinates
		case 17:
			return rotateAWiseAboutX(coordinates)
		case 18:
			return rotateAWiseAboutX(rotateAWiseAboutX(coordinates))
		case 19:
			return rotateCWiseAboutX(coordinates)
		}
	} else {
		coordinates = rotateAWiseAboutX(rotateAWiseAboutX(coordinates))
		switch orientation {
		case 20:
			return coordinates
		case 21:
			return rotateAWiseAboutY(coordinates)
		case 22:
			return rotateAWiseAboutY(rotateAWiseAboutY(coordinates))
		case 23:
			return rotateCWiseAboutY(coordinates)
		}
	}
	return []int{}
}

// functions for rotating scanners
// a clockwise rotation of the scanner is equivalent to an anticlockwise rotation of beacon positions and vice versa
func rotateCWiseAboutX(coordinates []int) []int {
	return []int{coordinates[0], coordinates[2], -coordinates[1]}
}

func rotateCWiseAboutY(coordinates []int) []int {
	return []int{-coordinates[2], coordinates[1], coordinates[0]}
}

func rotateCWiseAboutZ(coordinates []int) []int {
	return []int{coordinates[1], -coordinates[0], coordinates[2]}
}

func rotateAWiseAboutX(coordinates []int) []int {
	return []int{coordinates[0], -coordinates[2], coordinates[1]}
}

func rotateAWiseAboutY(coordinates []int) []int {
	return []int{coordinates[2], coordinates[1], -coordinates[0]}
}

func rotateAWiseAboutZ(coordinates []int) []int {
	return []int{-coordinates[1], coordinates[0], coordinates[2]}
}

// count the number of distinct beacons when all scanner positions are found
func countBeacons(oceanScanners []OceanScanner) int {
	beaconsFound := map[[3]int]bool{}
	for _, oceanScanner := range oceanScanners {
		for _, beacon := range oceanScanner.beacons {
			relativeCoordinates := [3]int{beacon[0] + oceanScanner.coordinates[0], beacon[1] + oceanScanner.coordinates[1], beacon[2] + oceanScanner.coordinates[2]}
			beaconsFound[relativeCoordinates] = true
		}
	}
	return len(beaconsFound)
}

// calculate the maximum Manhattan distance between pairs of scanners
func maxManhattan(coordinates [][3]int) int {
	maxManhattan := 0
	for i := 0; i < len(coordinates); i++ {
		for j := 0; j < len(coordinates); j++ {
			if i == j {
				continue
			}
			xDifference := int(math.Abs(float64(coordinates[i][0] - coordinates[j][0])))
			yDifference := int(math.Abs(float64(coordinates[i][1] - coordinates[j][1])))
			zDifference := int(math.Abs(float64(coordinates[i][2] - coordinates[j][2])))
			manhattan := xDifference + yDifference + zDifference
			if manhattan > maxManhattan {
				maxManhattan = manhattan
			}
		}
	}
	return maxManhattan
}

func readData(fileName string) []OceanScanner {
	oceanScanners := []OceanScanner{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	beacons := [][]int{}
	index := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			newOceanScanner := OceanScanner{beacons: beacons, number: index}
			oceanScanners = append(oceanScanners, newOceanScanner)
			beacons = [][]int{}
			index++
		} else if scanner.Text()[1] == '-' {
			continue
		} else {
			coordinateStrings := strings.Split(scanner.Text(), ",")
			coordinates := []int{}
			for i := 0; i <= 2; i++ {
				coordinate, err := strconv.Atoi(coordinateStrings[i])
				if err != nil {
					log.Fatal(err)
				}
				coordinates = append(coordinates, coordinate)
			}
			beacons = append(beacons, coordinates)
		}
	}
	newOceanScanner := OceanScanner{beacons: beacons, number: index}
	oceanScanners = append(oceanScanners, newOceanScanner)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return oceanScanners
}

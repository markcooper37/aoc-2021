package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Room struct {
	amphipod string
	column   int
	spaces   []string
}

func main() {
	fmt.Println(findLeastEnergy(readData("input.txt")))
	fmt.Println(findLeastEnergy(readData("input2.txt")))
}

func findLeastEnergy(hallway []string, rooms []Room) int {
	leastEnergy := 0
	minEnergy := &leastEnergy
	occupiable := []int{0, 1, 3, 5, 7, 9, 10}
	roomPositions := map[string]int{"A": 2, "B": 4, "C": 6, "D": 8}
	roomAvailable := []bool{false, false, false, false}
	letterToIndex := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}
	energies := map[string]int{"A": 1, "B": 10, "C": 100, "D": 1000}
	for roomIndex := range rooms {
		for _, occupiableSpace := range occupiable {
			makeMove(0, rooms, hallway, roomAvailable, roomPositions, occupiable, letterToIndex, energies, minEnergy, roomIndex, occupiableSpace)
		}
	}
	return leastEnergy
}

func isComplete(rooms []Room) bool {
	for _, room := range rooms {
		for _, space := range room.spaces {
			if space != room.amphipod {
				return false
			}
		}
	}
	return true
}

func isOccupiable(roomPosition, newPosition int, hallway []string, occupiable []int) bool {
	if hallway[newPosition] != "." {
		return false
	}
	if roomPosition > newPosition {
		for _, space := range occupiable {
			if roomPosition > space && space > newPosition {
				if hallway[space] != "." {
					return false
				}
			}
		}
		return true
	} else {
		for _, space := range occupiable {
			if roomPosition < space && space < newPosition {
				if hallway[space] != "." {
					return false
				}
			}
		}
		return true
	}
}

func canMoveToRoom(position int, amphipod string, roomAvailable []bool, roomPositions map[string]int, occupiable []int, hallway []string, letterToIndex map[string]int) bool {
	if !roomAvailable[letterToIndex[amphipod]] {
		return false
	}
	if position == roomPositions[amphipod] {
		return false
	}
	roomPosition := roomPositions[amphipod]
	if roomPosition > position {
		for _, space := range occupiable {
			if roomPosition > space && space > position {
				if hallway[space] != "." {
					return false
				}
			}
		}
		return true
	} else {
		for _, space := range occupiable {
			if roomPosition < space && space < position {
				if hallway[space] != "." {
					return false
				}
			}
		}
		return true
	}
}

func isRoomAvailable(room Room) bool {
	if room.spaces[0] != "." {
		return false
	}
	for _, space := range room.spaces {
		if space != room.amphipod && space != "." {
			return false
		}
	}
	return true
}

func isInCorrectPosition(rooms []Room, roomIndex int) bool {
	space := firstIndex(rooms[roomIndex])
	amphipod := rooms[roomIndex].spaces[space]
	if amphipod != rooms[roomIndex].amphipod {
		return false
	}
	for index := space + 1; index < len(rooms[roomIndex].spaces); index++ {
		if rooms[roomIndex].spaces[index] != rooms[roomIndex].amphipod {
			return false
		}
	}
	return true
}

func firstIndex(room Room) int {
	for index := range room.spaces {
		if room.spaces[index] != "." {
			return index
		}
	}
	return len(room.spaces)
}

func moveHallwayToRoom(rooms []Room, amphipod string, roomIndex, currentPosition int, hallway []string, total int, energies map[string]int) (int, []Room, []string) {
	spaceInRoom := firstIndex(rooms[roomIndex]) - 1
	rooms[roomIndex].spaces[spaceInRoom] = amphipod
	hallway[currentPosition] = "."
	spacesMoved := spaceInRoom + 1 + int(math.Abs(float64(rooms[roomIndex].column-currentPosition)))
	total += spacesMoved * energies[amphipod]
	return total, rooms, hallway
}

func moveRoomToRoom(rooms []Room, oldRoomIndex, newRoomIndex int, amphipod string, total int, energies map[string]int) (int, []Room) {
	moveFromIndex := firstIndex(rooms[oldRoomIndex])
	moveToIndex := firstIndex(rooms[newRoomIndex]) - 1
	rooms[oldRoomIndex].spaces[moveFromIndex] = "."
	rooms[newRoomIndex].spaces[moveToIndex] = amphipod
	spacesMoved := moveFromIndex + 1 + moveToIndex + 1 + int(math.Abs(float64(rooms[oldRoomIndex].column-rooms[newRoomIndex].column)))
	total += spacesMoved * energies[amphipod]
	return total, rooms
}

func moveToHallway(rooms []Room, roomIndex int, amphipod string, hallway []string, hallwayIndex int, total int, energies map[string]int) (int, []Room, []string) {
	hallway[hallwayIndex] = amphipod
	rooms[roomIndex].spaces[firstIndex(rooms[roomIndex])] = "."
	spacesMoves := firstIndex(rooms[roomIndex]) + int(math.Abs(float64(rooms[roomIndex].column-hallwayIndex)))
	total += spacesMoves * energies[amphipod]
	return total, rooms, hallway
}

func makeMove(total int, rooms []Room, hallway []string, roomAvailable []bool, roomPositions map[string]int, occupiable []int, letterToIndex map[string]int, energies map[string]int, minEnergy *int, roomIndex int, hallwayPosition int) {
	newRoomAvailable := []bool{}
	for _, value := range roomAvailable {
		newRoomAvailable = append(newRoomAvailable, value)
	}
	newHallway := []string{}
	for _, value := range hallway {
		newHallway = append(newHallway, value)
	}
	newRooms := []Room{}
	for _, room := range rooms {
		newRoom := Room{}
		newSpaces := []string{}
		for _, space := range room.spaces {
			newSpaces = append(newSpaces, space)
		}
		newRoom.spaces = newSpaces
		newRoom.amphipod = room.amphipod
		newRoom.column = room.column
		newRooms = append(newRooms, newRoom)
	}
	if firstIndex(newRooms[roomIndex]) == len(newRooms[roomIndex].spaces) {
		return
	}
	if isInCorrectPosition(newRooms, roomIndex) {
		return
	}
	if isOccupiable(newRooms[roomIndex].column, hallwayPosition, newHallway, occupiable) {
		total, newRooms, newHallway = moveToHallway(newRooms, roomIndex, rooms[roomIndex].spaces[firstIndex(newRooms[roomIndex])], newHallway, hallwayPosition, total, energies)
	} else {
		return
	}
	for _, room := range newRooms {
		newRoomAvailable[letterToIndex[room.amphipod]] = isRoomAvailable(room)
	}
	movedToRoom := true
	for movedToRoom {
		movedToRoom = false
		for position, space := range newHallway {
			if space == "." {
				continue
			} else {
				if canMoveToRoom(position, space, newRoomAvailable, roomPositions, occupiable, newHallway, letterToIndex) {
					total, newRooms, newHallway = moveHallwayToRoom(newRooms, space, letterToIndex[space], position, newHallway, total, energies)
					for _, room := range newRooms {
						newRoomAvailable[letterToIndex[room.amphipod]] = isRoomAvailable(room)
					}
					movedToRoom = true
				}
			}
		}
		for i, room := range newRooms {
			if firstIndex(room) == len(room.spaces) {
				continue
			}
			if canMoveToRoom(room.column, room.spaces[firstIndex(room)], newRoomAvailable, roomPositions, occupiable, newHallway, letterToIndex) {
				total, newRooms = moveRoomToRoom(newRooms, i, letterToIndex[room.spaces[firstIndex(room)]], room.spaces[firstIndex(room)], total, energies)
				for _, room := range newRooms {
					newRoomAvailable[letterToIndex[room.amphipod]] = isRoomAvailable(room)
				}
				movedToRoom = true
				break
			}
		}
	}
	if isComplete(newRooms) {
		if *minEnergy == 0 {
			*minEnergy = total
		}
		if total < *minEnergy {
			*minEnergy = total
		}
		return
	}
	for index := range newRooms {
		for _, space := range occupiable {
			makeMove(total, newRooms, newHallway, newRoomAvailable, roomPositions, occupiable, letterToIndex, energies, minEnergy, index, space)
		}
	}
}

func readData(fileName string) ([]string, []Room) {
	hallway := []string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."}
	rooms := []Room{{column: 2, amphipod: "A", spaces: []string{}}, {column: 4, amphipod: "B", spaces: []string{}}, {column: 6, amphipod: "C", spaces: []string{}}, {column: 8, amphipod: "D", spaces: []string{}}}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	index := 1
	for scanner.Scan() {
		stringData := strings.Split(scanner.Text(), "")
		if stringData[3] == "#" {
			continue
		}
		for i := 0; i <= 3; i++ {
			rooms[i].spaces = append(rooms[i].spaces, stringData[2*i+3])
		}
		index++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return hallway, rooms
}

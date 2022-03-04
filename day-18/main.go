package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
	parent     *Node
	depth      int
}

type Tree struct {
	root *Node
}

func main() {
	fmt.Println(part1(readData("input.txt")))
	fmt.Println(part2(readData("input.txt")))
}

func part1(numbers []string) int {
	trees := []Tree{}
	for _, number := range numbers {
		trees = append(trees, createTree(number))
	}
	tree := trees[0]
	for i := 1; i <= len(trees)-1; i++ {
		newTree := addTrees(tree, trees[i])
		newTree = reduceTree(newTree)
		tree = newTree
	}
	return tree.root.magnitude()
}

func part2(numbers []string) int {
	biggestMagnitude := 0
	for i := 0; i <= len(numbers)-1; i++ {
		for j := 0; j <= len(numbers)-1; j++ {
			if i == j {
				continue
			}
			tree1 := createTree(numbers[i])
			tree2 := createTree(numbers[j])
			newTree := addTrees(tree1, tree2)
			newTree = reduceTree(newTree)
			newTreeMagnitude := newTree.root.magnitude()
			if newTreeMagnitude > biggestMagnitude {
				biggestMagnitude = newTree.root.magnitude()
			}
		}
	}
	return biggestMagnitude
}

func createTree(input string) Tree {
	root := &Node{depth: 0}
	currentNode := root
	newTree := Tree{root: root}
	for _, char := range input {
		if char == '[' {
			currentNode.leftChild = &Node{parent: currentNode, depth: currentNode.depth + 1}
			currentNode = currentNode.leftChild
		} else if char == ']' {
			currentNode = currentNode.parent
		} else if char == ',' {
			currentNode = currentNode.parent
			currentNode.rightChild = &Node{parent: currentNode, depth: currentNode.depth + 1}
			currentNode = currentNode.rightChild
		} else {
			currentNode.value = int(char) - 48
		}
	}
	return newTree
}

func addTrees(tree1, tree2 Tree) Tree {
	tree1.root.incrementDepths()
	tree2.root.incrementDepths()
	newRoot := Node{depth: 0, leftChild: tree1.root, rightChild: tree2.root}
	newRoot.leftChild.parent = &newRoot
	newRoot.rightChild.parent = &newRoot
	newTree := Tree{root: &newRoot}
	return newTree
}

func reduceTree(tree Tree) Tree {
	checksNeeded := true
	for checksNeeded {
		tree.root.explodeNodes()
		splitOccured := tree.root.splitNode()
		if !splitOccured {
			checksNeeded = false
		}
	}
	return tree
}

func (n *Node) incrementDepths() {
	n.depth += 1
	if n.leftChild != nil {
		n.leftChild.incrementDepths()
	}
	if n.rightChild != nil {
		n.rightChild.incrementDepths()
	}
}

func (n *Node) explodeNodes() {
	if n.depth == 4 && n.leftChild != nil && n.rightChild != nil {
		leftValue := n.leftChild.value
		n.leftChild = nil
		leftReached := false
		currentNode := n
		for !leftReached {
			if currentNode.parent == nil {
				leftReached = true
				continue
			}
			if currentNode.parent.leftChild == currentNode {
				currentNode = currentNode.parent
				continue
			}
			currentNode = currentNode.parent.leftChild
			leftValueUpdated := false
			for !leftValueUpdated {
				if currentNode.rightChild != nil {
					currentNode = currentNode.rightChild
				} else {
					currentNode.value += leftValue
					leftValueUpdated = true
				}
			}
			leftReached = true
		}
		rightValue := n.rightChild.value
		n.rightChild = nil
		rightReached := false
		currentNode = n
		for !rightReached {
			if currentNode.parent == nil {
				rightReached = true
				continue
			}
			if currentNode.parent.rightChild == currentNode {
				currentNode = currentNode.parent
				continue
			}
			currentNode = currentNode.parent.rightChild
			rightValueUpdated := false
			for !rightValueUpdated {
				if currentNode.leftChild != nil {
					currentNode = currentNode.leftChild
				} else {
					currentNode.value += rightValue
					rightValueUpdated = true
				}
			}
			rightReached = true
		}
		n.value = 0
		return
	}
	if n.leftChild != nil {
		n.leftChild.explodeNodes()
	}
	if n.rightChild != nil {
		n.rightChild.explodeNodes()
	}
}

func (n *Node) splitNode() bool {
	if n.value >= 10 {
		n.leftChild = &Node{depth: n.depth + 1, value: n.value / 2, parent: n}
		n.rightChild = &Node{depth: n.depth + 1, value: (n.value + 1) / 2, parent: n}
		n.value = 0
		return true
	}
	if n.leftChild != nil {
		leftChildSplit := n.leftChild.splitNode()
		if leftChildSplit {
			return true
		}
	}
	if n.rightChild != nil {
		rightChildSplit := n.rightChild.splitNode()
		if rightChildSplit {
			return true
		}
	}
	return false
}

func (n *Node) magnitude() int {
	if n.leftChild == nil || n.rightChild == nil {
		return n.value
	} else {
		return 3*n.leftChild.magnitude() + 2*n.rightChild.magnitude()
	}
}

func readData(fileName string) []string {
	numbers := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers
}

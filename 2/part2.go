package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Tree struct {
	value    rune
	children map[rune]*Tree
}

func main() {
	t := time.Now()
	boxIds, err := readBoxIds("input.txt")
	if err != nil {
		panic(err)
	}
	root := &Tree{}
	for _, boxId := range boxIds {
		substr, t := root.MaxSubstring(boxId)
		for _, child := range t.children {
			if child.Contains(boxId[len(substr)+1:]) {
				fmt.Println(boxId)
				fmt.Printf("%s%c...\n", substr, child.value)
			}
		}
		root.Insert(boxId)
	}
	fmt.Println(time.Since(t))
}

func (t *Tree) Insert(s string) {
	// t is the root
	node := t
	for _, c := range s {
		if child, ok := node.children[c]; ok {
			// this rune is in the tree, advance the node pointer
			node = child
			continue
		}
		// not in the tree, create the child node
		if node.children == nil {
			node.children = make(map[rune]*Tree)
		}
		child := &Tree{value: c}
		node.children[c] = child
		node = child
	}
}

func (t *Tree) MaxSubstring(s string) (string, *Tree) {
	node := t
	var substr string
	for _, c := range s {
		if child, ok := node.children[c]; ok {
			node = child
			substr += string(c)
			continue
		}
		return substr, node
	}
	return substr, node
}

func (t *Tree) Contains(s string) bool {
	node := t
	for _, c := range s {
		if child, ok := node.children[c]; ok {
			node = child
			continue
		}
		return false
	}
	return true
}

func readBoxIds(file string) ([]string, error) {
	input, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var boxIds []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		boxIds = append(boxIds, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return boxIds, nil
}

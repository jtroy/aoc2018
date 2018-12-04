package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	boxIds, err := readBoxIds("input.txt")
	if err != nil {
		panic(err)
	}
	var (
		twos   int
		threes int
	)
	for _, boxId := range boxIds {
		h := letterHistogram(boxId)
		if hasExactly(h, 2) {
			twos++
		}
		if hasExactly(h, 3) {
			threes++
		}
	}
	fmt.Println(twos * threes)
}

func hasExactly(histogram map[rune]int, n int) bool {
	for _, count := range histogram {
		if count == n {
			return true
		}
	}
	return false
}

func letterHistogram(s string) map[rune]int {
	histogram := make(map[rune]int)

	for _, c := range s {
		histogram[c]++
	}
	return histogram
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	offsets, err := readOffsets("input.txt")
	if err != nil {
		panic(err)
	}

	visited := make(map[int64]struct{})
	var (
		i    int
		freq int64
	)
	for {
		if i == len(offsets) {
			i = 0
		}
		if _, seen := visited[freq]; seen {
			break
		}
		visited[freq] = struct{}{}
		freq += offsets[i]
		i++
	}

	fmt.Printf("Found repeat frequency %d\n", freq)
}

func readOffsets(file string) ([]int64, error) {
	input, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var offsets []int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		off, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		offsets = append(offsets, off)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return offsets, nil
}

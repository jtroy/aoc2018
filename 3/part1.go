package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	Id int
	X  int
	Y  int
	W  int
	H  int
}

func main() {
	claims, maxX, maxY, err := parseClaims("input.txt")
	if err != nil {
		panic(err)
	}
	fabric := make([][]int, maxX)
	for i := range fabric {
		fabric[i] = make([]int, maxY)
	}
	for _, claim := range claims {
		for i := claim.X; i < claim.X+claim.W; i++ {
			for j := claim.Y; j < claim.Y+claim.H; j++ {
				fabric[i][j]++
			}
		}
	}
	var contested int
	for _, strip := range fabric {
		for _, inch := range strip {
			if inch > 1 {
				contested++
			}
		}
	}
	fmt.Println(contested)
}

func parseClaims(file string) ([]Claim, int, int, error) {
	input, err := os.Open(file)
	if err != nil {
		return nil, 0, 0, err
	}
	defer input.Close()

	var (
		claims []Claim
		maxX   int
		maxY   int
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var (
			err   error
			claim Claim
		)

		// #1373 @ 808,590: 13x18
		fields := strings.Fields(scanner.Text())
		claim.Id, err = strconv.Atoi(fields[0][1:]) // #1373
		if err != nil {
			return nil, 0, 0, err
		}
		// @
		xy := strings.Split(strings.TrimRight(fields[2], ":"), ",") // 808,590:
		claim.X, err = strconv.Atoi(xy[0])
		claim.Y, err = strconv.Atoi(xy[1])
		wh := strings.Split(fields[3], "x") // 13x18
		claim.W, err = strconv.Atoi(wh[0])
		claim.H, err = strconv.Atoi(wh[1])

		edgeX := claim.X + claim.W
		edgeY := claim.Y + claim.H

		if edgeX > maxX {
			maxX = edgeX
		}
		if edgeY > maxY {
			maxY = edgeY
		}

		claims = append(claims, claim)
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	return claims, maxX, maxY, nil
}

package day2

import (
	"fmt"
	"os"
	"strings"
)

func part1(puzzleInput string) int {
	lines := strings.Split(puzzleInput, "\n")

	totalStore := 0
	for _, line := range lines {
		roundScore := 0
		shapes := strings.Split(line, " ")
		opponentsShape, yourShape := shapes[0], shapes[1]

		// What's your shape
		switch yourShape {
		case "X":
			roundScore += 1
		case "Y":
			roundScore += 2
		case "Z":
			roundScore += 3
		}
		// Who wins
		if (yourShape == "X" && opponentsShape == "A") || (yourShape == "Y" && opponentsShape == "B") || (yourShape == "Z" && opponentsShape == "C") {
			roundScore += 3
		} else if (yourShape == "X" && opponentsShape == "C") || (yourShape == "Y" && opponentsShape == "A") || (yourShape == "Z" && opponentsShape == "B") {
			roundScore += 6
		}

		totalStore += roundScore
	}

	return totalStore
}

func part2(puzzleInput string) int {
	lines := strings.Split(puzzleInput, "\n")

	totalStore := 0
	for _, line := range lines {
		roundScore := 0
		shapes := strings.Split(line, " ")
		opponentsShape, yourShape := shapes[0], shapes[1]

		// What's your shape
		if (yourShape == "X" && opponentsShape == "B") || (yourShape == "Y" && opponentsShape == "A") || (yourShape == "Z" && opponentsShape == "C") {
			roundScore += 1
		} else if (yourShape == "X" && opponentsShape == "C") || (yourShape == "Y" && opponentsShape == "B") || (yourShape == "Z" && opponentsShape == "A") {
			roundScore += 2
		} else if (yourShape == "X" && opponentsShape == "A") || (yourShape == "Y" && opponentsShape == "C") || (yourShape == "Z" && opponentsShape == "B") {
			roundScore += 3
		}
		// Who wins
		switch yourShape {
		case "X":
			roundScore += 0
		case "Y":
			roundScore += 3
		case "Z":
			roundScore += 6
		}

		totalStore += roundScore
	}

	return totalStore
}

func Day2() {
	puzzleRead, err := os.ReadFile("../Input/day2.txt")
	if err != nil {
		fmt.Println(err)
	}
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 2 Part 1 result is: %v\n", part1(puzzleCleaned))
	fmt.Printf("Day 2 Part 2 result is: %v\n", part2(puzzleCleaned))
}

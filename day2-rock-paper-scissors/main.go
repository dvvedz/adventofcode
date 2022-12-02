package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := openFile("hands.txt")
	var hands []Hand = parseFile(file)

	var scoreCounter1 int
	var scoreCounter2 int

	for _, hand := range hands {
		scoreCounter1 += hand.playStratOne()
		scoreCounter2 += hand.playStratTwo()
	}

	fmt.Println(scoreCounter1)
	fmt.Println(scoreCounter2)
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("error opening file", err)
	}

	return f
}

func parseFile(f *os.File) []Hand {
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		getRow := strings.Split(line, "\n")
		getValCol := strings.Split(getRow[0], " ")

		var val int

		var val2 int

		switch getValCol[0] {
		case "A":
			val = 1
		case "B":
			val = 2
		case "C":
			val = 3
		}

		switch getValCol[1] {
		case "X":
			val2 = 1
		case "Y":
			val2 = 2
		case "Z":
			val2 = 3
		}

		myM := map[string]int{getValCol[1]: val2}
		opM := map[string]int{getValCol[0]: val}

		h := Hand{
			my:       myM,
			opponent: opM,
		}

		hands = append(hands, h)
	}
	return hands
}

type Hand struct {
	my       map[string]int
	opponent map[string]int
}

func (h Hand) playStratTwo() int {
	// X means i have to LOOSE
	// Y means i have to DRAW
	// Z means i have to WIN

	var score int

	if h.my["X"] == 1 {
		// LOOSE
		if h.opponent["A"] == 1 {
			score = 0 + 3
		} else if h.opponent["B"] == 2 {
			score = 0 + 1
		} else if h.opponent["C"] == 3 {
			score = 0 + 2
		}

	} else if h.my["Y"] == 2 {
		// DRAW
		if h.opponent["A"] == 1 {
			score = 1 + 3
		} else if h.opponent["B"] == 2 {
			score = 2 + 3
		} else if h.opponent["C"] == 3 {
			score = 3 + 3
		}
	} else if h.my["Z"] == 3 {
		// WIN
		if h.opponent["A"] == 1 {
			score = 2 + 6
		} else if h.opponent["B"] == 2 {
			score = 3 + 6
		} else if h.opponent["C"] == 3 {
			score = 1 + 6
		}
	}
	return score
}

func (h Hand) playStratOne() int {
	// My Values:
	// X (ROCK) 	= 1
	// Y (Paper) 	= 2
	// Z (Scissors) = 3

	// Opponenet Values:
	// A (ROCK) 	= 1
	// B (Paper) 	= 2
	// C (Scissors) = 3

	// A WIN is worth 	= 6
	// A DRAW is worth	= 3
	// A LOSS is worth 	= 0

	// Counting Scores
	var score int

	if h.my["X"] == 1 && h.opponent["A"] == 1 {
		// ROCK VS ROCK = DRAW
		score = 3 + 1
	} else if h.my["X"] == 1 && h.opponent["B"] == 2 {
		// ROCK VS PAPER = LOSS
		score = 1
	} else if h.my["X"] == 1 && h.opponent["C"] == 3 {
		// ROCK VS SCISSORS = WIN
		score = 1 + 6
	} else if h.my["Y"] == 2 && h.opponent["A"] == 1 {
		// PAPER VS ROCK = WIN
		score = 2 + 6
	} else if h.my["Y"] == 2 && h.opponent["B"] == 2 {
		// PAPER VS PAPER = DRAW
		score = 2 + 3
	} else if h.my["Y"] == 2 && h.opponent["C"] == 3 {
		// PAPER VS SCISSORS = LOSS
		score = 2
	} else if h.my["Z"] == 3 && h.opponent["A"] == 1 {
		// SCISSORS VS ROCK = LOSS
		score = 3
	} else if h.my["Z"] == 3 && h.opponent["B"] == 2 {
		// SCISSORS VS PAPER = WIN
		score = 3 + 6
	} else if h.my["Z"] == 3 && h.opponent["C"] == 3 {
		// SCISSORS VS SCISSORS = DRAW
		score = 3 + 3
	}

	return score
}

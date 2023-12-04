package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var playCards [][]int
var winningCards [][]int

func main() {
	playCards = make([][]int, 0)
	winningCards = make([][]int, 0)

	var input *[]string
	input = ReadInput()
	ParseInput(input)
	PartOne()
}

func PartOne() {
	var totalPoints int
	for game, cards := range playCards {
		var points = 0
		match := 0
		for _, card := range cards {
			if slices.Contains(winningCards[game], card) {
				match++
				if match == 1 {
					points = 1
				} else {
					points *= 2
				}

			}
		}
		totalPoints += points
	}
	fmt.Printf("Total points: %d\r\n", totalPoints)
}

func ReadInput() *[]string {
	var input = make([]string, 0)

	inputFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return &input
}

func ParseInput(input *[]string) {
	for _, currentLine := range *input {
		currentLineSplit := strings.Split(currentLine, "|")

		// Parse winning numbers
		playingCardsLine := make([]int, 0)
		for _, item := range strings.Split(currentLineSplit[1], " ") {
			if item != " " && item != "" {
				value, _ := strconv.ParseInt(item, 10, 64)
				playingCardsLine = append(playingCardsLine, int(value))
			}
		}
		playCards = append(playCards, playingCardsLine)

		// Parse playing cards
		winningCardsLine := make([]int, 0)
		playCardData := strings.Split(currentLineSplit[0], ":")
		for _, item := range strings.Split(playCardData[1], " ") {

			if item != " " && item != "" {
				value, _ := strconv.ParseInt(item, 10, 64)
				winningCardsLine = append(winningCardsLine, int(value))
			}

		}
		winningCards = append(winningCards, winningCardsLine)
	}
}

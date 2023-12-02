package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	Id    int64
	Blue  []int
	Red   []int
	Green []int
}

const MaxRed int = 12
const MaxGreen int = 13
const MaxBlue int = 14

func main() {
	var input *[]string
	var games *[]Game
	fmt.Print("Read input\r\n")
	input = ReadInput()
	games = ConvertInput(input)

	var total int64

	for _, game := range *games {
		sort.Ints(game.Red)
		maxRed := game.Red[len(game.Red)-1]

		sort.Ints(game.Blue)
		maxBlue := game.Blue[len(game.Blue)-1]

		sort.Ints(game.Green)
		maxGreen := game.Green[len(game.Green)-1]

		if maxRed <= MaxRed && maxBlue <= MaxBlue && maxGreen <= MaxGreen {
			total += game.Id
		}
	}
	fmt.Println(total)
}

func ReadInput() *[]string {
	var input = make([]string, 0)

	// Read the input from file
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

func ConvertInput(input *[]string) *[]Game {
	var games = make([]Game, 0)

	for _, rawData := range *input {
		var convertedGame Game

		game := strings.Split(rawData, ":")

		// Get the game id
		idSplit := strings.Split(game[0], " ")
		number, _ := strconv.ParseInt(idSplit[1], 10, 64)
		convertedGame.Id = number

		cubeData := strings.Split(game[1], ";")
		for _, cubes := range cubeData {
			cubesSplit := strings.Split(cubes, ",")

			for _, colors := range cubesSplit {
				colorData := strings.Split(colors, " ")
				amount, _ := strconv.ParseInt(colorData[1], 10, 64)

				color := colorData[2]

				switch color {
				case "red":
					convertedGame.Red = append(convertedGame.Red, int(amount))
				case "blue":
					convertedGame.Blue = append(convertedGame.Blue, int(amount))
				case "green":
					convertedGame.Green = append(convertedGame.Green, int(amount))
				}
			}
		}

		games = append(games, convertedGame)
	}

	return &games
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var timeDistance = make(map[int]int)

func main() {
	var input *[]string
	input = ReadInput()
	ParseInput(input)
	PartOne()
}

func PartOne() {
	sNew := 0
	total := 0
	totalSolutions := make([]int, 0)

	for t, s := range timeDistance {
		possibleSolutions := 0
		for i := 0; i < t; i++ {
			v := t - (t - i)
			sNew = v * (t - i)
			if sNew > s {
				possibleSolutions++
			}
		}
		totalSolutions = append(totalSolutions, possibleSolutions)
	}

	for i := 0; i < len(totalSolutions); i++ {
		if i == 0 {
			total = totalSolutions[i]
			continue
		}
		total *= totalSolutions[i]
	}

	fmt.Println(total)
}

func ParseInput(input *[]string) {
	time := make([]int, 0)
	distance := make([]int, 0)

	for line, currentLine := range *input {
		currentLineSplit := strings.Split(currentLine, ":")[1]
		currentLineSplit = strings.Replace(currentLineSplit, " ", ";", -1)
		for _, number := range strings.Split(currentLineSplit, ";") {
			if number == "" {
				continue
			}
			numberToFloat, _ := strconv.ParseInt(number, 10, 64)
			if line == 0 {
				time = append(time, int(numberToFloat))
			}
			if line == 1 {
				distance = append(distance, int(numberToFloat))
			}
		}
	}
	for i := 0; i < len(time); i++ {
		timeDistance[time[i]] = distance[i]
	}
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

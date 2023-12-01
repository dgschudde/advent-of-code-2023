package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var input *[]string

	fmt.Print("Read input\r\n")
	input = ReadInput()
	fmt.Printf("Read %d lines\r\n", len(*input))
	FilterDigits(input)
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

func FilterDigits(input *[]string) {
	var total int64

	var mapping = map[string]string{
		"one":   "o1ne",
		"two":   "t2wo",
		"three": "th3ree",
		"four":  "fo4ur",
		"five":  "fi5ve",
		"six":   "si6x",
		"seven": "sev7en",
		"eight": "eig8ht",
		"nine":  "ni9ne",
	}

	var expression = regexp.MustCompile("[0-9]")

	for _, inputLine := range *input {
		for k, v := range mapping {
			inputLine = strings.Replace(inputLine, k, v, -1)
		}

		convertedLine := expression.FindAllString(inputLine, -1)

		outerDigits := convertedLine[0] + convertedLine[len(convertedLine)-1]
		decimalValue, _ := strconv.ParseInt(outerDigits, 10, 64)
		total += decimalValue
	}

	fmt.Println(total)
}

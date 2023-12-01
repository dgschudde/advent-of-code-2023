package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	expression := regexp.MustCompile("[0-9]")

	var total int64

	for _, value := range *input {
		digits := expression.FindAllString(value, -1)
		outerDigits := digits[0] + digits[len(digits)-1]
		decimalValue, _ := strconv.ParseInt(outerDigits, 10, 64)
		total += decimalValue
	}

	fmt.Println(total)
}

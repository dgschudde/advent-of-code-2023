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
	var input []string

	input = ReadInput()

	var expression = regexp.MustCompile("[0-9]+")
	var total int

	for line, currentLine := range input {
		matches := expression.FindAllStringIndex(currentLine, -1)

		// Iterate matches
		for _, match := range matches {
			left := match[0]
			right := match[1]

			isValid := true

			// Check above
			if line > 0 {
				above := input[line-1]
				check := above[left:right]
				count := strings.Count(check, ".")

				if count != (right - left) {
					isValid = false
				}

				if left > 0 {
					if above[left-1:left] != "." {
						isValid = false
					}
				}

				if right < len(currentLine) {
					if above[right:right+1] != "." {
						isValid = false
					}
				}
			}

			// Check below
			if line < len(input)-1 {
				below := input[line+1]
				check := below[left:right]
				count := strings.Count(check, ".")

				if count != (right - left) {
					isValid = false
				}

				if left > 0 {
					if below[left-1:left] != "." {
						isValid = false
					}
				}

				if right < len(currentLine) {
					if below[right:right+1] != "." {
						isValid = false
					}
				}
			}

			// Check left
			if left > 0 {
				if currentLine[left-1:left] != "." {
					isValid = false
				}
			}

			// Check right
			if right < len(currentLine) {
				if currentLine[right:right+1] != "." {
					isValid = false
				}
			}

			if isValid == false {
				value := currentLine[left:right]
				number, _ := strconv.ParseInt(value, 10, 64)
				total += int(number)
			}
		}
	}
	fmt.Println(total)
}

func ReadInput() []string {
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

	return input
}

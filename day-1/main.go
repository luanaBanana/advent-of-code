package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(e error, message string) {
	if e != nil {
		log.Fatalf("%s. error: %v", message, e)
	}
}

func findAllOverlapping(str string, reg *regexp.Regexp) []string {
	var matches []string
	for i := 0; i < len(str); i++ {
		substr := str[i:]
		match := reg.FindString(substr)
		if match != "" {
			matches = append(matches, match)
		}
	}
	return matches
}

func main() {

	result := 0
	dat, err := os.Open("./one.txt")
	check(err, "Failed to open file")
	defer dat.Close()

	scanner := bufio.NewScanner(dat) //scan the contents of a file and print line by line
	re := regexp.MustCompile("\\d")

	for scanner.Scan() {
		line := scanner.Text()
		digits := re.FindAllString(line, -1)
		firstAndLast := digits[0] + digits[len(digits)-1]
		value, err := strconv.Atoi(firstAndLast)
		check(err, "Failed to convert to int")
		result = result + value //cumulus

	}

	log.Printf("Result Part 1: %d", result)

	result = 0

	wordToNumber := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	reg := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|\\d)")

	dat2, err := os.Open("./one.txt")
	check(err, "Failed to openn file")
	defer dat2.Close()
	scanner2 := bufio.NewScanner(dat2) //scan the contents of a file
	value := 0
	for scanner2.Scan() { // print line by line
		var firstNum, lastNum int
		line := scanner2.Text()

		findings := findAllOverlapping(line, reg)

		if first, found := wordToNumber[findings[0]]; found {
			firstNum = first
		} else {
			firstNum, err = strconv.Atoi(findings[0])
			check(err, "Failed to convert to int")
		}

		if last, found := wordToNumber[findings[len(findings)-1]]; found {
			lastNum = last
		} else {
			lastNum, err = strconv.Atoi(findings[len(findings)-1])
			check(err, "Failed to convert to int")
		}

		value = firstNum*10 + lastNum
		result = result + value //cumulus

	}

	log.Printf("Result Part 2: %d", result)

}

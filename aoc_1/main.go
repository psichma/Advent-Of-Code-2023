package aoc_1

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func load_file(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New("error reading file, " + filename)
	}
	return string(file), nil
}

func part1(line string) (int64) {
	found_first := false
	var first string
	var last string
	for _, rune_char := range line {
		str_char := string(rune_char)
		_, err := strconv.ParseInt(str_char, 10, 32)
		if err == nil && !found_first {
			found_first = true
			first = str_char
		} else if err == nil {
			last = str_char
		}
	}
	if(last == "") {
		last = first
	}

	concat := first + last

	num, err := strconv.ParseInt(concat, 10, 32)
	if err != nil {
		return 0
	}
	return num
}

func part2(line string) (int64) {
	number_strings := map[string]string {
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	found_first := false
	var first string
	var last string
	
	for len(line) > 0 {
		current_rune := line[0]
		current_rune_string := string(current_rune)
		_, err := strconv.ParseInt(current_rune_string, 10, 32)
		if err == nil && !found_first {
			found_first = true
			first = current_rune_string
		} else if err == nil {
			last = current_rune_string
		// end check for actual digit
		} else { // is a character, so it could be a number

			// loop through numbers
			for number_full, number_int := range number_strings {
				idx := strings.Index(line, number_full)
				if idx == 0 {
					// found it
					if !found_first {
						found_first = true
						first = number_int
					} else {
						last = number_int
					}
					break
				}
			}
		}
		line = line[1:]
	}
	if(last == "") {
		last = first
	}

	concat := first + last

	num, err := strconv.ParseInt(concat, 10, 32)
	if err != nil {
		return 0
	}
	return num
}

func Run(file_name string) {
	content, err := load_file(file_name)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(content, "\n")

	var sum1 int64
	var sum2 int64

	for _, line := range lines {
		sum1 += part1(line) // 54159
		sum2 += part2(line) // 53866
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)

	
}
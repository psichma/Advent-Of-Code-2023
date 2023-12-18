package aoc_2

import (
	"errors"
	"log"
	"os"
	"strings"
)

func load_file(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New("error reading file, " + filename)
	}
	return string(file), nil
}

func part1(line string) {

}

func Run(file_name string) {
	content, err := load_file(file_name)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		part1(line)
	}
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const inputFileName = "input"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0

	for scanner.Scan() {
		line := scanner.Text()

		res1 += part1(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
}

func part1(line string) int {
	re := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	parts := re.FindAllStringSubmatch(line, -1)

	res := 0

	for _, part := range parts {
		n1, _ := strconv.Atoi(part[1])
		n2, _ := strconv.Atoi(part[2])
		res += n1 * n2
	}

	return res
}

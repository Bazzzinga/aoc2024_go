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

const inputFileName = "input"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0
	res2 := 0

	ll := ""

	for scanner.Scan() {
		line := scanner.Text()

		ll += line
	}

	res1 += part1(ll)
	res2 += part2(ll)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
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

func part2(line string) int {
	line = strings.Replace(line, "\n", "", -1)

	doRe := regexp.MustCompile("do\\(\\)")
	doParts := doRe.FindAllStringSubmatchIndex(line, -1)

	for i, p := range doParts {
		line = line[:p[0]+i] + "\n" + line[p[0]+i:]
	}

	dontRe := regexp.MustCompile("don't\\(\\)")
	dontParts := dontRe.FindAllStringSubmatchIndex(line, -1)

	for i, p := range dontParts {
		line = line[:p[0]+i] + "\n" + line[p[0]+i:]
	}

	parts := strings.Split(line, "\n")

	res := 0

	for _, p := range parts {
		if !strings.HasPrefix(p, "don't()") {
			res += part1(p)
		}
	}

	return res
}

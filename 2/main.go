package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, " ")
		levels := make([]int, len(tmp))

		for i := range tmp {
			l, _ := strconv.Atoi(tmp[i])
			levels[i] = l
		}

		if part1(levels) {
			res1++
		}

		if part2(levels) {
			res2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func part2(level []int) bool {
	l := len(level)
	for i := 0; i < l; i++ {
		tmp := make([]int, l)
		copy(tmp, level)

		sub := append(tmp[:i], tmp[i+1:]...)

		safe := part1(sub)
		if safe {
			return true
		}

	}

	return false
}

func part1(level []int) bool {
	if len(level) < 2 {
		return true
	}

	inc := level[1] > level[0]
	prev := level[0]
	for i, l := range level {
		if i == 0 {
			continue
		}

		if inc {
			if l <= prev || l-prev > 3 {
				return false
			}
		} else {
			if l >= prev || prev-l > 3 {
				return false
			}
		}

		prev = l
	}

	return true
}

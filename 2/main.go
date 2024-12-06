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

	//reports := make([][]int, 0)
	res1 := 0

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, " ")
		levels := make([]int, len(tmp))

		for i := range tmp {
			l, _ := strconv.Atoi(tmp[i])
			levels[i] = l
		}

		isSafe := part1(levels)
		if isSafe {
			res1++
		}
		//reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
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

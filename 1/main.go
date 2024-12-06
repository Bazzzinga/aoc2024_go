package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	left := make([]int, 0)
	right := make([]int, 0)

	rightMap := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, " ")
		n, _ := strconv.Atoi(tmp[0])
		left = append(left, n)

		n, _ = strconv.Atoi(tmp[len(tmp)-1])
		right = append(right, n)

		rightMap[n]++
	}

	res1 := part1(left, right)
	res2 := part2(left, rightMap)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func part2(left []int, rightMap map[int]int) int {
	res := 0

	for _, l := range left {
		res += l * rightMap[l]
	}

	return res
}

func part1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	res := 0

	for i := range left {
		d := left[i] - right[i]
		if d < 0 {
			d *= -1
		}
		res += d
	}

	return res
}

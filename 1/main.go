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

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, " ")
		n, _ := strconv.Atoi(tmp[0])
		left = append(left, n)

		n, _ = strconv.Atoi(tmp[len(tmp)-1])
		right = append(right, n)
	}

	res1 := part1(left, right)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
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

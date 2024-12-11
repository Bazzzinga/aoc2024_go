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

var cache map[string]int

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cache = make(map[string]int)

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
	res := 0

	ss := strings.Split(line, " ")

	blinks := 25

	for _, s := range ss {
		stone, _ := strconv.Atoi(s)

		res += calcStone1(int64(stone), blinks)
	}

	return res
}

func calcStone1(n int64, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := fmt.Sprintf("%d_%d", n, blinks)
	res, ok := cache[key]
	if ok {
		return res
	}

	if n == 0 {
		res = calcStone1(1, blinks-1)

	} else {
		if numberLength(n)%2 == 0 {
			n1, n2 := splitNumber(n)

			res = calcStone1(n1, blinks-1) + calcStone1(n2, blinks-1)
		} else {
			res = calcStone1(n*2024, blinks-1)
		}
	}

	cache[key] = res

	return res
}

func numberLength(n int64) int {
	res := 0

	for n > 0 {
		res++
		n /= 10
	}

	return res
}

func splitNumber(n int64) (int64, int64) {
	l := numberLength(n)

	d := int64(1)

	for i := 0; i < l/2; i++ {
		d *= 10
	}

	return n / d, n % d
}

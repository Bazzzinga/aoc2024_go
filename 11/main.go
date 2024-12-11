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
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		res1 += solve(line, 25)
		res2 += solve(line, 75)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func solve(line string, blinks int) int {
	res := 0

	ss := strings.Split(line, " ")

	for _, s := range ss {
		stone, _ := strconv.Atoi(s)

		res += calcStone(int64(stone), blinks)
	}

	return res
}

func calcStone(n int64, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := fmt.Sprintf("%d_%d", n, blinks)
	res, ok := cache[key]
	if ok {
		return res
	}

	if n == 0 {
		res = calcStone(1, blinks-1)

	} else {
		if numberLength(n)%2 == 0 {
			n1, n2 := splitNumber(n)

			res = calcStone(n1, blinks-1) + calcStone(n2, blinks-1)
		} else {
			res = calcStone(n*2024, blinks-1)
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

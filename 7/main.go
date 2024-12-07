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

type Test struct {
	Result  int
	Numbers []int
}

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

		parts1 := strings.Split(line, ": ")

		r, _ := strconv.Atoi(parts1[0])

		parts2 := strings.Split(parts1[1], " ")

		t := Test{
			Result:  r,
			Numbers: make([]int, len(parts2)),
		}

		for i, p := range parts2 {
			n, _ := strconv.Atoi(p)
			t.Numbers[i] = n
		}

		if t.isPossible() {
			res1 += t.Result
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
}

func (t Test) isPossible() bool {
	return t.recursion(t.Numbers[0], t.Numbers[1:])
}

func (t Test) recursion(r int, nn []int) bool {
	if len(nn) == 0 {
		return r == t.Result
	}

	ok1 := t.recursion(r+nn[0], nn[1:])
	ok2 := t.recursion(r*nn[0], nn[1:])

	return ok1 || ok2
}

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
	res2 := 0

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

		if t.isPossible1() {
			res1 += t.Result
		}
		if t.isPossible2() {
			res2 += t.Result
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func (t Test) isPossible1() bool {
	return t.recursion1(t.Numbers[0], t.Numbers[1:])
}

func (t Test) isPossible2() bool {
	return t.recursion2(t.Numbers[0], t.Numbers[1:])
}

func (t Test) recursion1(r int, nn []int) bool {
	if len(nn) == 0 {
		return r == t.Result
	}

	ok1 := t.recursion1(r+nn[0], nn[1:])
	ok2 := t.recursion1(r*nn[0], nn[1:])

	return ok1 || ok2
}

func (t Test) recursion2(r int, nn []int) bool {
	if len(nn) == 0 {
		return r == t.Result
	}

	ok1 := t.recursion2(r+nn[0], nn[1:])
	ok2 := t.recursion2(r*nn[0], nn[1:])

	ok3 := t.recursion2(t.concat(r, nn[0]), nn[1:])

	return ok1 || ok2 || ok3
}

func (t Test) concat(a, b int) int {
	tmp := strconv.Itoa(b)
	l := len(tmp)

	for i := 0; i < l; i++ {
		a *= 10
	}

	return a + b
}

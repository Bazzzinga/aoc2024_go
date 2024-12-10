package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFileName = "input"

type Map struct {
	h [][]int
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := Map{
		h: make([][]int, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		rs := []rune(line)

		tmp := make([]int, len(rs))

		for i, r := range rs {
			tmp[i] = int(r) - '0'
		}

		m.h = append(m.h, tmp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1 := m.part1()
	res2 := m.part2()

	fmt.Println(res1)
	fmt.Println(res2)
}

func (m *Map) part1() int {
	res := 0

	for y, row := range m.h {
		for x, h := range row {
			if h == 0 {
				reached := make(map[string]struct{})
				res += m.countNines(x, y, 0, reached)
			}
		}
	}

	return res
}

func (m *Map) part2() int {
	res := 0

	for y, row := range m.h {
		for x, h := range row {
			if h == 0 {
				res += m.countNines2(x, y, 0)
			}
		}
	}

	return res
}

func (m *Map) countNines2(x, y, count int) int {
	h := m.h[y][x]

	if h == 9 {
		return count + 1
	}

	res := count

	//up
	if y > 0 && m.h[y-1][x]-h == 1 {
		res += m.countNines2(x, y-1, count)
	}

	//down
	if y < len(m.h)-1 && m.h[y+1][x]-h == 1 {
		res += m.countNines2(x, y+1, count)
	}

	//left
	if x > 0 && m.h[y][x-1]-h == 1 {
		res += m.countNines2(x-1, y, count)
	}

	//right
	if x < len(m.h[0])-1 && m.h[y][x+1]-h == 1 {
		res += m.countNines2(x+1, y, count)
	}

	return res
}

func (m *Map) countNines(x, y, count int, reached map[string]struct{}) int {
	h := m.h[y][x]

	key := fmt.Sprintf("%d_%d", x, y)

	if h == 9 {
		_, ok := reached[key]
		if ok {
			return count
		}

		reached[key] = struct{}{}

		return count + 1
	}

	res := count

	//up
	if y > 0 && m.h[y-1][x]-h == 1 {
		res += m.countNines(x, y-1, count, reached)
	}

	//down
	if y < len(m.h)-1 && m.h[y+1][x]-h == 1 {
		res += m.countNines(x, y+1, count, reached)
	}

	//left
	if x > 0 && m.h[y][x-1]-h == 1 {
		res += m.countNines(x-1, y, count, reached)
	}

	//right
	if x < len(m.h[0])-1 && m.h[y][x+1]-h == 1 {
		res += m.countNines(x+1, y, count, reached)
	}

	return res
}

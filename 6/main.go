package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFileName = "input"

type Field struct {
	Cells [][]rune
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0

	f := Field{
		Cells: make([][]rune, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		f.Cells = append(f.Cells, []rune(line))
	}

	res1 = f.part1()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
}

func (f Field) part1() int {
	x, y, dir := f.start()

	checkedMap := make(map[string]struct{})

	for {
		if x < 0 || y < 0 || x >= len(f.Cells[0]) || y >= len(f.Cells) {
			break
		}

		checkedMap[fmt.Sprintf("%d_%d", x, y)] = struct{}{}

		x, y, dir = f.checkDir(x, y, dir)
	}

	return len(checkedMap)
}

func (f Field) checkDir(x, y int, dir rune) (int, int, rune) {
	ok := false

	var nx, ny int

	for !ok {
		dx, dy := f.dirToVector(dir)

		nx = x + dx
		ny = y + dy

		if nx < 0 || ny < 0 || nx >= len(f.Cells[0]) || ny >= len(f.Cells) {
			ok = true
			break
		}

		ok = f.Cells[ny][nx] != '#'

		if !ok {
			dir = f.nextDir(dir)
		}
	}

	return nx, ny, dir
}

func (f Field) nextDir(dir rune) rune {
	switch dir {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}

	return '^'
}

func (f Field) dirToVector(dir rune) (int, int) {
	switch dir {
	case 'v':
		return 0, 1
	case '^':
		return 0, -1
	case '<':
		return -1, 0
	case '>':
		return 1, 0
	}

	return 0, 0
}

func (f Field) start() (int, int, rune) {
	for y := 0; y < len(f.Cells); y++ {
		for x := 0; x < len(f.Cells[y]); x++ {
			r := f.Cells[y][x]
			if r == '^' || r == 'v' || r == '<' || r == '>' {
				f.Cells[y][x] = '.'
				return x, y, r
			}
		}
	}

	return 0, 0, '^'
}

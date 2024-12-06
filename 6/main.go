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

	f := Field{
		Cells: make([][]rune, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		f.Cells = append(f.Cells, []rune(line))
	}

	//res1 := f.part1()
	res2 := f.part2()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(res1)
	fmt.Println(res2)
}

func (f Field) part2() int {
	//цикл = повторный проход по одной клетке в том же направлении
	x, y, _ := f.start()
	basePath := f.mapPath()
	delete(basePath, fmt.Sprintf("%d_%d", x, y))

	res := 0

	for c := range basePath {
		parts := strings.Split(c, "_")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		if f.hasLoop(x, y) {
			res++
		}
	}

	return res
}

func (f Field) hasLoop(ox, oy int) bool {
	pathDirMap := make(map[string]struct{})

	x, y, dir := f.start()

	for {
		key := fmt.Sprintf("%d_%d_%s", x, y, string(dir))

		_, ok := pathDirMap[key]
		if ok {
			return true
		}

		pathDirMap[key] = struct{}{}

		if x < 0 || y < 0 || x >= len(f.Cells[0]) || y >= len(f.Cells) {
			return false
		}

		f.Cells[oy][ox] = '#'
		x, y, dir = f.checkDir(x, y, dir)
		f.Cells[oy][ox] = '.'
	}
}

func (f Field) part1() int {
	return len(f.mapPath())
}

func (f Field) mapPath() map[string]struct{} {
	x, y, dir := f.start()

	checkedMap := make(map[string]struct{})

	for {
		if x < 0 || y < 0 || x >= len(f.Cells[0]) || y >= len(f.Cells) {
			break
		}

		checkedMap[fmt.Sprintf("%d_%d", x, y)] = struct{}{}

		x, y, dir = f.checkDir(x, y, dir)
	}

	return checkedMap
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
				return x, y, r
			}
		}
	}

	return 0, 0, '^'
}

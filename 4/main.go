package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFileName = "input"

type Field struct {
	f [][]rune
	d map[Coord][]Coord
}

type Coord struct {
	x int
	y int
}

var ALL = Coord{x: 0, y: 0}
var UP = Coord{x: 0, y: -1}
var DOWN = Coord{x: 0, y: 1}
var LEFT = Coord{x: -1, y: 0}
var RIGHT = Coord{x: 1, y: 0}
var UPLEFT = Coord{x: -1, y: -1}
var UPRIGHT = Coord{x: 1, y: -1}
var DOWNLEFT = Coord{x: -1, y: 1}
var DOWNRIGHT = Coord{x: 1, y: 1}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := Field{
		f: make([][]rune, 0),
		d: map[Coord][]Coord{
			ALL:       {UP, DOWN, LEFT, RIGHT, UPLEFT, UPRIGHT, DOWNLEFT, DOWNRIGHT},
			UP:        {UPLEFT, UP, UPRIGHT},
			DOWN:      {DOWNLEFT, DOWN, DOWNRIGHT},
			LEFT:      {UPLEFT, LEFT, DOWNLEFT},
			RIGHT:     {UPRIGHT, RIGHT, DOWNRIGHT},
			UPLEFT:    {LEFT, UPLEFT, UP},
			UPRIGHT:   {RIGHT, UPRIGHT, UP},
			DOWNLEFT:  {DOWN, DOWNLEFT, LEFT},
			DOWNRIGHT: {DOWN, DOWNRIGHT, RIGHT},
		},
	}

	for scanner.Scan() {
		line := scanner.Text()

		field.f = append(field.f, []rune(line))
	}

	res1 := field.part1()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
}

func (f *Field) part1() int {
	res := 0

	for y, row := range f.f {
		for x := range row {
			if f.f[y][x] == 'X' {
				res += f.countXMAS(Coord{x: x, y: y})
			}
		}
	}

	return res
}

func (f *Field) countXMAS(c Coord) int {
	res := 0
	mCoords := f.searchNear('M', c, ALL)

	for _, mCoord := range mCoords {
		d := Coord{
			x: mCoord.x - c.x,
			y: mCoord.y - c.y,
		}
		aCoords := f.searchNear('A', mCoord, d)

		for _, aCoord := range aCoords {
			d := Coord{
				x: aCoord.x - mCoord.x,
				y: aCoord.y - mCoord.y,
			}

			sCoords := f.searchNear('S', aCoord, d)

			res += len(sCoords)
		}
	}

	return res
}

func (f *Field) searchNear(r rune, at, dir Coord) []Coord {
	res := make([]Coord, 0)

	if dir == ALL {
		ds := f.d[dir]

		for _, d := range ds {
			x := at.x + d.x
			y := at.y + d.y

			if x >= 0 && x < len(f.f[0]) && y >= 0 && y < len(f.f) {
				if f.f[y][x] == r {
					res = append(res, Coord{x: x, y: y})
				}
			}
		}
	} else {
		x := at.x + dir.x
		y := at.y + dir.y

		if x >= 0 && x < len(f.f[0]) && y >= 0 && y < len(f.f) {
			if f.f[y][x] == r {
				res = append(res, Coord{x: x, y: y})
			}
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}

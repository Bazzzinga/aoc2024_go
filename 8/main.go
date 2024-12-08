package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFileName = "input"

type Coord struct {
	x int
	y int
}

type City struct {
	maxX     int
	maxY     int
	antennas map[string][]Coord
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	c := &City{
		antennas: make(map[string][]Coord),
		maxY:     -1,
	}

	for scanner.Scan() {
		line := scanner.Text()

		c.parseLine(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1 := c.part1()
	res2 := c.part2()

	fmt.Println(res1)
	fmt.Println(res2)
}

func (c *City) parseLine(line string) {
	if c.maxX == 0 {
		c.maxX = len(line) - 1
	}

	c.maxY++
	y := c.maxY

	ss := strings.Split(line, "")
	for i, s := range ss {
		if s != "." {
			_, ok := c.antennas[s]
			if !ok {
				c.antennas[s] = make([]Coord, 0)
			}
			c.antennas[s] = append(c.antennas[s], Coord{x: i, y: y})
		}
	}
}

func (c *City) part2() int {
	nodes := make(map[string]struct{})

	for _, as := range c.antennas {
		l := len(as)
		if l > 1 {
			for i, a1 := range as {
				tmp := make([]Coord, l)
				copy(tmp, as)
				tmp[i] = tmp[l-1]
				tmp = tmp[:l-1]

				nodes[fmt.Sprintf("%d_%d", a1.x, a1.y)] = struct{}{}

				for _, a2 := range tmp {
					ns := c.nodes(a1, a2)

					for _, n := range ns {
						nodes[fmt.Sprintf("%d_%d", n.x, n.y)] = struct{}{}
					}
				}
			}
		}
	}

	return len(nodes)
}

func (c *City) nodes(a1, a2 Coord) []Coord {
	dx := a1.x - a2.x
	dy := a1.y - a2.y

	res := make([]Coord, 0)

	nx := a1.x
	ny := a1.y

	for {
		nx += dx
		ny += dy

		if nx >= 0 && ny >= 0 && nx <= c.maxX && ny <= c.maxY {
			res = append(res, Coord{x: nx, y: ny})
		} else {
			break
		}
	}

	return res
}

func (c *City) part1() int {
	nodes := make(map[string]struct{})

	for _, as := range c.antennas {
		l := len(as)
		if l > 1 {
			for i, a1 := range as {
				tmp := make([]Coord, l)
				copy(tmp, as)
				tmp[i] = tmp[l-1]
				tmp = tmp[:l-1]

				for _, a2 := range tmp {
					n := c.node(a1, a2)
					if n.x >= 0 && n.y >= 0 && n.x <= c.maxX && n.y <= c.maxY {
						nodes[fmt.Sprintf("%d_%d", n.x, n.y)] = struct{}{}
					}
				}
			}
		}
	}

	return len(nodes)
}

func (c *City) node(a1, a2 Coord) Coord {
	dx := a1.x - a2.x
	dy := a1.y - a2.y

	return Coord{
		x: a1.x + dx,
		y: a1.y + dy,
	}
}

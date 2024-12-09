package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFileName = "input"

type FS struct {
	mem []*int
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fs := &FS{}

	for scanner.Scan() {
		line := scanner.Text()

		fs.Parse(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1 := fs.checksum1()

	fmt.Println(res1)
}

func (s *FS) Parse(line string) {
	rs := []rune(line)

	s.mem = make([]*int, 0)

	free := false
	id := 0

	for _, r := range rs {
		c := int(r) - '0'

		var data *int
		if !free {
			tmp := id
			data = &tmp
			id++
		}
		for i := 0; i < c; i++ {
			s.mem = append(s.mem, data)
		}

		free = !free
	}
}

func (s *FS) print() {
	for _, m := range s.mem {
		if m == nil {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", *m)
		}
	}
	fmt.Printf("\n")
}

func (s *FS) checksum1() int {
	//pack

	l := -1
	r := len(s.mem) - 1

	for l < r {
		l++
		if s.mem[l] != nil {
			continue
		}

		s.mem[l], s.mem[r] = s.mem[r], s.mem[l]
		for s.mem[r] == nil {
			r--
		}
	}

	//checksum
	res := 0

	i := 0
	for s.mem[i] != nil {
		res += i * *s.mem[i]

		i++
	}

	return res
}

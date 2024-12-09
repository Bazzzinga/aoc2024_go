package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const inputFileName = "input"

type File struct {
	idx  int
	size int
	pos  int
}

type FileSystem struct {
	mem       []*int
	fileSize  map[int]int
	files     []File
	freeSpace []File
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fs := &FileSystem{
		fileSize: make(map[int]int),
	}

	for scanner.Scan() {
		line := scanner.Text()

		fs.Parse(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//res1 := fs.checksum1()
	res2 := fs.checksum2()

	//fmt.Println(res1)
	fmt.Println(res2)
}

func (s *FileSystem) Parse(line string) {
	rs := []rune(line)

	s.mem = make([]*int, 0)
	s.files = make([]File, 0)

	free := false
	id := 0

	pos := 0

	for _, r := range rs {
		c := int(r) - '0'

		var data *int
		if !free {
			tmp := id
			data = &tmp
			s.fileSize[id] = c
			s.files = append(s.files, File{idx: id, size: c, pos: pos})
			id++
		}
		for i := 0; i < c; i++ {
			s.mem = append(s.mem, data)
		}
		pos += c
		free = !free
	}

	sort.SliceStable(s.files, func(i, j int) bool {
		return s.files[i].idx > s.files[j].idx
	})
}

func (s *FileSystem) rebuildFreeSpaceMap() {
	s.freeSpace = make([]File, 0)

	idx := -1
	l := 0
	for i := 0; i < len(s.mem); i++ {
		if s.mem[i] == nil {
			if idx < 0 {
				idx = i
				l = 1
			} else {
				l++
			}
		} else if idx > 0 {
			s.freeSpace = append(s.freeSpace, File{idx: idx, size: l})
			idx = -1
			l = 0
		}
	}
}

func (s *FileSystem) checksum2() int {
	//pack

	for _, file := range s.files {
		if file.idx == 0 {
			break
		}
		s.rebuildFreeSpaceMap()

		for _, freeSpace := range s.freeSpace {
			if freeSpace.size >= file.size && freeSpace.idx < file.pos {
				//move file
				for j := 0; j < file.size; j++ {
					s.mem[freeSpace.idx+j], s.mem[file.pos+j] = s.mem[file.pos+j], s.mem[freeSpace.idx+j]
				}
				break
			}
		}
	}

	return s.calcChecksum2()
}

func (s *FileSystem) print() {
	for _, m := range s.mem {
		if m == nil {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", *m)
		}
	}
	fmt.Printf("\n")
}

func (s *FileSystem) checksum1() int {
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

	return s.calcChecksum1()
}

func (s *FileSystem) calcChecksum1() int {
	res := 0

	i := 0
	for s.mem[i] != nil {
		res += i * *s.mem[i]

		i++
	}

	return res
}

func (s *FileSystem) calcChecksum2() int {
	res := 0

	for i, d := range s.mem {
		if d != nil {
			res += i * *d

			i++
		}
	}

	return res
}

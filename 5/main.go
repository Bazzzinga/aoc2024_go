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

type Rule struct {
	Before int
	After  int
}

type Updates struct {
	Rules map[int][]Rule
	Cache map[string]bool
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	u := Updates{
		Rules: make(map[int][]Rule),
		Cache: make(map[string]bool),
	}

	rules := true

	res1 := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			rules = false
			continue
		}

		if rules {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])

			_, ok := u.Rules[before]
			if !ok {
				u.Rules[before] = make([]Rule, 0)
			}
			u.Rules[before] = append(u.Rules[before], Rule{Before: before, After: after})
		} else {
			ss := strings.Split(line, ",")

			nn := make([]int, len(ss))
			for i, s := range ss {
				n, _ := strconv.Atoi(s)
				nn[i] = n
			}

			res1 += u.part1(nn)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res1)
}

func (u *Updates) part1(nn []int) int {
	l := len(nn)

	middle := nn[l/2]

	for i := 0; i < l-1; i++ {
		n := nn[i]
		for j := i + 1; j < l; j++ {
			m := nn[j]

			key := fmt.Sprintf("%d_%d", n, m)

			cached, ok := u.Cache[key]
			if ok {
				if !cached {
					return 0
				}
			} else {
				rules := u.Rules[m]

				for _, r := range rules {
					if r.After == n {
						u.Cache[key] = false
						return 0
					}
				}

				u.Cache[key] = true
			}
		}
	}

	return middle
}

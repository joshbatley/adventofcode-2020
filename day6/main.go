package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputs := loadInput()
	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	sum := 0
	for _, v := range inputs {
		var m = make(map[string]bool)
		for _, c := range v {
			if c != '\n' {
				m[string(c)] = true
			}
		}
		sum += len(m)
	}
	log.Println("Part 1 sum is", sum)
}

func part2(inputs []string) {
	sum := 0
	for _, v := range inputs {
		var mappp = make(map[string]int)

		groupCount := 1

		for i, char := range v {
			if char != '\n' {
				as := string(char)
				mappp[as] = mappp[as] + 1
			}
			if char == '\n' && i != len(v)-1 {
				groupCount++
			}
		}
		for i, k := range mappp {
			if k != groupCount {
				delete(mappp, i)
			}
		}
		sum += len(mappp)
	}
	log.Println("part 2 sum is", sum)
}

func loadInput() []string {
	data, err := ioutil.ReadFile("./inputs.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n\n")
}

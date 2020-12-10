package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputs := loadInputs()
	sort.Ints(inputs)
	part1(inputs)
	part2(inputs)
}

func part1(inputs []int) {
	one := 1
	three := 1
	for i, k := range inputs {
		if i+1 >= len(inputs) {
			break
		}
		switch inputs[i+1] - k {
		case 1:
			one++
		case 3:
			three++
		}
	}

	log.Println("Number of 1-jolt differences multiplied by the number of 3-jolt", one, three)
}

var counted map[int]int

func part2(inputs []int) {
	counted = make(map[int]int)
	route := make(map[int]bool)

	for _, i := range inputs {
		route[i] = true
	}

	route[inputs[len(inputs)-1]+3] = true
	val := looper(0, route)
	log.Println("Total number of possible sizes", val)
}

func looper(start int, r map[int]bool) int {
	count, ok := counted[start]
	if ok {
		return count
	}

	res := 0
	q := true

	for i := 1; i <= 3; i++ {
		pos := start + i
		_, ok := r[pos]
		if ok {
			res += looper(pos, r)
			q = false
		}
	}

	if q {
		res++
	}

	counted[start] = res

	return res
}

func loadInputs() []int {
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		d = append(d, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

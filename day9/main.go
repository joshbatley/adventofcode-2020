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
	num := part1(inputs)
	part2(inputs, num)
}

func part2(inputs []int, num int) {
	// find num
	for i, v := range inputs {
		list := []int{v}
		sum := v
		found := false

		for _, next := range inputs[i+1:] {
			if sum > num {
				break
			}
			if sum+next <= num {
				sum += next
				list = append(list, next)
				if sum == num {
					found = true
					break
				}
			}
		}
		if found {
			sort.Ints(list)
			log.Println("The encryption weakness is", list[0]+list[len(list)-1])
			break
		}
	}
}

func part1(inputs []int) int {
	offset := 25
	tt := 0
	for i := 0; i <= len(inputs); i++ {
		top := i + offset
		if top >= len(inputs) {
			break
		}
		r := inputs[i:top]
		tt = inputs[top]
		found := false
		for _, k := range r {
			for _, k2 := range r {
				if k == k2 {
					continue
				}
				if k+k2 == tt {
					found = true
				}
			}
		}
		if !found {
			log.Println("Broken value is", tt)
			break
		}
	}
	return tt
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

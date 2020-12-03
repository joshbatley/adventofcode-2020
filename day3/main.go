package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	inputs := loadInputs()
	part1(inputs)
	part2(inputs)

}

func part1(i []string) {
	pos := 0
	trees := 0
	for _, r := range i {
		if string(r[pos]) == "#" {
			trees++
		}
		pos = pos + 3
		if pos > len(r)-1 {
			pos = pos - len(r)
		}
	}
	log.Println("the number of tress hit are", trees)
}

func part2(i []string) {
	rule := [][]int{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}
	pr := 1
	for _, r := range rule {
		pos := 0
		trees := 0
		for idx, in := range i {
			if idx%r[1] != 0 {
				continue
			}
			if string(in[pos]) == "#" {
				trees++
			}
			pos = pos + r[0]
			if pos > len(in)-1 {
				pos = pos - len(in)
			}
		}
		pr = pr * trees
	}
	log.Println("the total tree multipled together is", pr)
}

func loadInputs() []string {
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d = append(d, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

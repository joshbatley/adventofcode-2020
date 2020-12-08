package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var count = make(map[string]bool, 0)

func main() {
	inputs := loadInputs()
	m := createStuct(inputs)
	part1(m)
	part2(m)
}

func part2(m map[string][]string) {
	log.Println("there are", loopy2(m, "shiny gold"), "inside one shiny gold bag")
}

func loopy2(m map[string][]string, k string) int {
	total := 0
	for b, i := range m {
		if b != k {
			continue
		}
		for _, t := range i {
			if t == "no other" {
				return 0
			}
			s := strings.SplitN(t, " ", 2)
			num, _ := strconv.Atoi(s[0])
			total += num + (num * loopy2(m, s[1]))
		}
	}
	return total
}

func part1(m map[string][]string) {
	for k, v := range m {
		loopy(v, m, k)
	}
	log.Println("number of possible gold bags", len(count))
}

func loopy(v []string, m map[string][]string, k string) {
	for _, b := range v {
		if b == "no other" || count[k] == true {
			break
		}
		if strings.Contains(b, "shiny gold") {
			count[k] = true
			break
		}
		q := regexp.MustCompile(`\d+ `)
		b = q.ReplaceAllString(b, "")
		loopy(m[b], m, k)
	}
}

func createStuct(inputs []string) map[string][]string {
	m := make(map[string][]string, 0)
	rgx := regexp.MustCompile(`( bags?)((\, )|( contain  ?)|(\.?))`)
	for _, i := range inputs {
		s := rgx.Split(i, -1)
		m[s[0]] = s[1:(len(s[1:]))]
	}

	return m
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

package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var mustHave = [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

type pp struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func main() {
	i := loadInput()
	zp := regexp.MustCompile(`\n|\s`)
	valid := make([][][]string, 0)
	for _, v := range i {
		s := zp.Split(string(v), -1)
		t := make([][]string, 0)
		for _, k := range s {
			key := strings.Split(k, ":")
			for _, m := range mustHave {
				if key[0] == m {
					i := []string{key[0], key[1]}
					t = append(t, i)
				}
			}
		}
		if len(t) >= 7 {
			valid = append(valid, t)
		}
	}
	log.Printf("Number of valid passports %v", len(valid))

	allValid := 0
	for _, r := range valid {
		c := 0
		for _, v := range r {
			switch v[0] {
			case "byr":
				y, err := strconv.Atoi(v[1])
				if err == nil && y >= 1920 && y <= 2002 {
					c++
				}
			case "iyr":
				y, err := strconv.Atoi(v[1])
				if err == nil && y >= 2010 && y <= 2020 {
					c++
				}
			case "eyr":
				y, err := strconv.Atoi(v[1])
				if err == nil && y >= 2020 && y <= 2030 {
					c++
				}
			case "hgt":
				if strings.Contains(v[1], "cm") {
					h, err := strconv.Atoi(strings.Split(v[1], "cm")[0])
					if err == nil && h >= 150 && h <= 193 {
						c++
					}
				}
				if strings.Contains(v[1], "in") {
					h, err := strconv.Atoi(strings.Split(v[1], "in")[0])
					if err == nil && h >= 59 && h <= 76 {
						c++
					}
				}

			case "hcl":
				zp := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
				if zp.Match([]byte(v[1])) {
					c++
				}
			case "ecl":
				ecs := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				for _, e := range ecs {
					if e == v[1] {
						c++
						break
					}
				}
			case "pid":
				_, err := strconv.Atoi(v[1])
				if err == nil && len(v[1]) == 9 {
					c++
				}
			}
		}
		if c >= 7 {
			allValid++
		}
	}
	log.Printf("Number of completely solid %v", allValid)
}

func loadInput() []string {
	data, err := ioutil.ReadFile("./inputs.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n\n")
}

// iyr:1928 cid:150 pid:476113241 eyr:2039 hcl:a5ac0f
// ecl:#25f8d2
// byr:2027 hgt:190

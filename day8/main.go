package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type data struct {
	n    string
	op   string
	val  int
	seen bool
}

func main() {
	inputs := loadInputs()
	// part1(inputs)
	part2(inputs)
}

func part1(inputs []data) {
	acc := 0
	i := 0
	for i < len(inputs) {
		if inputs[i].seen == true {
			break
		}

		switch inputs[i].n {
		case "nop":
			inputs[i].seen = true
		case "acc":
			inputs[i].seen = true
			updateValues(inputs[i], &acc)
		case "jmp":
			updateValues(inputs[i], &i)
			continue
		}
		i++
	}
	log.Println("acc total value was", acc)
}

func part2(inputs []data) {
	acc := 0

	for z := range inputs {
		co := make([]data, len(inputs))
		copy(co, inputs)

		if co[z].n == "acc" {
			continue
		}
		co[z] = flipFirst(co[z])
		f := false
		acc = 0
		i := 0
		q := 0
		for i < len(co) {
			q++
			if co[i].seen == true {
				f = true
				break
			}
			switch co[i].n {
			case "nop":
				co[i].seen = true
			case "acc":
				co[i].seen = true
				updateValues(co[i], &acc)
			case "jmp":
				co[i].seen = true
				updateValues(co[i], &i)
				continue
			}
			i++
		}

		if f != true {
			break
		}
	}

	log.Println("new acc total value was", acc)
}

func flipFirst(m data) data {
	n := m
	if m.n == "nop" {
		n.n = "jmp"
	}
	if m.n == "jmp" {
		n.n = "nop"
	}
	n.seen = false
	return n
}

func updateValues(d data, v *int) {
	if d.op == "-" {
		*v -= d.val

	} else {
		*v += d.val
	}
}

func loadInputs() []data {
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d []data
	scanner := bufio.NewScanner(file)
	rgx := regexp.MustCompile(`(\w+) (.)(\d+)`)
	for scanner.Scan() {
		text := scanner.Text()
		val, _ := strconv.Atoi(rgx.ReplaceAllString(text, "$3"))
		d = append(d, data{
			n:   rgx.ReplaceAllString(text, "$1"),
			op:  rgx.ReplaceAllString(text, "$2"),
			val: val,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

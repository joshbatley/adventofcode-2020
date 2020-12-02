package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type input struct {
	limit    string
	letter   string
	password string
}

func main() {
	inputs := loadInputs()
	valid := charWithinRange(inputs)
	log.Println("The number of valid passwords are", valid)

	valid = charValidPos(inputs)
	log.Println("The number of valid passwords are", valid)

}

func charValidPos(inputs []input) int {
	valid := 0
	for _, v := range inputs {

		lim := strings.Split(v.limit, "-")
		pos1, _ := strconv.Atoi(lim[0])
		pos2, _ := strconv.Atoi(lim[1])

		let1 := string(v.password[pos1-1])
		let2 := string(v.password[pos2-1])

		notBoth := !(let1 == v.letter && let2 == v.letter)
		butEither := (let1 == v.letter || let2 == v.letter)

		if notBoth && butEither {
			valid++
		}
	}
	return valid
}

func charWithinRange(inputs []input) int {
	valid := 0
	for _, v := range inputs {
		c := 0
		for _, l := range v.password {
			if string(l) == v.letter {
				c++
			}
		}
		lim := strings.Split(v.limit, "-")
		min, _ := strconv.Atoi(lim[0])
		max, _ := strconv.Atoi(lim[1])
		if c <= max && c >= min {
			valid++
		}
	}
	return valid
}

func loadInputs() []input {
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d []input
	scanner := bufio.NewScanner(file)
	// 9-10 m: mmmmnxmmmwm

	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ": ")
		ll := strings.Split(strs[0], " ")
		d = append(d, input{
			limit:    ll[0],
			letter:   ll[1],
			password: strs[1],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

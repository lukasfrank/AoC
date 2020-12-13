package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result = 0
	seen := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
		"cid": true,
	}

	rule := map[string]func(i string) bool{
		"byr": func(i string) bool {
			year, _ := strconv.ParseInt(i, 10, 32)
			return 1920 <= year && year <= 2002
		},
		"iyr": func(i string) bool {
			year, _ := strconv.ParseInt(i, 10, 32)
			return 2010 <= year && year <= 2020
		},
		"eyr": func(i string) bool {
			year, _ := strconv.ParseInt(i, 10, 32)
			return 2020 <= year && year <= 2030
		},
		"hgt": func(i string) bool {
			h, _ := strconv.ParseInt(i[:len(i)-2], 10, 32)
			if strings.HasSuffix(i, "cm") {
				return 150 <= h && h <= 193
			}
			return 59 <= h && h <= 76

		},
		"hcl": func(i string) bool {
			if len(i) != 7 {
				return false
			}
			r, _ := regexp.Compile("#([a-f]|[0-9]){6}")
			return r.MatchString(i)
		},
		"ecl": func(i string) bool {
			if len(i) != 3 {
				return false
			}
			r, _ := regexp.Compile("(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)")
			return r.MatchString(i)
		},
		"pid": func(i string) bool {
			if len(i) != 9 {
				return false
			}
			r, _ := regexp.Compile("[0-9]{9}")
			return r.MatchString(i)
		},
		"cid": func(i string) bool {
			return true
		},
	}
	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Split(txt, " ")

		if txt == "" {
			var temp = true
			for _, e := range seen {
				temp = temp && e
			}
			if temp {
				result++
			}

			seen = map[string]bool{
				"byr": false,
				"iyr": false,
				"eyr": false,
				"hgt": false,
				"hcl": false,
				"ecl": false,
				"pid": false,
				"cid": true,
			}
		}

		for _, el := range parts {
			el := strings.Split(el, ":")

			_, ok := seen[el[0]]
			if ok {
				seen[el[0]] = rule[el[0]](el[1])
			}
		}

	}
	file.Close()

	fmt.Println(result)
}

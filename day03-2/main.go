package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
				seen[el[0]] = true
			}
		}

	}
	file.Close()

	fmt.Println(result)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func bagNames(str string) []nameValue {
	names := []nameValue{}
	bags := strings.Split(str, ",")

	for _, bag := range bags {
		parts := strings.Split(strings.Trim(bag, " "), " ")
		val, _ := strconv.ParseInt(parts[0], 10, 32)
		names = append(names, nameValue{name: parts[1] + parts[2], value: int(val)})
	}
	return names
}

type nameValue struct {
	name  string
	value int
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	mapping := map[string][]nameValue{}

	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Split(txt, "bags contain")
		name := strings.Replace(parts[0], " ", "", -1)

		if _, ok := mapping[name]; !ok {
			mapping[name] = []nameValue{}
		}

		if !strings.Contains(parts[1], "no") {
			mapping[name] = append(mapping[name], bagNames(parts[1])...)
		}

	}
	file.Close()

	count := 0
	traverse := []nameValue{{name: "shinygold", value: 1}}
	var el nameValue
	for len(traverse) > 0 {
		el, traverse = traverse[len(traverse)-1], traverse[:len(traverse)-1]
		if bags, ok := mapping[el.name]; ok {
			for _, b := range bags {
				for i := 0; i < b.value; i++ {
					count++
					traverse = append(traverse, b)
				}
			}
		}
	}

	fmt.Printf("%d", count)
}

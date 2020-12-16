package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func bagNames(str string) []string {
	names := []string{}
	bags := strings.Split(str, ",")

	for _, bag := range bags {
		parts := strings.Split(bag, " ")
		names = append(names, parts[2]+parts[3])
	}
	return names
}
func unique(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}	
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	mapping := map[string][]string{}

	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Split(txt, "bags contain")
		name := strings.Replace(parts[0], " ", "", -1)

		if !strings.Contains(parts[1], "no") {
			for _, bag := range bagNames(parts[1]) {
				if _, ok := mapping[bag]; ok {
					mapping[bag] = append(mapping[bag], name)
				} else {
					mapping[bag] = []string{
						name,
					}
				}
			}
		}
	}
	file.Close()

	count := []string{}
	traverse := []string{"shinygold"}
	el := ""

	for len(traverse) > 0 {
		el, traverse = traverse[len(traverse)-1], traverse[:len(traverse)-1]
		if bags, ok := mapping[el]; ok {
			count = append(count, bags...)
			traverse = append(traverse, bags...)
			continue
		}
	}

	fmt.Println(len(unique(count)))
}

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
	sum := 0
	questions := map[string]bool{}
	first := true
	skip := false
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			sum += len(questions)
			questions = map[string]bool{}
			first = true
			skip = false
			continue
		}
		parts := strings.Split(txt, "")
		if skip {
			continue
		}
		if first {
			for _, s := range parts {
				questions[s] = true
			}
			first = false
			continue
		}

		for k := range questions {
			questions[k] = false
		}
		for _, s := range parts {
			_, ok := questions[s]
			if ok {
				questions[s] = true
			}

		}
		for k, v := range questions {
			if !v {
				delete(questions, k)
			}
		}
		if len(questions) == 0	 {
			skip = true
		}

	}
	sum += len(questions)
	file.Close()
	fmt.Print(sum)

}

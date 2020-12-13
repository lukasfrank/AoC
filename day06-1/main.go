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
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			sum += len(questions)
			questions = map[string]bool{}
			continue
		}
		parts := strings.Split(txt, "")
		for _, s := range parts {
			questions[s] = true
		}
	}
	sum += len(questions)
	file.Close()
	fmt.Print(sum)

}

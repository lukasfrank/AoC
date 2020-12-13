package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func sub(low, high int) int {
	return (high-low)/2 + (high-low)%2
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ids []int
	for scanner.Scan() {
		lowR, highR := 0, 127
		lowC, highC := 0, 7
		txt := scanner.Text()
		parts := strings.Split(txt, "")

		for _, e := range parts {
			switch e {
			case "F":
				highR -= sub(lowR, highR)
			case "B":
				lowR += sub(lowR, highR)
			case "R":
				lowC += sub(lowC, highC)
			case "L":
				highC -= sub(lowC, highC)
			default:
			}
		}

		ids = append(ids, lowR*8+lowC)

	}
	file.Close()

	sort.Ints(ids)

	last := ids[0]
	for _, id := range ids[1:] {
		if last+1 != id {
			fmt.Println(last + 1)
			break

		}
		last++

	}

}

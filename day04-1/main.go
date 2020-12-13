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

	var grid [][]string
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "")
		grid = append(grid, parts)

	}
	file.Close()

	res :=  []int{0, 0, 0, 0, 0}
	ptrI := []int{1, 3, 5, 7, 1}
	ptr := []int{1, 3, 5, 7, 1}
	for lineI, line := range grid[1:] {
		for i := range ptr {
			if i == 4 && lineI % 2 == 0 {
				continue
			}
			if line[ptr[i]%len(line)] == "#" {
				res[i]++
			}
			ptr[i] += ptrI[i]
		}

	}

	for _, v := range res[1:] {
		res[0] *= v
	}
	fmt.Println(res[0])
}

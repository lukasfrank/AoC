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
	
	
	res := 0
	ptr := 3
	for _, line := range grid[1:] {
		if line[ptr % len(line)] == "#" {
			res++
		}
		ptr += 3
	}
	fmt.Println(res)
}

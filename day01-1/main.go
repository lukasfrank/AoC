package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	entries := make(map[int]int)

	for scanner.Scan() {
		i64, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		i := int(i64)

		if val, ok := entries[i]; ok {
			fmt.Println(val * i)
		}
		entries[2020-i] = i
	}
	file.Close()

	fmt.Println(entries)
}

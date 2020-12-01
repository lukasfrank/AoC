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

	values := make([]int, 0)
	for scanner.Scan() {
		i64, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		values = append(values, int(i64))
	}
	file.Close()

	for _, v1 := range values {
		entries := make(map[int]int)
		for _, v2 := range values {
			if val, ok := entries[v2]; ok {
				fmt.Println(v1)
				fmt.Println(v2)
				fmt.Println(val)
				fmt.Println(v1*v2*val)
				return
			}
			entries[2020-v1-v2] = v2
		}
	}

}

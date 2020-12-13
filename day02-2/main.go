package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var pw = 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		rng := strings.Split(parts[0], "-")
		low, _ := strconv.ParseInt(rng[0], 10, 32)
		up, _ := strconv.ParseInt(rng[1], 10, 32)
		if (parts[2][low-1] == parts[1][0]) != (parts[2][up-1] == parts[1][0]) {
			pw++
		}
	}
	file.Close()

	fmt.Println(pw)

}

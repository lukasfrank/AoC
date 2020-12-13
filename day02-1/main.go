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

	// entries := make(map[int]int)
	var pw = 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		rng := strings.Split(parts[0], "-")
		low, _ := strconv.ParseInt(rng[0], 10, 32)
		up, _ := strconv.ParseInt(rng[1], 10, 32)
		var c = 0
		for _, l := range parts[2] {

			switch l {
			case rune(parts[1][0]):
				c++
			default:
			}
		}
		if int(low) <= c && c <= int(up) {
			pw++
		}
	}
	file.Close()

	fmt.Println(pw)

}

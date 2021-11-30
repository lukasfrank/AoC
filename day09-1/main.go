package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sums(el []int64) (r map[int64]bool) {
	r = make(map[int64]bool)
	for i := 0; i < len(el); i++ {
		for j := 0; j < i; j++ {
			r[el[i]+el[j]] = true
		}
	}

	return r
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := []int64{}
	for scanner.Scan() {
		txt := scanner.Text()
		val, _ := strconv.ParseInt(txt, 10, 64)
		numbers = append(numbers, int64(val))

	}
	file.Close()
	for i := 25; i < len(numbers); i++ {
		if _, ok := sums(numbers[i-25 : i])[numbers[i]]; !ok {
			fmt.Println(numbers[i])
			break
		}
	}
}

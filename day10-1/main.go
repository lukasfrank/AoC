package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := []int{}
	for scanner.Scan() {
		txt := scanner.Text()
		val, _ := strconv.ParseInt(txt, 10, 64)
		numbers = append(numbers, int(val))
	}
	file.Close()

	sort.Ints(numbers)

	dif := make(map[int]int)

	dif[numbers[0]] = 1
	for i, v := range numbers[1:] {
		d := v - numbers[i]
		if _, ok := dif[d]; ok {
			dif[d]++
		} else {
			dif[d] = 1
		}
	}
	dif[3]++
	fmt.Println(dif[3])
	fmt.Println(dif[1])
	fmt.Println(dif[1] * dif[3])
}

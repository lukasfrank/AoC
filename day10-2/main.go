package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	d1 := 0
	res := 0
	for i, v := range numbers[1:] {
		d := v - numbers[i]
		if d == 1 {
			d1++
			if d1 > 1 {
				res++
			}
		} else {
			d1 = 0
		}

		if _, ok := dif[d]; ok {
			dif[d]++
		} else {
			dif[d] = 1
		}
	}
	dif[3]++
	fmt.Println(res)
	fmt.Println(math.Pow(2, float64(res)))
}

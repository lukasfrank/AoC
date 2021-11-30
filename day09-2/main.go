package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	num := int64(0)
	for i := 25; i < len(numbers); i++ {
		if _, ok := sums(numbers[i-25 : i])[numbers[i]]; !ok {
			num = numbers[i]
			break
		}
	}

	l := 0
	r := 0
	sum := int64(numbers[0])
	for r < len(numbers) && l <= r {
		if sum == num {
			break
		} else if sum < num {
			r++
			sum += numbers[r]
		} else {
			sum -= numbers[l]
			l++
		}
	}
	min := int64(math.MaxInt64)
	max := int64(math.MinInt64)
	for _, v := range numbers[l:r] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}

	}
	fmt.Println(min + max)

}

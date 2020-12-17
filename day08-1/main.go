package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type op = func(i, v int) (int, int, bool)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ops := []op{}
	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Split(txt, " ")
		val, _ := strconv.ParseInt(parts[1], 10, 32)
		switch parts[0] {
		case "nop":
			ops = append(ops, func(i, v int) (int, int, bool) { return i + 1, v, false })
		case "acc":
			ops = append(ops, func(i, v int) (int, int, bool) { return i + 1, v + int(val), false })
		case "jmp":
			ops = append(ops, func(i, v int) (int, int, bool) { return i + int(val), v, false })
		default:
		}

	}
	file.Close()

	ptr := 0
	acc := 0
	stop := false

	for {
		ptrCopy := ptr
		ptr, acc, stop = ops[ptr](ptr, acc)
		ops[ptrCopy] = func(i, v int) (int, int, bool) { return i, v, true }
		if stop {
			break
		}
	}
	fmt.Println(acc)

}

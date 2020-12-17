package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type kv struct {
	k string
	v int
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ops := []kv{}
	for scanner.Scan() {
		txt := scanner.Text()
		parts := strings.Split(txt, " ")
		val, _ := strconv.ParseInt(parts[1], 10, 32)
		ops = append(ops, kv{k: parts[0], v: int(val)})
	}
	file.Close()

	ptr := 0
	acc := 0
	seen := make([]bool, len(ops), len(ops))
	done := false

	nextChange := 0
	changed := false

	for !done {
		ptr = 0
		acc = 0
		for k := range seen {
			seen[k] = false
		}
		changed = false
		for i := 0; ; i++ {
			if ptr >= len(ops) {
				done = true
				break
			}
			op := ops[ptr]
			if seen[ptr] {
				break
			}
			seen[ptr] = true

			switch op.k {
			case "nop":
				if i == nextChange && !changed {
					nextChange++
					ptr += op.v
					changed = true
				} else {
					ptr++
				}
			case "acc":
				if i == nextChange && !changed {
					nextChange++
					changed = true
				}
				ptr++
				acc += op.v
			case "jmp":
				if i == nextChange && !changed {
					nextChange++
					ptr++
					changed = true
				} else {
					ptr += op.v
				}

			}

		}
	}
	fmt.Println(acc)

}

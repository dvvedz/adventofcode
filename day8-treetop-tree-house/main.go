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
	f := readFile("sample_input.txt")
	f.pPrint()
}

func readFile(fn string) Trees {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var trees []string

	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, line)
	}
	return trees
}

type Trees []string

func (ts Trees) print() {
	for _, t := range ts {
		fmt.Println(t)
	}
}

func (ts Trees) pPrint() {
	var strees [][]string

	for _, t := range ts {
		sp := strings.Split(t, "")
		strees = append(strees, sp)
	}
	for _, t := range strees {
		fmt.Println(t)
	}
	var sum int
	for i, st := range strees {
		if i != 0 && i != len(st)-1 {
			// midtrees := st[1 : len(st)-1]
			for j, itm := range st {
				if j > 0 && j < len(st)-1 {
					left := st[j-1]
					right := st[j+1]
					top := strees[i-1][j]
					bottom := strees[i+1][j]

					if itm > left || itm > right || itm > top || itm > bottom {
						fmt.Println(i, itm)
						t, _ := strconv.Atoi(itm)
						sum += t
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

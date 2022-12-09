package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fmt.Println(strees)
	for i, st := range strees {

		// Or just add them to New()
		// g := color.New(color.FgGreen, color.Bold)
		// ma := color.New(color.FgMagenta, color.Bold)

		// g.Printf("%s", t[0])

		// mt := t[1 : len(t)-1]
		// for _, m := range mt {
		// ma.Printf("%s", m)
		// }

		// et := t[len(t)-1]
		// g.Printf("%s\n", et)

		// fmt.Printf("%s\n", st)
		if i != 0 && i != len(st)-1 {
			fmt.Println(st[1 : len(st)-1])
		}
	}
}

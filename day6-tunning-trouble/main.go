package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	sli := readFile("input.txt")

	result1 := parser(sli, 4)
	result2 := parser(sli, 14)

	fmt.Println(result1)
	fmt.Println(result2)
}

func readFile(fn string) []string {

	f, err := os.Open(fn)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	var chars []string

	for s.Scan() {
		char := s.Text()
		chars = append(chars, char)
	}
	return chars
}

func parser(sli []string, ran int) int {
	var c2 int = ran
	for c1 := range sli {

		if notDup(sli[c1:c2], ran) == true {
			fmt.Println(sli[c1:c2])
			break
		}
		c2++
	}
	return c2
}

type void struct{}

func notDup(comb []string, ran int) bool {
	set := make(map[string]void)

	for _, co := range comb {
		set[co] = void{}
	}
	t := reflect.ValueOf(set).MapKeys()
	if len(t) <= ran-1 {
		return false
	}
	return true
}

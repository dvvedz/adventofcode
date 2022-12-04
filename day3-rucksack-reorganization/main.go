package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f := openFile("sample_input.txt")
	rucksacks := parseFile(f)
	mapOfItmes := itemPrioList().getMap()

	var valuesAdded int

	for _, r := range rucksacks {
		valuesAdded += mapOfItmes[r.getDupItem()]
	}
	rucksacks.print()

	rucksacks.getGroup()
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)

	if err != nil {
		err := fmt.Errorf("File %v not found", fn)
		fmt.Println(err)
		os.Exit(1)
	}

	return f
}

func parseFile(file *os.File) Rucksacks {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rs Rucksacks

	for scanner.Scan() {
		line := scanner.Text()
		c1 := line[:len(line)/2]
		c2 := line[len(line)/2:]
		if len(line)%2 == 0 {
			r := Rucksack{
				Value:        line,
				Length:       len(line),
				Compartments: map[string]string{"first": c1, "second": c2},
			}
			rs = append(rs, r)
		}
	}
	return rs
}

type Rucksacks []Rucksack

func (rs Rucksacks) print() {
	for _, rucksack := range rs {
		s := `Rucksack value:         %s
Rucksack lenght: 	%d
Rucksack compartments:  %v
-------------------------------------
`
		fmt.Printf(s, rucksack.Value, rucksack.Length, rucksack.Compartments)
	}
}

type Rucksack struct {
	Value        string
	Length       int
	Compartments map[string]string
}

func (rs Rucksack) print() {
	s := `Rucksack value:         %s
Rucksack lenght: 	%d
Rucksack compartments:  %v
-------------------------------------
`
	fmt.Printf(s, rs.Value, rs.Length, rs.Compartments)
}

func (rs Rucksacks) getGroup() {
	var values []string

	for _, v := range rs {
		values = append(values, v.Value)
	}

	var group [][]string
	for i := 0; i < len(values); i += 3 {
		group = append(group, values[i:i+3])
	}

	for _, v := range group {
		// compare each char in first slice with second and third slice
		lstrucksacks := v[1:]
		items := strings.Split(v[0], "")

		fmt.Println("ALL", v)
		fmt.Println("Lasts", lstrucksacks)
		fmt.Println("Items", sort.StringSlice(items))

		for _, lst := range lstrucksacks {
			fmt.Println(lst)
			for _, itm := range items {
				if strings.Contains(lst, itm) {
					fmt.Println(itm)
				}
			}
		}

		// split first backback and compare each item with the two other slices

	}
}

func (rs Rucksack) getDupItem() string {
	// Compare the slices, find the letter that appears in both slices
	fs := rs.Compartments["first"]
	ss := rs.Compartments["second"]

	item := strings.Split(fs, "")
	var dupe string

	for _, v := range item {
		countDup := strings.Count(ss, v)
		if countDup >= 1 {
			dupe = v
			break
		}
	}
	return dupe
}

type Items []Item

func (itms Items) print() {
	for _, itm := range itms {
		fmt.Println(itm.Letter, itm.PrioValue)
	}
}

func (itms Items) getMap() map[string]int {
	//var itemsMap map[string]int
	itemsMap := make(map[string]int)

	for _, itm := range itms {
		itemsMap[itm.Letter] = itm.PrioValue
	}
	return itemsMap
}

func itemPrioList() Items {
	var items Items

	var c1 int = 27
	var c2 int = 1

	// Uppercase
	for i := 65; i <= 90; i++ {
		char := string(rune(i))
		item := Item{
			Letter:    char,
			PrioValue: c1,
		}
		items = append(items, item)
		c1++
	}

	//lowercase
	for i := 97; i <= 122; i++ {
		char := string(rune(i))
		item := Item{
			Letter:    char,
			PrioValue: c2,
		}
		items = append(items, item)
		c2++
	}

	return items
}

type Item struct {
	Letter    string
	PrioValue int
}

func (itm Item) print() {
	fmt.Println(itm.Letter, itm.PrioValue)
}

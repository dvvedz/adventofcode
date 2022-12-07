package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rec := readCsv("input.txt")
	prec := parseCsv(rec)
	var olc int

	var olc2 int

	prec.printNice()

	for _, p := range prec {
		if p.isOverlapping() {
			olc += 1
		}
		if p.isPartiallyOverlapping() {
			olc2 += 1
		}
	}
	fmt.Println(olc2)
}

func readCsv(file string) [][]string {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, rerr := r.ReadAll()
	if rerr != nil {
		log.Fatal(rerr)
	}
	return records
}

func parseCsv(records [][]string) Pairs {
	var rs Pairs
	for _, rec := range records {
		sf := strings.Split(rec[0], "-")
		ss := strings.Split(rec[1], "-")

		var intSlice []int
		var intSlice2 []int

		for _, s := range sf {
			sti, cerr := strconv.Atoi(s)
			if cerr != nil {
				fmt.Println("error when converting string to int")
				return nil
			}

			intSlice = append(intSlice, sti)
		}

		for _, s := range ss {
			sti, cerr := strconv.Atoi(s)
			if cerr != nil {
				fmt.Println("error when converting string to int")
				return nil
			}

			intSlice2 = append(intSlice2, sti)
		}

		cs := map[string][]int{
			"one": intSlice,
			"two": intSlice2,
		}

		as := AssignedSection{
			Pair:  rec,
			Elves: cs,
		}

		rs = append(rs, as)
	}
	return rs
}

type Pairs []AssignedSection

func (pairs Pairs) printNice() {
	for _, p := range pairs {
		p.printNice()
		fmt.Println()
		// debug commment
		// fmt.Println(p.Elves)
	}
}

type AssignedSection struct {
	Pair  []string
	Elves map[string][]int
}

func (as AssignedSection) printNice() {
	// var fos string
	// for _, v := range as.Elves {
	// v = sort.IntSlice(v)
	// for i := 1; i < 100; i++ {
	// fos += "."
	// fos = fos[:i-1] + strconv.Itoa(i) + fos[i:]
	// }
	// fmt.Println(fos, v)
	// fos = "........."
	// }
	var fos string
	var fos2 string
	for i := 1; i < 100; i++ {
		if i >= as.Elves["one"][0] && i <= as.Elves["one"][1] {
			fos += "X"
		} else {
			fos += "."
		}
		if i >= as.Elves["two"][0] && i <= as.Elves["two"][1] {
			fos2 += "X"
		} else {
			fos2 += "."
		}
	}
	fmt.Println(as.Pair)
	fmt.Println(fos)
	fmt.Println(fos2)
}

func (as AssignedSection) isOverlapping() bool {
	// fmt.Println(as.Elves["one"], as.Elves["two"])

	f := as.Elves["one"]
	s := as.Elves["two"]

	if s[0] <= f[0] && f[1] <= s[1] || f[0] <= s[0] && s[1] <= f[1] {
		return true
	}
	return false
}

func (as AssignedSection) isPartiallyOverlapping() bool {
	// fmt.Println(as.Elves["one"], as.Elves["two"])

	f := as.Elves["one"]
	s := as.Elves["two"]
	if s[0] <= f[1] && f[0] <= s[1] {
		return true
	}
	return false
}

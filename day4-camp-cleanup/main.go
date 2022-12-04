package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rec := readCsv("sample_input.txt")
	prec := parseCsv(rec)
	prec.printNice()

	for _, p := range prec {
		p.isOverlapping()
	}
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

type AssignedSection struct {
	Pair  []string
	Elves map[string][]int
}

type Pairs []AssignedSection

func (pairs Pairs) printNice() {
	fos := "........."
	for _, p := range pairs {
		for _, v := range p.Elves {
			v = sort.IntSlice(v)
			for i := v[0]; i <= v[1]; i++ {
				fos = fos[:i-1] + strconv.Itoa(i) + fos[i:]
			}
			fmt.Println(fos)
			fos = "........."
		}
		// debug commment
		fmt.Println()
		// fmt.Println(p.Elves)
	}
}

func (as AssignedSection) isOverlapping() {
	fmt.Println(as.Elves)
}

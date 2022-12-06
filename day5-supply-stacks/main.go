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
	fo := openFile("sample_input.txt")
	ps := parseStacks(fo)
	msc := ps.mapStacksToCrates()

	msc.moveCrates(ps.Moves)
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func parseStacks(f *os.File) Cargo {
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ids []int
	var crates [][]string
	var mov Moves

	for scanner.Scan() {
		line := scanner.Text()
		var trimString string
		// fmt.Println(line)
		if strings.Contains(line, "[") {

			if strings.Contains(line, "    ") {
				trimString = strings.ReplaceAll(line, "    ", " ")
				// fmt.Println(trimString, len(trimString))
			} else {
				trimString = line
			}

			sl := strings.Split(trimString, " ")
			crates = append(crates, sl)
		}

		if strings.HasPrefix(line, " 1") {
			idsTrim := strings.Split(line, "")
			for _, id := range idsTrim {
				if id != " " {
					cs, err := strconv.Atoi(id)
					if err != nil {
						log.Fatal(err)
					}
					ids = append(ids, cs)
				}
			}
		}
		if strings.HasPrefix(line, "move") {
			line = strings.ReplaceAll(line, " ", "")
			line = strings.ReplaceAll(line, "move", "")
			line = strings.ReplaceAll(line, "from", "")
			line = strings.ReplaceAll(line, "to", "")
			nline := strings.Split(line, "")

			var intsli []int
			for _, l := range nline {
				conv, cerr := strconv.Atoi(l)
				if cerr != nil {
					log.Fatal(cerr)
				}
				intsli = append(intsli, conv)
			}

			inst := Instructions{
				Move: intsli[0],
				From: intsli[1],
				To:   intsli[2],
			}

			mov = append(mov, inst)

		}
	}

	sta := Cargo{
		StackIds: ids,
		Crates:   crates,
		Moves:    mov,
	}
	// fmt.Println(crates)
	return sta
}

type Instructions struct {
	Move int
	From int
	To   int
}

type Moves []Instructions

type Cargo struct {
	StackIds []int
	Crates   [][]string
	Moves    Moves
}

func revSlice(rev []string) []string {

	reverse := []string{}

	for i := range rev {
		n := rev[len(rev)-1-i]
		reverse = append(reverse, n)
	}
	return reverse
}

func (c Cargo) mapStacksToCrates() Stacks {
	var sli []Stack
	var s Stack
	for _, v := range c.StackIds {
		var crateAppend []string

		for _, cr := range c.Crates {
			crate := cr[v-1]
			if crate != "" {
				crateAppend = append(crateAppend, crate)
			}
		}
		s = Stack{
			Id:     v,
			Crates: crateAppend,
		}
		sli = append(sli, s)
	}

	return sli
}

type Stack struct {
	Id     int
	Crates []string
}

type Stacks []Stack

func (ss Stacks) moveCrates(mov Moves) {
	// var tmp []Stack
	fmt.Println(ss)
	for _, m := range mov {
		// fmt.Println("From:", m.From, "Move:", m.Move, "to:", m.To)
		move := ss[m.Move-1].Id
		from := ss[m.From-1].Id - 1
		to := ss[m.To-1].Id - 1
		//fmt.Println(ss)
		// fmt.Println("From", from, "to", to, "Move", move)
		fmt.Println("Move", from, "from", move, "to", to)

		var t []string

		t = ss[move].Crates[:move]
		c := revSlice(ss[to].Crates)
		ss[to].Crates = append(c, t[0])
		fmt.Println(ss)
		break
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fc := openFile("calories.txt")
	pf := parseFile(fc)
	tcf := totalCaloriesElvs(pf)

	fmt.Println("Elf with most calories, has", elfeMostCalories(tcf), "calories.")
	fmt.Println("Top 3 elfes combined calories is:", topElvesCalories(tcf, 3))
	// fmt.Println(pf)
}

func openFile(fn string) *os.File {
	// Read file (calories.txt)
	f, err := os.Open(fn)

	if err != nil {
		fmt.Println("Error opening file")
	}

	return f
}

func parseFile(f *os.File) [][]int {
	var elfes [][]int
	var calories []int
	defer f.Close()

	fc := bufio.NewScanner(f)
	for fc.Scan() {
		line := fc.Text()
		n, _ := strconv.Atoi(line)

		if line == "" {
			elfes = append(elfes, calories)
			calories = nil
		} else {
			calories = append(calories, n)
		}
	}

	return elfes
}

func sumOfSlice(elfe []int) int {
	var result int
	for _, calorie := range elfe {
		result += calorie
	}
	return result
}

func totalCaloriesElvs(allElfes [][]int) []int {
	var result []int
	for _, elfe := range allElfes {
		result = append(result, sumOfSlice(elfe))
	}
	return result
}

func elfeMostCalories(elfes []int) int {
	sort.Ints(elfes)

	result := elfes[len(elfes)-1]
	// for _, elfe := range elfes {
	// }
	return result
}

func topElvesCalories(elfes []int, amount int) int {
	sort.Ints(elfes)
	allResults := elfes[len(elfes)-amount:]
	fmt.Println(allResults)
	result := sumOfSlice(allResults)
	return result
}

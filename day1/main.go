package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func toInt(freq string) int {
	i, err := strconv.Atoi(freq)
	if err != nil {
		panic(err)
	}
	return i
}

func part1(start int, filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		freq := scanner.Text()
		start += toInt(freq)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return start

}

func part2(init int, filename string) *int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var nums []int
	seen := map[int]bool{init: true}
	for scanner.Scan() {
		freq := scanner.Text()
		i := toInt(freq)
		nums = append(nums, i)
	}

	current := init
	found := false
	for !found {
		for _, i := range nums {
			current += i
			if _, found := seen[current]; found {
				return &current
			}
			seen[current] = true
		}
	}

	return nil
}

func main() {
	filename := os.Args[1]
	fmt.Println("part1:", part1(0, filename))
	fmt.Println("part2:", *part2(0, filename))
}

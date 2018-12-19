package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func diff(first, second string) (bool, string) {
	var count, diffIndex int
	for i := range first {
		if first[i] != second[i] {
			diffIndex = i
			count++
			if count > 1 {
				return false, ""
			}
		}
	}
	return true, first[:diffIndex] + first[diffIndex+1:]
}

func main() {
	fp, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var data []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	sort.Strings(data)
	for i := 0; i < len(data)-1; i++ {
		first, second := data[i], data[i+1]
		ok, answer := diff(first, second)
		if ok {
			fmt.Println(answer)
		}
	}
}

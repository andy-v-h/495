package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func fourNinetyFive(in int) (int, int) {
	input := sorted(strconv.Itoa(in))

	ordered, _ := strconv.Atoi(input)
	reverse, _ := strconv.Atoi(Reverse(input))

	var (
		count   int
		lastRes int
		res     int
	)
	for {
		res = reverse - ordered
		if res == lastRes {
			break
		}
		// fmt.Printf("%d - %d = %d\n", reverse, ordered, res)
		input = sorted(strconv.Itoa(res))
		ordered, _ = strconv.Atoi(input)
		reverse, _ = strconv.Atoi(Reverse(input))

		lastRes = res
		count++
	}
	return res, count
}

func main() {
	for i := 1; i <= 1000; i++ {
		result, count := fourNinetyFive(i)
		if result != 495 && result != 6174 {
			fmt.Printf("%d took %d step to arrive at %d\n", i, count, result)
		}
	}
}

package main

import (
	"fmt"
	"sort"
)

var (
	prices = []int {146,132,129,178,85,143,1,7,184,139,200,22,42,67,96,72,34,174,156,25,114,178,188,187,46,171,193,149,116,150,27,200,42,100,140,122,187,25,187,67,82,174,37,178,47,158,188,141}
)

func main() {
	pHash := make(map[int][]int, 0)
	for ix, p := range prices {
		if _, ok := pHash[p]; !ok {
			pHash[p] = []int {ix}
		} else {
			pHash[p] = append(pHash[p], ix)
		}
	}
	fmt.Printf("%v\n", pHash)
	sort.Ints(prices)
	fmt.Printf("%v\n", prices)
}

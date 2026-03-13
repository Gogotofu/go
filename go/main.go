package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Function to convert a string into an array of integers
func convertStringToIntArray(m string) []int {
	parts := strings.Split(m, ",")
	result := make([]int, len(parts))
	for i, p := range parts {
		result[i], _ = strconv.Atoi(p)
	}
	return result
}

// Function to count the number of occurrences of each integer
func countNumberFrequency(a []int) map[int]int {
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	return freq
}

// Function to count how many ways there are to make the sum equal to s
func countCardCombinations(a []int, s int) int {
	count := 0
	n := len(a)
	for i := 0; i < (1 << n); i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				sum += a[j]
			}
		}
		if sum == s {
			count++
		}
	}
	
	if count == 0 {
		return -1
	}
	return count
}

// Function to output the key and value of map[int]int in ascending order of key
func printMapKeyAndValue(m map[int]int) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func main() {
	var m string
	var s int

	fmt.Scan(&m)
	fmt.Scan(&s)

	a := convertStringToIntArray(m)
	frequencyCount := countNumberFrequency(a)
	combinationCount := countCardCombinations(a, s)

	printMapKeyAndValue(frequencyCount)
	fmt.Println(combinationCount)
}
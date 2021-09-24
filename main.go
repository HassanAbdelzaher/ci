package main

import (
	"fmt"
	"math"
)

func found(a []string, b string) bool {
	for _, s := range a {
		if s == b {
			return true
		}
	}
	return false
}

func uniqueNames(a, b []string) []string {
	result := []string{}
	for _, s1 := range a {
		if !found(result, s1) {
			result = append(result, s1)
		}
	}
	for _, s1 := range b {
		if !found(result, s1) {
			result = append(result, s1)
		}
	}
	return result
}

func sqr(a float64) float64 {
	return math.Pow(a, 2.0)
}

func main() {
	fmt.Println(sqr(3))
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}

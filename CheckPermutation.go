package main

import (
	"log"
	"sort"
	"strings"
)

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	arr1 := strings.Split(s1, "")
	arr2 := strings.Split(s2, "")

	sort.Strings(arr1)
	sort.Strings(arr2)

	return strings.Join(arr1, "") == strings.Join(arr2, "")
}

func main() {
	s1 := "abc"
	s2 := "bca"
	log.Println(CheckPermutation(s1, s2))
}
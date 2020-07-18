package main

import "fmt"

func kmpSearch(input string, pattern string) []int {
	M := len(input)
	N := len(pattern)
	if M == 0 || N == 0 {
		return nil
	}
	result := make([]int, 0)

	next := make([]int, N)
	next[0] = 0
	i := 1
	l := 0
	for i < N {
		if pattern[i] == pattern[l] {
			l++
			next[i] = l
			i++
		} else {
			if l > 0 {
				l = next[l-1]
			} else {
				next[i] = l
				i++
			}
		}
	}
	next = append([]int{-1}, next...)[:len(next)]
	//fmt.Println(next)

	i = 0
	j := 0
	for i < M {
		if input[i] == pattern[j] {
			i++
			j++
			if j == N-1 {
				result = append(result, i-j)
				j = next[j]
			}
		} else {
			j = next[j]
			if j == -1 {
				i++
				j++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(kmpSearch("what is mind? no matter; what is matter? never mind.",
		"mind"))
}
// https://techdevguide.withgoogle.com/paths/foundational/find-longest-word-in-dictionary-that-subsequence-of-given-string#code-challenge

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/achim-k/go-vebt"
)

// TODO: Generate all possible subsequences, and check each one against the dictionary
// for a match: O(2^n)
func bruteForce(S string, D []string) string {
	return ""
}

// Check each dictionary word using a greedy algorithm: O(N*W)
func greedy(S string, D []string) string {
	// Sort in order of descending length
	sort.Slice(D, func(i, j int) bool {
		return len(D[i]) > len(D[j])
	})

	for _, d := range D {
		d := []rune(d)
		i := 0
		for _, s := range S {
			if s == d[i] {
				i++
				if i >= len(d) {
					break
				}
			}
		}
		// Pass
		if i == len(d) {
			return string(d)
		}
	}

	return ""
}

// Improving the greedy approach: O(N + L * log N)
func greedyMap(S string, D []string) string {
	// Sort in order of descending length
	sort.Slice(D, func(i, j int) bool {
		return len(D[i]) > len(D[j])
	})

	// Build a map of letter
	// S == "abppplee": m['p'] -> [2, 3, 4]
	m := map[rune][]int{}
	for i, c := range S {
		m[c] = append(m[c], i)
	}

	for _, d := range D {
		d := []rune(d)
		i := -1
		for i < len(d) {
			x, ok := m[d[i+1]]
			if !ok {
				break
			}
			// Find the smallest number n > i in x by binary search
			n := sort.Search(len(x), func(t int) bool {
				return x[t] > i
			})
			// Not found
			if n >= len(x) {
				break
			}
			i = x[n]
		}
		// Pass
		if i == len(d) {
			return string(d)
		}
	}

	return ""
}

// Improving the greedy approach with A vEB tree: O(N + L * log log N)
func greedyVEB(S string, D []string) string {
	// Sort in order of descending length
	sort.Slice(D, func(i, j int) bool {
		return len(D[i]) > len(D[j])
	})

	// Build a map of letter
	// S == "abppplee": m['p'] -> [2, 3, 4]
	m := map[rune]*vebt.VEB{}
	for i, c := range S {
		_, ok := m[c]
		if !ok {
			m[c] = vebt.CreateTree(len(S))
			if m[c] == nil {
				panic("vebt.CreateTree failed.")
			}
		}
		m[c].Insert(i)
	}

	for _, d := range D {
		d := []rune(d)
		i := -1
		for i < len(d) {
			x, ok := m[d[i+1]]
			if !ok {
				break
			}
			// Find the smallest number n > i in x
			i = x.Successor(i)
			// Not found
			if i < 0 {
				break
			}
		}
		// Pass
		if i == len(d) {
			return string(d)
		}
	}

	return ""
}

// An optimal O(N + L) approach for small alphabets: O(N + L)
func greedySmallAlphabets(S string, D []string) string {
	// Sort in order of descending length
	sort.Slice(D, func(i, j int) bool {
		return len(D[i]) > len(D[j])
	})

	// Build a dense of letter
	// S == "abppplee": m['p'] -> [2, 2, 3, 4, -1, -1, -1, -1]
	m := map[rune][]int{}
	for i, c := range S {
		_, ok := m[c]
		// []int{-1, -1, ...}
		if !ok {
			m[c] = make([]int, len(S)+1)
			for i := range m[c] {
				m[c][i] = -1
			}
		}
		for j := i - 1; j >= 0 && m[c][j] < 0; j-- {
			m[c][j] = i
		}
	}

	for _, d := range D {
		d := []rune(d)
		i := 0
		for i < len(d) {
			x, ok := m[d[i+1]]
			if !ok {
				break
			}
			i = x[i]
			// Not found
			if i < 0 {
				break
			}
		}
		// Pass
		if i == len(d) {
			return string(d)
		}
	}

	return ""
}

// An optimal O(N + L) approach for any alphabet: O(N + L)
func greedyAnyAlphabets(S string, D []string) string {
	type Pair struct {
		w []rune
		i int
	}

	// S == "abppplee": m['p'] -> [("able", 0), ("ale", 0), ("apple", 0)]
	m := map[rune][]Pair{}
	for _, d := range D {
		d := []rune(d)
		m[d[0]] = append(m[d[0]], Pair{d, 0})
	}

	ans := ""

	for _, c := range S {
		m2 := map[rune][]Pair{}
		for _, p := range m[c] {
			p.i++
			if p.i >= len(p.w) {
				// Update answer
				if len(p.w) > len(ans) {
					ans = string(p.w)
				}
			} else {
				m2[p.w[p.i]] = append(m2[p.w[p.i]], Pair{p.w, p.i})
			}
		}
		// Merge m m2
		for k, v := range m2 {
			m[k] = append(m[k], v...)
		}
	}

	return ans
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: S ...D")
		os.Exit(0)
	}

	S := os.Args[1]
	D := os.Args[2:]

	fmt.Println(greedy(S, D))
}

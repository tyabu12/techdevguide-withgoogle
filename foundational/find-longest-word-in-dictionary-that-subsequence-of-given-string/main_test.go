package main

import "testing"

const (
	S   = "abppplee"
	ans = "apple"
)

var D = []string{"able", "ale", "apple", "bale", "kangaroo"}

// TODO
/*
func TestBruleForce(t *testing.T) {
	out := bruteForce(S, D)
	if out != ans {
		t.Errorf("expected '%s', got '%s'", ans, out)
	}
}
*/

func TestGreedy(t *testing.T) {
	out := greedy(S, D)
	if out != ans {
		t.Errorf("expected '%s', got '%s'", ans, out)
	}
}

func TestGreedyMap(t *testing.T) {
	out := greedyMap(S, D)
	if out != ans {
		t.Errorf("expected '%s', got '%s'", ans, out)
	}
}

func TestGreedyVEB(t *testing.T) {
	out := greedyVEB(S, D)
	if out != ans {
		t.Errorf("expected '%s', got '%s'", ans, out)
	}
}

func TestGreedySmallAlphabets(t *testing.T) {
	out := greedySmallAlphabets(S, D)
	if out != ans {
		t.Errorf("expected '%s', got '%s'", ans, out)
	}
}

func TestGreedyAnyAlphabets(t *testing.T) {
	out := greedyAnyAlphabets(S, D)
	if out != ans {
		t.Errorf("expected '%s', got '%s'", ans, out)
	}
}

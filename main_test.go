package main

import "testing"

func TestCountGoodNodes(t *testing.T) {
	// Classic: 3(root), 3(left-left), 4, 5 are good -> 4.
	r1 := tn(3, tn(1, tn(3, nil, nil), nil), tn(4, tn(1, nil, nil), tn(5, nil, nil)))
	if got := CountGoodNodes(r1); got != 4 {
		t.Errorf("got %d, want 4", got)
	}

	// Descending chain: only the root is good.
	r2 := tn(5, tn(4, tn(3, nil, nil), nil), nil)
	if got := CountGoodNodes(r2); got != 1 {
		t.Errorf("descending = %d, want 1", got)
	}

	// Equal values count as good (>=, not >).
	r3 := tn(2, tn(2, nil, nil), tn(2, nil, nil))
	if got := CountGoodNodes(r3); got != 3 {
		t.Errorf("equal values = %d, want 3", got)
	}

	if got := CountGoodNodes(nil); got != 0 {
		t.Errorf("nil = %d, want 0", got)
	}
	// Negative values: root is always good even below zero.
	if got := CountGoodNodes(tn(-5, nil, nil)); got != 1 {
		t.Errorf("negative single = %d, want 1", got)
	}
}

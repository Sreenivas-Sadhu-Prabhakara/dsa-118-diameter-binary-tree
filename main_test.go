package main

import "testing"

func TestDiameter(t *testing.T) {
	// Path 4-2-5 or 4-2-1-3: longest is 4-2-1-3 with 3 edges.
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, nil, nil))
	if got := Diameter(root); got != 3 {
		t.Errorf("Diameter = %d, want 3", got)
	}

	if got := Diameter(nil); got != 0 {
		t.Errorf("Diameter(nil) = %d, want 0", got)
	}
	if got := Diameter(tn(1, nil, nil)); got != 0 {
		t.Errorf("single node = %d, want 0", got)
	}

	// Diameter not through the root: two deep arms under the left child.
	deep := tn(1,
		tn(2,
			tn(4, tn(6, nil, nil), nil),
			tn(5, nil, tn(7, tn(8, nil, nil), nil))),
		nil)
	// Path 6-4-2-5-7-8 has 5 edges.
	if got := Diameter(deep); got != 5 {
		t.Errorf("deep = %d, want 5", got)
	}
}

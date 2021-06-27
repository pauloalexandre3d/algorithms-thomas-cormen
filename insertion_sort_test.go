package algorithms_thomas_cormen

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := [6]int{5, 2, 4, 6, 1, 3}

	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1

		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}

	r := [6]int{1, 2, 3, 4, 5, 6}
	if a != r {
		t.Fail()
	}
}

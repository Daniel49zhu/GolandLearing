package treesort

import (
	"gopl.io/ch4/treesort"
	"math/rand"
	. "sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !IntsAreSorted(data) {
		t.Errorf("not sorted %v", data)
	}
}

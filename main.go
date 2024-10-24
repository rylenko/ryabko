package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

type P func(u []byte) float64

type QItem struct {
	U []byte
	Q float64
}

func v(items []QItem, u []byte) int {
	for i, item := range items {
		if bytes.Equal(item.U, u) {
			return i
		}
	}
	panic("v not found")
}

type QItemsSort []QItem

func (items QItemsSort) Len() int {
	return len(items)
}

func (items QItemsSort) Less(i, j int) bool {
	return items[i].Q < items[j].Q
}

func (a QItemsSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func computeAllU(n int) [][]byte {
	totalCombs := 1 << n
	combs := make([][]byte, totalCombs)

	for i := 0; i < totalCombs; i++ {
		comb := make([]byte, n)
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				comb[n - j - 1] = 1
			} else {
				comb[n - j - 1] = 0
			}
		}

		combs[i] = comb
	}
	return combs
}

func computePMax(ps []P, u []byte) float64 {
	sup := -math.MaxFloat64

	for _, p := range ps {
		pResult := p(u)
		if pResult > sup {
			sup = pResult
		}
	}

	return sup
}

func computeQ(ps []P, u []byte) float64 {
	return computePMax(ps, u) / computeSP(len(u), ps)
}

func computeOrderedQWithAllU(n int, ps []P) []QItem {
	us := computeAllU(n)
	qs := make([]QItem, len(us))

	for i, u := range us {
		qs[i] = QItem{
			U: u,
			Q: computeQ(ps, u),
		}
	}

	sort.Sort(QItemsSort(qs))
	return qs
}

func computeSP(n int, ps []P) float64 {
	sum := float64(0)
	us := computeAllU(n)

	for _, u := range us {
		sum += computePMax(ps, u)
	}
	return sum
}

func u1(u []byte) float64 {
	return 2 * (float64(u[0]) - float64(u[1]) + float64(u[2]))
}

func u2(u []byte) float64 {
	return 2 * (float64(u[0]) - float64(u[1]) - float64(u[2]))
}

func main() {
	ps := []P{u1, u2}

	fmt.Println(computeAllU(5))
	fmt.Println(computePMax(ps, []byte{1, 1, 1}))
	fmt.Println(computeSP(3, ps))
	fmt.Println(computeQ(ps, []byte{1, 1, 1}))

	items := computeOrderedQWithAllU(3, ps)
	fmt.Println(items)
	fmt.Println(v(items, []byte{1, 0, 1}))
}

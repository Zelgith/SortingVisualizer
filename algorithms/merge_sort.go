package algorithms

import "github.com/Zelgith/SortingVisualizer/utils"

type MergeSort struct{}

func (MergeSort) Sort(slice []any, drawBarChart func([]any, []int)) {
	mergeSort(slice, 0, len(slice)-1, drawBarChart)
}

func mergeSort(slice []any, l, r int, drawBarChart func([]any, []int)) {
	if l < r {
		m := l + (r-l)/2
		mergeSort(slice, l, m, drawBarChart)
		mergeSort(slice, m+1, r, drawBarChart)
		merge(slice, l, m, r, drawBarChart)
	}
}

func merge(slice []any, l, m, r int, drawBarChart func([]any, []int)) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]any, n1)
	R := make([]any, n2)

	for i := 0; i < n1; i++ {
		L[i] = slice[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = slice[m+1+j]
	}

	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if utils.Compare(L[i], R[j]) <= 0 {
			slice[k] = L[i]
			i++
		} else {
			slice[k] = R[j]
			j++
		}
		drawBarChart(slice, []int{k})
		k++
	}

	for i < n1 {
		slice[k] = L[i]
		drawBarChart(slice, []int{k})
		i++
		k++
	}

	for j < n2 {
		slice[k] = R[j]
		drawBarChart(slice, []int{k})
		j++
		k++
	}
}

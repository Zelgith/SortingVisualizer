package algorithms

import "github.com/Zelgith/SortingVisualizer/utils"

type SelectionSort struct{}

func (SelectionSort) Sort(slice []any, drawBarChart func([]any, []int)) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if utils.Compare(slice[j], slice[minIdx]) < 0 {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
		drawBarChart(slice, []int{i, minIdx})
	}
}

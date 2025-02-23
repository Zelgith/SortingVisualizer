package algorithms

import "github.com/Zelgith/SortingVisualizer/utils"

type BubbleSort struct{}

func (BubbleSort) Sort(slice []any, drawBarChart func([]any, []int)) {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if utils.Compare(slice[j], slice[j+1]) > 0 {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				drawBarChart(slice, []int{j, j + 1})
			}
		}
	}
}

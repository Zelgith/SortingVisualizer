package algorithms

import "github.com/Zelgith/SortingVisualizer/utils"

type QuickSort struct{}

func (QuickSort) Sort(slice []any, drawBarChart func([]any, []int)) {
	quickSort(slice, 0, len(slice)-1, drawBarChart)
}

func quickSort(slice []any, low, high int, drawBarChart func([]any, []int)) {
	if low < high {
		pi := partition(slice, low, high, drawBarChart)
		quickSort(slice, low, pi-1, drawBarChart)
		quickSort(slice, pi+1, high, drawBarChart)
	}
}

func partition(slice []any, low, high int, drawBarChart func([]any, []int)) int {
	pivot := slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if utils.Compare(slice[j], pivot) < 0 {
			i++
			slice[i], slice[j] = slice[j], slice[i]
			drawBarChart(slice, []int{i, j})
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	drawBarChart(slice, []int{i + 1, high})
	return i + 1
}

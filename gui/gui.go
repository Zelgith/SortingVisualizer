package gui

import (
	"context"
	"fmt"
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Zelgith/SortingVisualizer/algorithms"
	"github.com/Zelgith/SortingVisualizer/utils"
)

var data []any
var cancelFunc context.CancelFunc
var enableDelay bool

func StartGUI() {
	a := app.New()
	w := a.NewWindow("Sorting Visualizer")

	w.Resize(fyne.NewSize(1000, 800))
	w.CenterOnScreen()

	// Input fields
	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("Enter length of array (max 250)")

	// Dropdown for data type selection
	dataTypes := []string{"Integers", "Floats"}
	dataTypeSelect := widget.NewSelect(dataTypes, func(value string) {})

	// Dropdown for algorithm selection
	algorithmsStr := []string{"Bubble Sort", "Insertion Sort", "Selection Sort", "Merge Sort", "Quick Sort"}
	algorithmSelect := widget.NewSelect(algorithmsStr, func(value string) {})

	// Checkbox for enabling delay
	delayCheckbox := widget.NewCheck("Enable Delay", func(checked bool) {
		enableDelay = checked
	})

	// Canvas for bar chart
	barChart := container.NewWithoutLayout()

	// Function to check if an element is in the step slice
	contains := func(slice []int, element int) bool {
		for _, e := range slice {
			if e == element {
				return true
			}
		}
		return false
	}

	// Function to get the maximum value in the data slice
	getMaxValue := func(data []any) float64 {
		max := float64(0)
		for _, v := range data {
			switch v := v.(type) {
			case int:
				if float64(v) > max {
					max = float64(v)
				}
			case float64:
				if v > max {
					max = v
				}
			}
		}
		return max
	}

	// Function to draw bar chart
	drawBarChart := func(data []any, step []int) {
		barChart.Objects = nil
		width := float32(int(w.Canvas().Size().Width) / len(data))
		maxHeight := w.Canvas().Size().Height * 0.75
		maxValue := getMaxValue(data)
		for i, v := range data {
			var height float32
			switch v := v.(type) {
			case int:
				height = (float32(v) / float32(maxValue)) * maxHeight
			case float64:
				height = (float32(v) / float32(maxValue)) * maxHeight
			}
			height = float32(int(height + 0.5)) // Round to nearest integer
			bar := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
			if contains(step, i) {
				bar.FillColor = color.RGBA{255, 0, 0, 255}
			}
			bar.Resize(fyne.NewSize(width-2, height))
			bar.Move(fyne.NewPos(float32(i)*width, maxHeight-height))
			barChart.Add(bar)
		}
		barChart.Refresh()
		if enableDelay {
			time.Sleep(10 * time.Millisecond)
		}
		if step != nil {
			v := data[step[0]]
			utils.PlayTone(v, maxValue)
		}
	}

	// Function to generate data based on selected type
	generateData := func(length, maxValue int, dataType string) []any {
		switch dataType {
		case "Integers":
			return utils.GenerateIntData(length, maxValue)
		case "Floats":
			return utils.GenerateFloatData(length, maxValue)
		default:
			return nil
		}
	}

	// Label to display sorting time
	timeLabel := widget.NewLabel("Sorting Time: ")

	// Declare startButton variable
	var startButton *widget.Button

	// Initialize startButton
	startButton = widget.NewButton("Start Sorting", func() {
		length, err1 := strconv.Atoi(lengthEntry.Text)
		selectedType := dataTypeSelect.Selected
		selectedAlgorithm := algorithmSelect.Selected

		if err1 != nil || length <= 0 || length > 250 {
			dialog.ShowError(fmt.Errorf("invalid input. Length must be between 1 and 250"), w)
			return
		}
		if selectedType == "" {
			dialog.ShowError(fmt.Errorf("no data type selected"), w)
			return
		}
		if selectedAlgorithm == "" {
			dialog.ShowError(fmt.Errorf("no algorithm selected"), w)
			return
		}

		maxValue := 250
		data = generateData(length, maxValue, selectedType)
		drawBarChart(data, nil)

		var sorter utils.Sorter
		switch selectedAlgorithm {
		case "Bubble Sort":
			sorter = algorithms.BubbleSort{}
		case "Insertion Sort":
			sorter = algorithms.InsertionSort{}
		case "Selection Sort":
			sorter = algorithms.SelectionSort{}
		case "Merge Sort":
			sorter = algorithms.MergeSort{}
		case "Quick Sort":
			sorter = algorithms.QuickSort{}
		}

		sortTemplate := utils.SortTemplate{Slice: &data, Algorithm: sorter, DrawFunc: drawBarChart}
		startButton.Disable()

		// Create a context to manage the goroutine
		ctx, cancel := context.WithCancel(context.Background())
		cancelFunc = cancel

		start := time.Now()
		go func() {
			sortTemplate.SortWithContext(ctx)
			elapsed := time.Since(start)
			var delayStr string
			if enableDelay {
				delayStr = "(huge)"
			} else {
				delayStr = "(small)"
			}
			timeLabel.SetText(fmt.Sprintf("Sorting Time With %s Delay: %s", delayStr, elapsed))
			startButton.Enable()
		}()
	})

	// Button to reset
	resetButton := widget.NewButton("Reset", func() {
		if cancelFunc != nil {
			cancelFunc()
		}
		barChart.Objects = nil
		barChart.Refresh()
		timeLabel.SetText("Sorting Time: ")
		startButton.Enable()
	})

	// Layout
	w.SetContent(container.NewVBox(
		lengthEntry,
		dataTypeSelect,
		algorithmSelect,
		container.NewCenter(
			container.NewHBox(
				delayCheckbox,
				startButton,
				resetButton,
			),
		),
		timeLabel,
		barChart,
	))

	w.ShowAndRun()
}

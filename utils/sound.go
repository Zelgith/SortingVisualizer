package utils

import (
	"math"
	"sync"
	"time"

	"github.com/hajimehoshi/oto"
)

var (
	ctx     *oto.Context
	ctxOnce sync.Once
)

// initContext initializes the oto.Context only once
func initContext() {
	const sampleRate = 44100
	const channelCount = 2
	const bitDepthInBytes = 2
	const bufferSize = 8192

	var err error
	ctx, err = oto.NewContext(sampleRate, channelCount, bitDepthInBytes, bufferSize)
	if err != nil {
		panic(err)
	}
}

// PlayTone plays a tone based on the value of the current step
func PlayTone(value any, maxValue float64) {
	ctxOnce.Do(initContext)

	go func() {
		const sampleRate = 10000
		const duration = 1 * time.Millisecond
		const volume = 0.1
		var frequency float64

		switch v := value.(type) {
		case int:
			frequency = 50.0 + (float64(v)/maxValue)*5.0
		case float64:
			frequency = 50.0 + (v/maxValue)*5.0
		default:
			return
		}

		numSamples := int(sampleRate * duration.Seconds())
		samples := make([]byte, numSamples*2)

		for i := 0; i < numSamples; i++ {
			sample := int16(volume * math.Sin(2*math.Pi*frequency*float64(i)/sampleRate) * 5000)
			samples[2*i] = byte(sample)
			samples[2*i+1] = byte(sample >> 8)
		}

		player := ctx.NewPlayer()
		player.Write(samples)
		player.Close()
	}()
}

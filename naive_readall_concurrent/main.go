package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
)

type pos struct {
	from, to int
}

type result struct {
	city string
	temp float64
}

func main() {
	data := make(map[string]measurements, 10000)
	keys := make([]string, 0, 10000)
	f, _ := os.ReadFile("measurements.txt")

	i := 0

	work := make(chan pos)
	results := make(chan result)
	done := make(chan struct{})
	wg := &sync.WaitGroup{}

	go func() {
		defer close(done)
		for r := range results {
			if _, ok := data[r.city]; !ok {
				data[r.city] = measurements{min: r.temp, max: r.temp, mean: r.temp, numMeasurements: 1}
				keys = append(keys, r.city)
			} else {
				m := data[r.city]
				m.min = math.Min(m.min, r.temp)
				m.max = math.Max(m.max, r.temp)
				m.mean += r.temp
				m.numMeasurements++
				data[r.city] = m
			}
		}
	}()

	for i := 0; i <= 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range work {
				line := f[w.from:w.to]

				delim := bytes.IndexByte(line, ';')
				city := string(line[:delim])

				// todo: int instead of float, or try float32
				temp, _ := strconv.ParseFloat(string(line[delim+1:w.to]), 64)

				results <- result{city: city, temp: temp}
			}
		}()
	}

	newSlice := f
	start := 0
	for lineEnd := bytes.IndexByte(newSlice, '\n'); ; /*lineEnd != -1*/ lineEnd = bytes.IndexByte(newSlice, '\n') {

		i++
		if i == 1000000000 {
			break
		}

		start += lineEnd + 1

		work <- pos{from: start, to: lineEnd}

		newSlice = newSlice[start:]
	}

	close(work)
	wg.Wait()
	close(results)
	<-done

	sort.Strings(keys)

	fmt.Print(" {")
	first := data[keys[0]]
	fmt.Printf("%s=%.1f/%.1f/%.1f", keys[0], first.min, first.mean/float64(first.numMeasurements), first.max)
	for _, k := range keys[1:] {
		m := data[k]
		fmt.Printf(", %s=%.1f/%.1f/%.1f", k, m.min, m.mean/float64(m.numMeasurements), m.max)
	}
	fmt.Print("}\n")
}

type measurements struct {
	min, max, mean  float64
	numMeasurements int
}

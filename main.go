package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	data := make(map[string]measurements, 10000)
	keys := make([]string, 0, 10000)
	f, _ := os.ReadFile("measurements.txt")

	i := 0

	for lineEnd := bytes.IndexByte(f, '\n'); ; /*lineEnd != -1*/ lineEnd = bytes.IndexByte(f, '\n') {
		i++
		if i == 1000000000 {
			break
		}

		delim := bytes.IndexByte(f, ';')
		city := string(f[:delim])

		var temp float64
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("panic")
					fmt.Println("i", i)
					fmt.Println("lineEnd", lineEnd)
					fmt.Println("delim", delim)
					os.Exit(1)
				}
			}()
			// todo: int instead of float, or try float32
			temp, _ = strconv.ParseFloat(string(f[delim+1:lineEnd]), 64)
		}()
		if _, ok := data[city]; !ok {
			data[city] = measurements{min: temp, max: temp, mean: temp, numMeasurements: 1}
			keys = append(keys, city)
		} else {
			m := data[city]
			m.min = math.Min(m.min, temp)
			m.max = math.Max(m.max, temp)
			m.mean += temp
			m.numMeasurements++
			data[city] = m
		}

		f = f[lineEnd+1:]
	}

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

// todo: try tun without GC

package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/emirpasic/gods/v2/maps/treemap"
)

func main() {
	data := treemap.New[string, measurements]()

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

		if m, ok := data.Get(city); !ok {
			data.Put(city, measurements{min: temp, max: temp, mean: temp, numMeasurements: 1})
		} else {
			m.min = math.Min(m.min, temp)
			m.max = math.Max(m.max, temp)
			m.mean += temp
			m.numMeasurements++
			data.Put(city, m)
		}

		f = f[lineEnd+1:]
	}

	iter := data.Iterator()
	iter.First()

	fmt.Print(" {")

	fmt.Printf("%s=%.1f/%.1f/%.1f", iter.Key(), iter.Value().min, iter.Value().mean/float64(iter.Value().numMeasurements), iter.Value().max)
	for iter.Next() {
		m := iter.Value()
		fmt.Printf(", %s=%.1f/%.1f/%.1f", iter.Key(), m.min, m.mean/float64(m.numMeasurements), m.max)
	}
	fmt.Print("}\n")
}

type measurements struct {
	min, max, mean  float64
	numMeasurements int
}

// todo: try tun without GC

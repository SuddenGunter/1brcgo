package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := make(map[string]measurements, 10000)
	keys := make([]string, 0, 10000)
	f, _ := os.OpenFile("measurements.txt", os.O_RDONLY, 0644)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		l := strings.Split(scan.Text(), ";")
		city := l[0]
		temp, _ := strconv.ParseFloat(l[1], 64)

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

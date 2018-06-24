package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Sortable is a wrapper for slice of smth that can be sorted
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Exch(i, j int)
	Get(i int) interface{}
}

// Show print each element of sortable in os.Stdin
func Show(a Sortable) {
	for i := 0; i < a.Len(); i++ {
		fmt.Print(a.Get(i), " ")
	}
	fmt.Println()
}

// IsSorted checks whether sortable is sorted or not
func IsSorted(a Sortable) bool {
	for i := 1; i < a.Len(); i++ {
		if a.Less(i, i-1) {
			return false
		}
	}
	return true
}

func SelectionSort(a Sortable) {
	n := a.Len()
	var min int
	for i := 0; i < n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Exch(i, min)
	}
}

func InsertionSort(a Sortable) {
	n := a.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Exch(j, j-1)
		}
	}
}

func BubbleSort(a Sortable) {
	n := a.Len()
	var exch bool
	for i := n; i > 1; i-- {
		exch = false
		for j := 0; j < i-1; j++ {
			if a.Less(j+1, j) {
				a.Exch(j+1, j)
				exch = true
			}
		}
		if !exch {
			break
		}
	}
}

func ShellSort(a Sortable) {
	N := a.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && a.Less(j, j-h); j -= h {
				a.Exch(j, j-h)
			}
		}
		h = h / 3
	}
}

// Words is a slice of string that satisfied Sortable interface
type Words []string

func (w Words) Len() int              { return len(w) }
func (w Words) Less(i, j int) bool    { return w[i] < w[j] }
func (w Words) Exch(i, j int)         { w[i], w[j] = w[j], w[i] }
func (w Words) Get(i int) interface{} { return w[i] }

type Floats []float64

func (f Floats) Len() int              { return len(f) }
func (f Floats) Less(i, j int) bool    { return f[i] < f[j] }
func (f Floats) Exch(i, j int)         { f[i], f[j] = f[j], f[i] }
func (f Floats) Get(i int) interface{} { return f[i] }

func readStrings() ([]string, error) {
	strs := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return strs, nil
}

func measure(w Words, sortMethod func(Sortable), title string) {
	w2 := make(Words, w.Len())
	copy(w2, w)

	start := time.Now()
	sortMethod(w2)
	fmt.Printf("%s: %s\n", title, time.Now().Sub(start).String())
	if !IsSorted(w2) {
		fmt.Println("Error: not sorted")
	}
}

func measureTime(alg string, a Floats) float64 {
	start := time.Now()
	switch alg {
	case "SelectionSort":
		SelectionSort(a)
	case "InsertionSort":
		InsertionSort(a)
	case "BubbleSort":
		BubbleSort(a)
	case "ShellSort":
		ShellSort(a)
	default:
		panic(alg + " Not implemented")
	}
	if !IsSorted(a) {
		panic(alg + " is not sorted")
	}
	// Show(a)
	return start.Sub(time.Now()).Seconds()
}

func measureRandomFloats(alg string, N, T int) float64 {
	floats := make(Floats, N)
	var total float64
	for t := 0; t < T; t++ {
		for n := 0; n < N; n++ {
			floats[n] = rand.Float64()
		}
		total += measureTime(alg, floats)
	}
	return total
}

func compareAlgs(alg1, alg2 string, N, T int) {
	t1 := measureRandomFloats(alg1, N, T)
	t2 := measureRandomFloats(alg2, N, T)
	fmt.Printf("For %d random floats\n", N)
	fmt.Printf("%s is faster than %s %.3f\n", alg1, alg2, t2/t1)
}

// go run main.go ShellSort InsertionSort 100000 10
func main() {
	// a, err := readStrings()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	return
	// }
	// words := Words(a)
	// // Show(words)
	// fmt.Println("Len: ", words.Len())

	// measure(words, SelectionSort, "SelectionSort")
	// measure(words, InsertionSort, "InsertionSort")
	// measure(words, BubbleSort, "BubbleSort")

	alg1 := os.Args[1]
	alg2 := os.Args[2]
	N, _ := strconv.Atoi(os.Args[3])
	T, _ := strconv.Atoi(os.Args[4])
	compareAlgs(alg1, alg2, N, T)
}

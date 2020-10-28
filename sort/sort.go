package sort

import "math/rand"

// select sort
func SelectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

// insert sort
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j -= 1 {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}

// insert sort
func ShellSort(a []int) {
	h := 1
	// using knuth step
	for h < len(a)/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j-h], a[j] = a[j], a[j-h]

			}
		}

		h = h / 3
	}
}

// merge sort
func MergeSort(a []int) {
	aux := make([]int, len(a))
	sort(a, aux, 0, len(a)-1)
}

func sort(a, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, aux, lo, mid)
	sort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a, aux []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] > aux[i] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}

// add quick sort
func QuickSort(a []int) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := parition(a, lo, hi)
	quickSort(a, lo, j-1)
	quickSort(a, j+1, hi)
}
func parition(a []int, lo, hi int) int {
	i, j := lo+1, hi

	for {
		for i < hi && a[i] < a[lo] {
			i++
		}
		for j > lo && a[j] >= a[lo] {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[j], a[lo] = a[lo], a[j]
	return j
}

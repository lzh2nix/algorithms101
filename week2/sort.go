package week2

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

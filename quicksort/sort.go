package quicksort

func QuickSort(in []int) []int {
	if len(in) == 0 {
		return in
	}

	comp := func(in []int, pivot int, compFunc func(el, pivot int) bool) []int {
		out := make([]int, 0, len(in))
		for _, el := range in {
			if compFunc(el, pivot) {
				out = append(out, el)
			}

		}

		return out
	}

	pivot, res := in[0], in[1:]
	l := comp(res, pivot, func(el, pivot int) bool {
		return el <= pivot
	})

	r := comp(res, pivot, func(el, pivot int) bool {
		return el > pivot
	})

	out := append(QuickSort(l), pivot)
	return append(out, QuickSort(r)...)
}

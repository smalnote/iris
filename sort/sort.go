package sort

// Interface define a sortable collection
type Interface interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

// Sort sort a sortble collection
func Sort(a Interface) {
	quickSort(a, 0, a.Len())
}

func quickSort(a Interface, start int, end int) {
	if start >= end {
		return
	}
	var i, j int
	for i, j = start, start+1; j < end; j++ {
		if a.Less(j, start) {
			i++
			a.Swap(i, j)
		}
	}
	a.Swap(start, i)
	quickSort(a, start, i)
	quickSort(a, i+1, end)
}

type ints []int

func (a ints) Len() int {
	return len(a)
}

func (a ints) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a ints) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Ints sort a int slice
func Ints(a []int) {
	Sort(ints(a))
}

type strings []string

func (a strings) Len() int {
	return len(a)
}

func (a strings) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a strings) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Strings sort a string slice
func Strings(a []string) {
	Sort(strings(a))
}

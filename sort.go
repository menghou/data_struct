package data_struct


type Interface interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func BubbleSort(data Interface) {
	n := data.Len()
	notSorted := true
	for i := 0; i < n && notSorted; i ++ {
		notSorted = false
		for j := n - 1; j > i; j -- {
			if data.Less(j, j-1) {
				data.Swap(j-1,j)
				notSorted = true
			}
		}
	}
}

func SimpleSelectSort(data Interface) {
	for i := 0; i < data.Len(); i ++ {
		min := i
		for j := i; j < data.Len(); j ++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i {
			data.Swap(i,min)
		}
	}
}

func InsertSort(data Interface) {
	for i := 1; i < data.Len(); i ++ {
		for j := i - 1; j >= 0; j -- {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			} else {
				break
			}
		}
	}
}
func ShellSort(data Interface) {
	for increasement := data.Len() / 3 + 1; increasement >= 1; increasement /= 3 {
		for i := increasement; i < data.Len() ; i ++ {
			for j:= i; j >= increasement; j -= increasement {
				if data.Less(j, j-increasement) {
					data.Swap(j, j-increasement)
				} else {
					break
				}
			}
		}
	}
}

func HeapSort(data Interface) {
	heapSort(data, 0, data.Len())
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	for i := (hi - 1) / 2; i >= 0; i ++ {
		heapAdjust(data, lo, hi, first)
	}

	for i := (hi - 1); i >= 0; i ++ {
		data.Swap(first, first + i)
		heapAdjust(data, lo, i, first)
	}
}

func heapAdjust(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := root * 2 + 1
		if child > hi {
			break
		}
		if child + 1 < hi && data.Less(first + child, first + child + 1) {
			child += 1
		}
		if data.Less(first + root, first + child) {
			return
		}
		data.Swap(first + root,first + child)
		root = child
	}
}

func QuickSort(data Interface) {
	if data.Len() < 2 {
		return
	}
	low := 0
	high := data.Len() - 1
	qSort(data ,low, high)
}
func qSort(data Interface, low, high int) {
	var pivotIndex int
	if low < high {
		pivotIndex = PivotIndex(data, low, high)
		qSort(data, low , pivotIndex)
		qSort(data, pivotIndex + 1, high)
	}
}
//返回pivot的索引
func PivotIndex(data Interface, low, high int) int {
	for low < high {
		for low < high && data.Less(low, high) {
			high -= 1
		}
		data.Swap(low, high)
		for low < high && data.Less(low, high) {
			low += 1
		}
		data.Swap(low, high)
	}
	return low
}

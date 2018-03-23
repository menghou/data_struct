package data_struct

import (
	"testing"
	"sort"
)

func TestBubbleSort(t *testing.T) {
	a := []int{9,1,5,8,3,7,4,6,2}
	BubbleSort(sort.IntSlice(a))
	t.Logf("Bubble Sort: %v", a)

	a = []int{9,1,5,8,3,7,4,6,2}
	SimpleSelectSort(sort.IntSlice(a))
	t.Logf("Simple Select Sort: %v", a)
}

func TestSimpleSelectSort(t *testing.T) {
	a := []int{9,1,5,8,3,7,4,6,2}
	SimpleSelectSort(sort.IntSlice(a))
	t.Logf("Simple Select Sort: %v", a)
}

func TestInsertSort(t *testing.T) {
	a := []int{9,1,5,8,3,7,4,6,2}
	InsertSort(sort.IntSlice(a))
	t.Logf("Insert Sort: %v", a)
}
func TestShellSort(t *testing.T) {
	a := []int{9,1,5,8,3,7,4,6,2}
	ShellSort(sort.IntSlice(a))
	t.Logf("Shell Sort: %v", a)
}

func TestPivotIndex(t *testing.T) {
	a := []int{5,1,9,3,7,4,8,6,2}
	QuickSort(sort.IntSlice(a))
}
func TestPivotIndex1(t *testing.T) {
	a := []int{1,2,4,3,5,7,8,6,9}
	PivotIndex(sort.IntSlice(a), 0, 1)
	PivotIndex(sort.IntSlice(a), 2, 8)
}


//func TestQuickSort(t *testing.T) {
//	a := []int{5,1,9,3,7,4,8,6,2}
//	QuickSort(sort.IntSlice(a))
//	t.Logf("Quick Sort: %v", a)
//}

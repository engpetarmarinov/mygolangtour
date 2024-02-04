package main

import (
	"fmt"
	"slices"
	"sort"
	"time"

	"github.com/engpetarmarinov/mygolangtour/sorting"
	"github.com/engpetarmarinov/mygolangtour/utils"
)

func main() {
	slice := utils.GenerateRandomIntSlice(1000000, -1000000, 1000000)
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	slice3 := make([]int, len(slice))
	copy(slice3, slice)
	slice4 := make([]int, len(slice))
	copy(slice4, slice)

	start := time.Now()
	sorting.QuickSort(slice, 0, len(slice)-1)
	quickSortDuration := time.Since(start)

	start = time.Now()
	slice2 = sorting.QuickSortAlloc(slice2)
	quickSortAllocDuration := time.Since(start)

	start = time.Now()
	sort.Ints(slice3)
	sortIntsDuration := time.Since(start)

	start = time.Now()
	slices.Sort(slice4)
	slicesSortDuration := time.Since(start)

	fmt.Println("slice equals slice2:", utils.AreSlicesEqual(slice, slice2))
	fmt.Println("slice2 equals slice3:", utils.AreSlicesEqual(slice2, slice3))
	fmt.Println("slice equals slice3:", utils.AreSlicesEqual(slice, slice3))
	fmt.Println("slice equals slice4:", utils.AreSlicesEqual(slice, slice4))
	fmt.Println("QuickSort took", quickSortDuration)           // QuickSort took 59.456792ms
	fmt.Println("QuickSortAlloc took", quickSortAllocDuration) // QuickSortAlloc took 229.084292ms
	fmt.Println("sort.Ints took", sortIntsDuration)            // sort.Ints took 114.073708ms
	fmt.Println("slices.Sort took", slicesSortDuration)        // slices.Sort took 68.292292ms
}

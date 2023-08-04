package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/wildalmighty/mygolangtour/sorting"
	"github.com/wildalmighty/mygolangtour/utils"
)

func main() {
	slice := utils.GenerateRandomIntSlice(1000000, -1000000, 1000000)
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	slice3 := make([]int, len(slice))
	copy(slice3, slice)

	start := time.Now()
	sorting.QuickSort(slice, 0, len(slice)-1)
	quickSortDuration := time.Since(start)

	start = time.Now()
	slice2 = sorting.QuickSortAlloc(slice2)
	quickSortAllocDuration := time.Since(start)

	start = time.Now()
	sort.Ints(slice3)
	sortIntsDuration := time.Since(start)

	fmt.Println("slice equals slice2:", utils.AreSlicesEqual(slice, slice2))
	fmt.Println("slice2 equals slice3:", utils.AreSlicesEqual(slice2, slice3))
	fmt.Println("slice equals slice3:", utils.AreSlicesEqual(slice, slice3))
	fmt.Println("QuickSort took", quickSortDuration)           // QuickSort took 62.988459ms
	fmt.Println("QuickSortAlloc took", quickSortAllocDuration) // QuickSortAlloc took 227.06875ms
	fmt.Println("sort.Ints took", sortIntsDuration)            // sort.Ints took 119.821875ms
}

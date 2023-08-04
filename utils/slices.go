package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomIntSlice(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}
	return arr
}

func AreSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, value := range slice1 {
		if value != slice2[i] {
			return false
		}
	}

	return true
}

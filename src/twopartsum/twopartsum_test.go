package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test that ConcurrentSum sums an odd-length array correctly
func TestSumConcurrentCorrectlySumsOddArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 45

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test that ConcurrentSum sums unsorted array correctly
func TestSumConcurrentCorrectlySumsUnsortedArray(t *testing.T) {
	arr := []int{1, 5, 6, 3, 420, 69, 21, 1337}
	expected := 1862

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

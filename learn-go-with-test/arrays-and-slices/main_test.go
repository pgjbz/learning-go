package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(numbers []int, got, wanted int, t *testing.T) {
		t.Helper()
		if got != wanted {
			t.Errorf("got %d want %d given %v", got, wanted, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers[:])

		want := 15

		checkSum(numbers[:], got, want, t)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		checkSum(numbers, got, want, t)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {

		got := SumAllTail([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

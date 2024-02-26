package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d wanted %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum and return arrays of values in array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d wanted %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(got []int, want []int, t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test sum all tails with two arguments", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(got, want, t)
	})

	t.Run("test sum all tails with more arguments", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}

		checkSums(got, want, t)

	})
	t.Run("test sum all tails with empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}

		checkSums(got, want, t)
	})
}

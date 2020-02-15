package main

import "testing"

type MultipleTest struct {
	limit    int
	expected int
}

var multipleTests = []MultipleTest{
	{1, 0}, {10, 23}, {1000, 233168},
}

func TestSumMultiplesOf3Or5(t *testing.T) {
	for _, mt := range multipleTests {
		actual := SumMultiplesOf3Or5(mt.limit)
		if actual != mt.expected {
			t.Errorf("SumMultiplesOf3Or5(%d): expected %d, actual %d", mt.limit, mt.expected, actual)
		}
	}
}

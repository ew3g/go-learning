package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	minimum float64
	maximum float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, -1, 1},
	{[]float64{}, 0, 0, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"Expected", pair.average,
				"But got", v)
		}
	}
}

func TestMinimum(t *testing.T) {
	for _, pair := range tests {
		v := Minimum(pair.values)
		if v != pair.minimum {
			t.Error(
				"For", pair.values,
				"Expected", pair.minimum,
				"But got", v)
		}
	}
}

func TestMaximum(t *testing.T) {
	for _, pair := range tests {
		v := Maximum(pair.values)
		if v != pair.maximum {
			t.Error("For", pair.values,
				"Expected", pair.maximum,
				"But got", v)
		}
	}
}

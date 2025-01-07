package main

import "testing"

func TestGetNumberSafe(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		expecteed bool
	}{
		{"Simple Ascending", "1 2 3 4", true},
		{"Ascending At Most 3", "1 4 7 10", true},
		{"Ascending At Most 3; skip first", "99 4 7 10", true},
		{"Ascending At Most 3; skip second", "4 99 7 10", true},
		{"Ascending At Most 3; skip middle", "1 4 99 7 10", true},
		{"Ascending At Most 3; skip last", "1 4 7 10 99", true},

		{"Simple Descending", "4 3 2 1", true},
		{"Descending At Most 3", "10 7 4 1", true},
		{"Ascending At Most 3; skip second", "10 99 7 4", true},
		{"Ascending At Most 3; skip middle", "10 7 99 4 1", true},
		{"Ascending At Most 3; skip last", "10 7 4 1 99", true},

		{"Two Bads last", "10 7 4 1 99 99", false},
		{"Two Bads beginning", "99 99 10 7 4 1", false},
		{"Two Bads middle", "10 7 99 99 4 1", false},

		{"Difference of 4", "20 16 12 8", false},
		{"Difference of 4", "8 12 16 20", false},
		{"Difference of 0", "8 8", false},
		{"Difference of 0", "8 8 8", false},

		{"No Increase", "1 1 1", false},
		{"Increase By More Than 3", "1 5 10", false},
		{"Decrease By More Than 3", "10 5 1", false},

		{"Edge case", "71 69 70 71 72 75", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := isReportSafe(tt.input)
			if err != nil {
				t.Error(err)
			}
			if result != tt.expecteed {
				t.Errorf("getNumberSafe(%v) = %v; want %v", tt.input, result, tt.expecteed)
			}
		})
	}
}

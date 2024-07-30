package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		subjects map[string]float64
		expected float64
	}{
		{
			name: "Test with multiple subjects",
			subjects: map[string]float64{
				"Math":    90.0,
				"Science": 80.0,
				"History": 85.0,
			},
			expected: 85.0,
		},
		{
			name: "Test with a single subject",
			subjects: map[string]float64{
				"Math": 100.0,
			},
			expected: 100.0,
		},
		{
			name:     "Test with no subjects",
			subjects: map[string]float64{},
			expected: 0.0, // Adjust based on your function's behavior
		},
		{
			name: "Test with zero grades",
			subjects: map[string]float64{
				"Math": 0.0,
				"PE":   0.0,
			},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := average(tt.subjects)
			if result != tt.expected {
				t.Errorf("expected %.2f, got %.2f", tt.expected, result)
			}
		})
	}
}

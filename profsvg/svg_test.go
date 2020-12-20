package profsvg

import (
	"math"
	"testing"
)

func TestArcByLevel(t *testing.T) {
	const radius = 5
	circum := circumference(radius)

	testcases := []struct {
		level    int
		expected float32
	}{
		{projectEulerLevelMax, 0.0},
		{0, (1 - 1.0/(projectEulerLevelMax+1)) * circum},
		{-1, circum},
		{projectEulerLevelMax + 1, 0.0},
	}

	for _, tc := range testcases {
		actual := arcByLevel(radius, tc.level)
		if !equalsFloat32(actual, tc.expected) {
			t.Errorf("expected = %v, but actual = %v, level = %v",
				tc.expected, actual, tc.level)
		}
	}
}

func equalsFloat32(a, b float32) bool {
	diff := float64(a - b)
	return math.Abs(diff) < 1e-3
}

func TestLevelPositionX(t *testing.T) {
	testcases := []struct {
		level    int
		expected int
	}{
		{9, 230},
		{10, 220},
	}

	for _, tc := range testcases {
		actual := levelPositionX(tc.level)
		if actual != tc.expected {
			t.Errorf("expected = %v, but actual = %v, level = %v",
				tc.expected, actual, tc.level)
		}
	}
}

package basic

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		x, y, res int
	}{
		{
			x:   3,
			y:   1,
			res: 4,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d", tt.x, tt.y), func(*testing.T) {
			res := Add(tt.x, tt.y)
			if res != tt.res {
				t.Errorf("wanted %d, got %d", tt.res, res)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		x, y, res int
	}{
		{
			x:   3,
			y:   1,
			res: 2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d - %d", tt.x, tt.y), func(*testing.T) {
			res := Subtract(tt.x, tt.y)
			if res != tt.res {
				t.Errorf("wanted %d, got %d", tt.res, res)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		x, y, res int
	}{
		{
			x:   3,
			y:   1,
			res: 3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d * %d", tt.x, tt.y), func(*testing.T) {
			res := Multiply(tt.x, tt.y)
			if res != tt.res {
				t.Errorf("wanted %d, got %d", tt.res, res)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		x, y, res int
	}{
		{
			x:   3,
			y:   1,
			res: 3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d / %d", tt.x, tt.y), func(*testing.T) {
			res := Divide(tt.x, tt.y)
			if res != tt.res {
				t.Errorf("wanted %d, got %d", tt.res, res)
			}
		})
	}
}

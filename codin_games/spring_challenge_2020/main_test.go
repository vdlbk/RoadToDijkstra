package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortCells(t *testing.T) {
	Point := struct {
		X int
		Y int
	}{
		X: 2,
		Y: 3,
	}
	slice := Cells{
		{X: 0, Y: 1, Value: 1, Dist: 0},
		{X: 4, Y: 8, Value: 1, Dist: 0},
		{X: 2, Y: 1, Value: 1, Dist: 0},
		{X: 1, Y: 0, Value: 1, Dist: 0},
		{X: 3, Y: 3, Value: 1, Dist: 0},
		{X: 4, Y: 4, Value: 10, Dist: 0},
		{X: 6, Y: 2, Value: 1, Dist: 0},
		{X: 5, Y: 9, Value: 10, Dist: 0},
	}

	for _, cell := range slice {
		cell.Dist = computeDistance(convert(Point.X, Point.Y, cell.X, cell.Y))
	}

	fmt.Println(slice)

	sort.Sort(slice)
	fmt.Println(slice)

	expectedSlice := Cells{
		{X: 4, Y: 4, Value: 10, Dist: 0},
		{X: 5, Y: 9, Value: 10, Dist: 0},
		{X: 3, Y: 3, Value: 1, Dist: 0},
		{X: 2, Y: 1, Value: 1, Dist: 0},
		{X: 0, Y: 1, Value: 1, Dist: 0},
		{X: 1, Y: 0, Value: 1, Dist: 0},
		{X: 6, Y: 2, Value: 1, Dist: 0},
		{X: 4, Y: 8, Value: 1, Dist: 0},
	}
	fmt.Println(expectedSlice)
}

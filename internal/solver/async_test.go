package solver

import (
	"learning-go-sudoku/internal/solver/dataset"
	"reflect"
	"runtime"
	"testing"
)

var sampleGrid = dataset.Grid{
	{0, 2, 0, 8, 1, 0, 7, 4, 0},
	{7, 0, 0, 0, 0, 3, 1, 0, 0},
	{0, 9, 0, 0, 0, 2, 8, 0, 5},
	{0, 0, 9, 0, 4, 0, 0, 8, 7},
	{4, 0, 0, 2, 0, 8, 0, 0, 3},
	{1, 6, 0, 0, 3, 0, 2, 0, 0},
	{3, 0, 2, 7, 0, 0, 0, 6, 0},
	{0, 0, 5, 6, 0, 0, 0, 0, 8},
	{0, 7, 6, 0, 5, 1, 0, 9, 0},
}

var sampleInvalidGrid = dataset.Grid{
	{0, 2, 0, 8, 1, 0, 7, 4, 0},
	{7, 0, 0, 0, 0, 3, 1, 0, 0},
	{0, 9, 0, 0, 0, 2, 8, 0, 5},
	{0, 0, 9, 0, 4, 0, 0, 8, 7},
	{4, 0, 0, 2, 0, 8, 0, 0, 3},
	{1, 6, 0, 0, 3, 0, 2, 0, 0},
	{3, 0, 2, 7, 0, 0, 0, 6, 0},
	{0, 0, 5, 6, 0, 0, 0, 0, 8},
	{0, 7, 6, 0, 5, 1, 0, 9, 9},
}

var sampleGridSolved = dataset.Grid{
	{5, 2, 3, 8, 1, 6, 7, 4, 9},
	{7, 8, 4, 5, 9, 3, 1, 2, 6},
	{6, 9, 1, 4, 7, 2, 8, 3, 5},
	{2, 3, 9, 1, 4, 5, 6, 8, 7},
	{4, 5, 7, 2, 6, 8, 9, 1, 3},
	{1, 6, 8, 9, 3, 7, 2, 5, 4},
	{3, 4, 2, 7, 8, 9, 5, 6, 1},
	{9, 1, 5, 6, 2, 4, 3, 7, 8},
	{8, 7, 6, 3, 5, 1, 4, 9, 2},
}

func TestSolveAsync(t *testing.T) {
	data := dataset.CopyGrid(sampleGrid)

	res, err := SolveAsync(data, runtime.NumCPU())
	if err != nil {
		t.Fatalf("Error during solving test grid: %q", err.Error())
	}

	if !reflect.DeepEqual(res, sampleGridSolved) {
		t.Error("Sudoku was not solved correctly")
	}
}

func TestSolveAsyncInvalid(t *testing.T) {
	data := dataset.CopyGrid(sampleInvalidGrid)

	_, err := SolveAsync(data, runtime.NumCPU())
	if err == nil {
		t.Fatalf("Did not report invalid grid")
	}
}

func BenchmarkSolveAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		SolveAsync(grid, runtime.NumCPU())
	}
}
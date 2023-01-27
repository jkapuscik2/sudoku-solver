package solver

import (
	"errors"
	"learning-go-sudoku/internal/solver/dataset"
	"reflect"
	"testing"
)

func TestSolveBacktrace(t *testing.T) {
	data := dataset.CopyGrid(sampleGrid)
	res, err := SolveBacktrace(data)
	if err != nil {
		t.Errorf("Error during solving test grid: %q", err.Error())
	}

	if !reflect.DeepEqual(res, sampleGridSolved) {
		t.Error("Sudoku was not solved correctly")
	}
}

func TestSolveBacktraceInvalid(t *testing.T) {
	data := dataset.CopyGrid(sampleInvalidGrid)

	_, err := SolveBacktrace(data)
	if !errors.Is(err, ErrNoSolutions) {
		t.Error("Did not report invalid grid")
	}
}

func TestSolveBacktraceSolved(t *testing.T) {
	data := dataset.CopyGrid(sampleGridSolved)

	res, err := SolveBacktrace(data)
	if !reflect.DeepEqual(res, sampleGridSolved) {
		t.Error("Sudoku was not solved correctly")
	}

	if err != nil {
		t.Errorf("Error during solving test grid: %q", err.Error())
	}
}

func TestSolveBacktraceSolvedInvalid(t *testing.T) {
	data := dataset.CopyGrid(sampleGridSolvedInvalid)

	_, err := SolveBacktrace(data)

	if !errors.Is(err, ErrNoSolutions) {
		t.Errorf("No error during solving an invliad test grid")
	}
}

func BenchmarkTestSolveBacktrace(b *testing.B) {

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		SolveBacktrace(grid)
	}
}

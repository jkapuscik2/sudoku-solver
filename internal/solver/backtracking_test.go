package solver_test

import (
	"errors"
	"github.com/jkapuscik2/sudoku-solver/internal/dataset"
	"github.com/jkapuscik2/sudoku-solver/internal/solver"
	"reflect"
	"testing"
)

func TestSolveBacktrace(t *testing.T) {
	type args struct {
		grid dataset.Grid
	}
	tests := []struct {
		name              string
		args              args
		want              dataset.Grid
		wantErr           bool
		errType           error
		multipleSolutions bool
	}{
		{
			name:    "correct dataset",
			args:    args{grid: sampleGrid},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "unsolvable dataset",
			args:    args{grid: sampleInvalidGrid},
			want:    sampleInvalidGrid,
			wantErr: true,
			errType: solver.ErrNoSolutions,
		},
		{
			name:    "blocked dataset",
			args:    args{grid: sampleGridSolved},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "wrongly solved dataset",
			args:    args{grid: sampleGridSolvedInvalid},
			want:    sampleGridSolvedInvalid,
			wantErr: true,
			errType: solver.ErrNoSolutions,
		},
		{
			name:              "empty dataset",
			args:              args{grid: emptyGrid},
			want:              emptyGrid,
			multipleSolutions: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solver.SolveBacktrace(tt.args.grid)
			if tt.wantErr && err != nil && !errors.Is(err, tt.errType) {
				t.Errorf("SolveBacktrace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.multipleSolutions && !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveBacktrace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTestSolveBacktrace(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		solver.SolveBacktrace(grid)
	}
}

func BenchmarkTestSolveBacktraceSimple(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGridSimple)
		solver.SolveBacktrace(grid)
	}
}

func BenchmarkTestSolveBacktraceEmpty(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(emptyGrid)
		solver.SolveBacktrace(grid)
	}
}

func BenchmarkTestSolveBacktraceHard(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGridHard)
		solver.SolveBacktrace(grid)
	}
}

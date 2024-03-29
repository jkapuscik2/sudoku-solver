package dataset

import "fmt"

const (
	EmptyVal                    = 0
	BoxSize                     = 3
	MinVal                      = 1
	MaxVal                      = 9
	GridSize                    = 9
	ErrInvalidData              = GridError("Invalid data provided")
	ErrTooManyCols              = GridError("Too many columns in dataset")
	ErrTooManyRows              = GridError("Too many rows in dataset")
	ErrDuplicatedValuesInRow    = GridError("There are duplicated values in a row of the Grid")
	ErrDuplicatedValuesInColumn = GridError("There are duplicated values in a column of the Grid")
	ErrDuplicatedValuesInBox    = GridError("There are duplicated values in one of the boxes of the Grid")
	ErrFieldNoExists            = GridError("Filed does not exists")
	ErrIncompleteData           = GridError("Provided dataset is not a valid sudoku grid")
)

type GridError string

func (e GridError) Error() string {
	return string(e)
}

type Position struct {
	X int
	Y int
}

type Grid [GridSize][GridSize]int64

func (dataset Grid) IsFilled() bool {
	for _, arr := range dataset {
		for _, val := range arr {
			if val == EmptyVal {
				return false
			}
		}
	}
	return true
}

func (dataset Grid) Rebuild(pos Position, val int64) Grid {
	rebuilt := CopyGrid(dataset)
	rebuilt[pos.Y][pos.X] = val

	return rebuilt
}

func CopyGrid(matrix Grid) Grid {
	var duplicate Grid

	for i := range matrix {
		duplicate[i] = matrix[i]
	}

	return duplicate
}

func (dataset Grid) GetValue(pos Position) (int64, error) {
	if !has(dataset, pos) {
		return 0, ErrFieldNoExists
	}

	return dataset[pos.Y][pos.X], nil
}

func Validate(dataset Grid) error {
	columns := Grid{}
	for idx, row := range dataset {
		// check if there are duplicated values in rows
		if hasDuplicates(row) {
			return ErrDuplicatedValuesInRow
		}

		// check if cell vales are correct
		for x, item := range row {
			if item != EmptyVal && (item < MinVal || item > MaxVal) {
				return ErrInvalidData
			}

			columns[x][idx] = item
		}
	}

	// check if there are duplicated values in columns
	for _, column := range columns {
		if hasDuplicates(column) {
			return ErrDuplicatedValuesInColumn
		}
	}

	// check if there duplicated values in 3x3 boxes
	for x := 0; x <= GridSize-1; x = x + BoxSize {
		subset := dataset[x : x+BoxSize]

		for y := 0; y <= GridSize-1; y = y + BoxSize {
			var boxVars [GridSize]int64

			for idx, values := range subset {
				columnVales := values[y : y+BoxSize]

				for colIdx, columnVal := range columnVales {
					boxVars[idx*len(columnVales)+colIdx] = columnVal
				}
			}

			if hasDuplicates(boxVars) {
				return ErrDuplicatedValuesInBox
			}
		}
	}

	return nil
}

func PrettyPrint(dataset Grid) {
	for _, row := range dataset {
		fmt.Println(row)
	}
}

func IsEmpty(cell int64) bool {
	return int(cell) == EmptyVal
}

func has(grid Grid, pos Position) bool {
	if len(grid) < pos.Y {
		return false
	}

	for _, row := range grid {
		if len(row) < pos.X {
			return false
		}
	}

	return true
}

func hasDuplicates(row [GridSize]int64) bool {
	checked := make(map[int64]int64, len(row))
	for _, val := range row {
		if _, ok := checked[val]; ok {
			return true
		}

		if val != EmptyVal {
			checked[val] = val
		}
	}

	return false
}

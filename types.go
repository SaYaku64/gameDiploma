package main

// Params - parameters that has everyone
type Params struct {
	ID    uint64
	Size  int
	Name  string
	Allow int // 0 - allowed; 1 - only change; 2 - disallowed
	Type  int
}

// Decoration - decoration object
type Decoration struct {
	Params
	Color string
}

// Building - Building object
type Building struct {
	Params
	Decoration
}

// Territory - Territory object
type Territory struct {
	Params
	Building
}

// Square - tile on map
type Square struct {
	ID          uint64
	Territories []Territory
}

// Count - type for grid
type Count struct {
	Count int
}

// TilesSlice - grid of tiles
var TilesSlice [][]Count
var numbers []Count

// FillSlice - filles grid of tiles
func FillSlice() {
	var rowsSlice []Count
	t := 0

	for a := 1; a <= 10; a++ {

		for row := t + 1; row <= t+12; row++ {
			numbers = append(numbers, Count{row})
			rowsSlice = append(rowsSlice, Count{row})
		}
		TilesSlice = append(TilesSlice, rowsSlice)
		rowsSlice = nil
		t = a*10 + a*2
	}
	TilesSlice = transpose(TilesSlice)
}

func transpose(slice [][]Count) [][]Count {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]Count, yl)
	for i := range result {
		result[i] = make([]Count, xl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[j][i] = numbers[j+i*10]
		}
	}

	return result
}

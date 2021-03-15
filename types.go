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
	Angles []CoordPoint
}

// Count - type for grid
type Count struct {
	Count int
}

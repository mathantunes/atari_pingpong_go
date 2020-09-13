package domain

// KeyboardEvent Representation of a user input on a keyboard
type KeyboardEvent struct {
	Keydown uint8
	Key     Key
}

// Key Represents a generic key
type Key int

const (
	// ArrowUp internal representation of Arrow Up key
	ArrowUp Key = iota
	// ArrowDown internal representation of Arrow Down key
	ArrowDown
	// Space internal representation of Space key
	Space
)

type ScoreEvent struct {
	Left  bool
	Right bool
}

type ScoreNumber map[int][]byte

var (
	PossibleScores = ScoreNumber{
		0: {1, 1, 1,
			1, 0, 1,
			1, 0, 1,
			1, 0, 1,
			1, 1, 1,
		},
		1: {1, 1, 0,
			0, 1, 0,
			0, 1, 0,
			0, 1, 0,
			1, 1, 1},
		2: {
			1, 1, 1,
			0, 0, 1,
			1, 1, 1,
			1, 0, 0,
			1, 1, 1,
		},
		3: {
			1, 1, 1,
			0, 0, 1,
			1, 1, 1,
			0, 0, 1,
			1, 1, 1,
		},
	}
)

const (
	// Width of screen
	Width = 800
	// Height of screen
	Height = 600
)

var (
	// White RGB Color
	White = Color{255, 255, 255}
)

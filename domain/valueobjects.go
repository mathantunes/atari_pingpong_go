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
)

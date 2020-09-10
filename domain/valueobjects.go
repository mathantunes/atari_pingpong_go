package domain

// KeyboardEvent Representation of a user input on a keyboard
type KeyboardEvent struct {
	Keydown uint8
	Key     Key
}

type Key int

const (
	ArrowUp Key = iota
	ArrowDown
)

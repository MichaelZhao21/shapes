package game

// Game states
type GameState int

const (
	LOADING GameState = iota
	MAIN_MENU
	PLAYING
)

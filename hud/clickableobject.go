package hud

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ClickableObject struct {
	// The image to display
	Image *ebiten.Image

	// The position of the object
	X, Y int

	// The size of the object
	Width, Height int

	// The function to call when the object is clicked
	OnClick func()

	// The function to call when the object is hovered over
	OnHover func()

	// The function to call when the object is no longer hovered over
	OnUnhover func()

	// Whether or not the object is currently hovered over
	Hovered bool

	// Whether or not the object is currently clicked
	Clicked bool

	// Whether or not the object is currently visible
	Visible bool

	// Whether or not the object is currently enabled
	Disabled bool
}

// Creates a new ClickableObject
func NewClickableObject(img *ebiten.Image, x, y, width, height int, onClick func(), onHover func(), onUnhover func()) *ClickableObject {
	return &ClickableObject{
		Image:     img,
		X:         x,
		Y:         y,
		Width:     width,
		Height:    height,
		OnClick:   onClick,
		OnHover:   onHover,
		OnUnhover: onUnhover,
		Visible:   true,
		Disabled:  false,
	}
}

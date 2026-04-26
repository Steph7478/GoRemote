package controller

import (
	"github.com/go-vgo/robotgo"
)

type Robot struct{}

func New() *Robot { return &Robot{} }

func (r *Robot) Move(x, y float64)  { robotgo.MoveRelative(int(x), int(y)) }
func (r *Robot) LeftClick()             { robotgo.Click("left") }
func (r *Robot) RightClick()        { robotgo.Click("right") }
func (r *Robot) MouseDown()         { robotgo.MouseDown("left") }
func (r *Robot) MouseUp()           { robotgo.MouseUp("left") }
func (r *Robot) Press(key string)   { robotgo.KeyTap(key) }
func (r *Robot) Type(text string)   { robotgo.Type(text) }
func (r *Robot) Scroll(x, y int)    { robotgo.Scroll(x, y) }

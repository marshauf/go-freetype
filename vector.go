package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type Vector struct {
	handle C.FT_Vector
}

func NewVector() *Vector {
	return &Vector{C.FT_Vector{}}
}

func (v *Vector) X() int64 {
	return int64(v.handle.x)
}

func (v *Vector) Y() int64 {
	return int64(v.handle.y)
}

func (v *Vector) SetX(x int64) { v.handle.x = C.FT_Pos(x) }
func (v *Vector) SetY(y int64) { v.handle.y = C.FT_Pos(y) }

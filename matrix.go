package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type Matrix struct {
	handle C.FT_Matrix
}

func NewMatrix() *Matrix {
	return &Matrix{C.FT_Matrix{}}
}

func (m *Matrix) XX() int64 { return int64(m.handle.xx) }
func (m *Matrix) XY() int64 { return int64(m.handle.xy) }
func (m *Matrix) YX() int64 { return int64(m.handle.yx) }
func (m *Matrix) YY() int64 { return int64(m.handle.yy) }

func (m *Matrix) SetXX(xx int64) { m.handle.xx = C.FT_Fixed(xx) }
func (m *Matrix) SetXY(xy int64) { m.handle.xy = C.FT_Fixed(xy) }
func (m *Matrix) SetYX(yx int64) { m.handle.yx = C.FT_Fixed(yx) }
func (m *Matrix) SetYY(yy int64) { m.handle.yy = C.FT_Fixed(yy) }

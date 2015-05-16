package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
#include FT_OUTLINE_H
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type Outline struct {
	handle C.FT_Outline
}

// number of contours in glyph
func (o *Outline) NumContours() int {
	return int(o.handle.n_contours)
}

// number of points in the glyph
func (o *Outline) NumPoints() int {
	return int(o.handle.n_points)
}

// the outline's points
func (o *Outline) Points() []Vector {
	size := o.NumPoints()
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(o.handle.points)),
		Len:  size,
		Cap:  size,
	}
	return *(*[]Vector)(unsafe.Pointer(&header))
}

// the points flags
func (o *Outline) Tags() []byte {
	size := o.NumPoints()
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(o.handle.tags)),
		Len:  size,
		Cap:  size,
	}
	return *(*[]byte)(unsafe.Pointer(&header))
}

// the contour end points
func (o *Outline) Contours() []int16 {
	size := o.NumContours()
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(o.handle.contours)),
		Len:  size,
		Cap:  size,
	}
	return *(*[]int16)(unsafe.Pointer(&header))
}

// outline masks
func (o *Outline) Flags() int {
	return int(o.handle.flags)
}

// functions

func (o *Outline) Render(library *Library, params *RasterParams) error {
	errno := C.FT_Outline_Render(library.handle, &o.handle, &params.handle)
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

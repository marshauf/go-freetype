package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type SpanFunc func(y, count int, spans []Span)

type Span struct {
	handle C.FT_Span
}

func (s *Span) X() int16 {
	return int16(s.handle.x)
}

func (s *Span) Len() uint16 {
	return uint16(s.handle.len)
}

func (s *Span) Coverage() byte {
	return byte(s.handle.coverage)
}

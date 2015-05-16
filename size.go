package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type Size struct {
	handle C.FT_Size
}

// parent face object
func (s *Size) Face() Face {
	return Face{s.handle.face}
}

// size metrics
func (s *Size) Metrics() SizeMetrics {
	return SizeMetrics{s.handle.metrics}
}

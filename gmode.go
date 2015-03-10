package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type GMode uint32

const (
	GMODE_PAGE_DESCRIPTION GMode = C.HPDF_GMODE_PAGE_DESCRIPTION
	GMODE_PATH_OBJECT      GMode = C.HPDF_GMODE_PATH_OBJECT
	GMODE_TEXT_OBJECT      GMode = C.HPDF_GMODE_TEXT_OBJECT
	GMODE_CLIPPING_PATH    GMode = C.HPDF_GMODE_CLIPPING_PATH
	GMODE_SHADING          GMode = C.HPDF_GMODE_SHADING
	GMODE_INLINE_IMAGE     GMode = C.HPDF_GMODE_INLINE_IMAGE
	GMODE_EXTERNAL_OBJECT  GMode = C.HPDF_GMODE_EXTERNAL_OBJECT
)

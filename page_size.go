package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type PageSize int

const (
	PAGE_SIZE_LETTER    PageSize = C.HPDF_PAGE_SIZE_LETTER
	PAGE_SIZE_LEGAL     PageSize = C.HPDF_PAGE_SIZE_LEGAL
	PAGE_SIZE_A3        PageSize = C.HPDF_PAGE_SIZE_A3
	PAGE_SIZE_A4        PageSize = C.HPDF_PAGE_SIZE_A4
	PAGE_SIZE_A5        PageSize = C.HPDF_PAGE_SIZE_A5
	PAGE_SIZE_B4        PageSize = C.HPDF_PAGE_SIZE_B4
	PAGE_SIZE_B5        PageSize = C.HPDF_PAGE_SIZE_B5
	PAGE_SIZE_EXECUTIVE PageSize = C.HPDF_PAGE_SIZE_EXECUTIVE
	PAGE_SIZE_US4x6     PageSize = C.HPDF_PAGE_SIZE_US4x6
	PAGE_SIZE_US4x8     PageSize = C.HPDF_PAGE_SIZE_US4x8
	PAGE_SIZE_US5x7     PageSize = C.HPDF_PAGE_SIZE_US5x7
	PAGE_SIZE_COMM10    PageSize = C.HPDF_PAGE_SIZE_COMM10
)

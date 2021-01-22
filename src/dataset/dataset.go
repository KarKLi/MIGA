// Package dataset implements the dataset processing layer.
package dataset

// Annoopt interface, provides one primitive.
//
// CheckTotalAnno, check the 'total annotation' file contains correct annotations.
//
// CheckEachAnno, check the 'each annotation' file contains correct annotations.
type Annoopt interface {
	CheckTotalAnno(path string) error
	CheckEachAnno(path string, classLength int) error
}

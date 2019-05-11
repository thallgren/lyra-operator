// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Candela represents a luminous intensity in candela.
type LuminousIntensity float64

const Candela LuminousIntensity = 1

// Unit converts the LuminousIntensity to a *Unit
func (j LuminousIntensity) Unit() *Unit {
	return New(float64(j), Dimensions{
		LuminousIntensityDim: 1,
	})
}

// LuminousIntensity allows LuminousIntensity to implement a LuminousIntensityer interface
func (j LuminousIntensity) LuminousIntensity() LuminousIntensity {
	return j
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (j *LuminousIntensity) From(u Uniter) error {
	if !DimensionsMatch(u, Candela) {
		*j = LuminousIntensity(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*j = LuminousIntensity(u.Unit().Value())
	return nil
}

func (j LuminousIntensity) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", j, float64(j))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " cd"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(j))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(j))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(j))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(j))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g cd)", c, j, float64(j))
	}
}

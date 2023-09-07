package decimal

import (
	gomath "math"
	"strings"

	"github.com/ericlagergren/decimal"
	"github.com/ericlagergren/decimal/math"
)

// Smallest positive number approaching zero in Golang
// It is a very small nonzero number representin needed
// as threshold denominator so we don't divide by zero in AlmostEqual
var epsilon = NewFromFloat64(gomath.Nextafter(1, 2) - 1)

// Cmp compares n to the decimal instance
func (dec Decimal) Cmp(n Decimal) int {
	return dec.native().Cmp(n.native())
}

// Cmp compares a to b
func Cmp(a Decimal, b Decimal) int {
	return a.Cmp(b)
}

// Equals returns true if n has the same value as the decimal instance
func (dec Decimal) Equals(n Decimal) bool {
	return dec.Cmp(n) == 0
}

// Equals returns true if a has the same value as b
func Equals(a Decimal, b Decimal) bool {
	return a.Equals(b)
}

// EqualsInterface returns true if v is an decimal and has the same value as the decimal instance
func (dec Decimal) EqualsInterface(v interface{}) bool {
	switch x := v.(type) {
	case Decimal:
		return dec.Cmp(x) == 0
	case *Decimal:
		return dec.Cmp(*x) == 0
	}
	return false
}

// EqualsInterface returns true if v is an decimal and has the same value as d
func EqualsInterface(d Decimal, v interface{}) bool {
	return d.EqualsInterface(v)
}

// Add adds n to the decimal instance
func (dec Decimal) Add(n Decimal) Decimal {
	dec.native().Add(dec.native(), n.native())
	return Decimal{dec.native()}
}

// Add adds a to b and returns a new decimal instance
// a and b will not be modified
func Add(a Decimal, b Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Add(b)
}

// Sub substracts n to the decimal instance
func (dec Decimal) Sub(n Decimal) Decimal {
	dec.native().Sub(dec.native(), n.native())
	return Decimal{dec.native()}
}

// Sub substracts b from a and returns a new decimal instance
// a and b will not be modified
func Sub(a Decimal, b Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Sub(b)
}

// Div divides n on the decimal instance
func (dec Decimal) Div(n Decimal) Decimal {
	dec.native().Quo(dec.native(), n.native())
	return Decimal{dec.native()}
}

// Div divides b from a and returns a new decimal instance
// a and b will not be modified
func Div(a Decimal, b Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Div(b)
}

// Mul multiplies n to the decimal instance
func (dec Decimal) Mul(n Decimal) Decimal {
	dec.native().Mul(dec.native(), n.native())
	return Decimal{dec.native()}
}

// Mul multiplies a to b and returns a new decimal instance
// a and b will not be modified
func Mul(a Decimal, b Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Mul(b)
}

// Mod modulos n on the decimal instance
func (dec Decimal) Mod(n Decimal) Decimal {
	dec.native().Rem(dec.native(), n.native())
	return Decimal{dec.native()}
}

// Mod modulos b on a and returns a new decimal instance
// a and b will not be modified
func Mod(a Decimal, b Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Mod(b)
}

// Floor rounds the instance down to the next whole number
func (dec Decimal) Floor() Decimal {
	math.Floor(dec.native(), dec.native())
	return Decimal{dec.native()}
}

// Floor rounds d down to the next whole number and returns it as a new instance
// d will not be modified
func Floor(a Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Floor()
}

// Ceil rounds the instance up to the next whole number
func (dec Decimal) Ceil() Decimal {
	math.Ceil(dec.native(), dec.native())
	return Decimal{dec.native()}
}

// Ceil rounds d up to the next whole number and returns it as a new instance
// d will not be modified
func Ceil(a Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Ceil()
}

// Round rounds the instance to the specific digits
func (dec Decimal) Round(digits int) Decimal {
	dec.native().Round(digits)
	return Decimal{dec.native()}
}

// Round rounds d to the specific digits and returns it as a new instance
// d will not be modified
func Round(a Decimal, digits int) Decimal {
	d := NewFromDecimal(a)
	return d.Round(digits)
}

// RoundDown rounds the instance down to the specific digits
func (dec Decimal) RoundDown(digits int) Decimal {
	roundingMode := dec.nat.Context.RoundingMode
	dec.nat.Context.RoundingMode = decimal.ToZero
	dec.native().Round(digits)
	dec.nat.Context.RoundingMode = roundingMode
	return Decimal{dec.native()}
}

// RoundDown rounds d down to the specific digits and returns it as a new instance
// d will not be modified
func RoundDown(a Decimal, digits int) Decimal {
	d := NewFromDecimal(a)
	return d.RoundDown(digits)
}

// RoundToInt rounds the instance to the nearest integer
func (dec Decimal) RoundToInt() Decimal {
	d := dec.native().RoundToInt()
	return Decimal{d}
}

func RoundToInt(a Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.RoundToInt()
}

// Truncate truncates the instance to the specific digits
func (dec Decimal) Truncate(digits int) Decimal {
	parts := strings.SplitN(dec.String(), ".", 2)
	if len(parts) <= 1 {
		v, _ := NewFromString(parts[0])
		dec.native().Copy(v.native())
		return v
	}
	if digits > len(parts[1])-1 {
		digits = len(parts[1])
	}
	v, _ := NewFromString(parts[0] + "." + parts[1][:digits])
	dec.native().Copy(v.native())
	return v
}

// Truncate truncates d to the specific digits and returns it as a new instance
// d will not be modified
func Truncate(a Decimal, digits int) Decimal {
	d := NewFromDecimal(a)
	return d.Truncate(digits)
}

// Quantize sets dec to the number equal in value and sign to dec with the scale, digits.
func (dec Decimal) Quantize(digits int) Decimal {
	dec.native().Quantize(digits)
	return Decimal{dec.native()}
}

// Quantize sets a to the number equal in value and sign to a with the scale, digits.
// a will not be modified.
func Quantize(a Decimal, digits int) Decimal {
	d := NewFromDecimal(a)
	return d.Quantize(digits)
}

// RoundToDigits rounds a to make it have as many digits if possible.
func (dec Decimal) RoundToDigits(digits int) Decimal {
	prec := dec.native().Precision()
	scale := dec.native().Scale()

	// if we have more significant digits (prec) than
	// digits after decimal point (scale) then we want
	// to have the digits after decimal point
	// set to (digits - (prec - scale))
	// for example, 1.23 has prec=3 and scale=2
	// we want to round it to 2 digits
	// so we want the new scale to equal (2 - (3 - 2))=1
	if scale < prec {
		left := prec - scale
		digits = digits - left
	}

	if digits < 0 {
		digits = 0
	}

	dec.native().Quantize(digits)
	return Decimal{dec.native()}
}

// RoundToDigits rounds a to make it have as many digits if possible.
// a will not be modified.
func RoundToDigits(a Decimal, digits int) Decimal {
	d := NewFromDecimal(a)
	return d.RoundToDigits(digits)
}

// Precision returns precision of dec.
func (dec Decimal) Precision() int {
	return dec.native().Precision()
}

func (dec Decimal) Compact() *uint64 {
    compact, _ := decimal.Raw(dec.native())
    return compact
}

// Scale returns scale of dec.
func (dec Decimal) Scale() int {
	return dec.native().Scale()
}

// Abs returns absolute value of a
func Abs(a Decimal) Decimal {
	d := NewFromDecimal(a)
	return d.Abs()
}

// Abs returns absolute value of dec
func (dec Decimal) Abs() Decimal {
	dec.native().Abs(dec.native())
	return Decimal{dec.native()}
}

// Min returns the smallest of a and b
func Min(a, b Decimal) Decimal {
	if a.Cmp(b) <= 0 {
		return Decimal{a.native()}
	}
	return Decimal{b.native()}
}

// Min returns the largest of a and b
func Max(a, b Decimal) Decimal {
	if a.Cmp(b) >= 0 {
		return Decimal{a.native()}
	}
	return Decimal{b.native()}
}

// AlmostEquals checks if n almost equal to dec within a relative tolerance
// Relative tolerance is the ratio of the difference between the two values over the smallest of them
func (dec Decimal) AlmostEquals(n, tolerance Decimal) bool {
	// check if the 2 numbers are exactly equal first
	if dec.Equals(n) {
		return true
	}

	min := Min(dec, n)
	diff := Abs(Sub(dec, n))
	if Abs(min).Equals(Zero()) {
		return diff.Cmp(tolerance) < 0
	}
	return Div(diff, Max(epsilon, min)).Cmp(tolerance) < 0
}

// AlmostEquals checks if a almost equal to b within a relative tolerance
func AlmostEquals(a, b, tolerance Decimal) bool {
	return a.AlmostEquals(b, tolerance)
}

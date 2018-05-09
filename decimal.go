package decimal

import (
	"errors"
	"strings"
	"unicode"

	"github.com/ericlagergren/decimal"
	"github.com/ericlagergren/decimal/math"
)

type Decimal struct {
	native decimal.Big
}

var Zero = New(0, 0)

func New(value int64, scale int32) Decimal {
	return Decimal{
		*decimal.New(value, int(scale)),
	}
}

func NewFromInt(i int) Decimal {
	return NewFromInt64(int64(i))
}

func NewFromInt8(i int8) Decimal {
	return NewFromInt64(int64(i))
}

func NewFromInt16(i int16) Decimal {
	return NewFromInt64(int64(i))
}

func NewFromInt32(i int32) Decimal {
	return NewFromInt64(int64(i))
}

func NewFromInt64(i int64) Decimal {
	return Decimal{
		*decimal.New(i, 0),
	}
}

func NewFromUint(i uint) Decimal {
	return NewFromUint64(uint64(i))
}

func NewFromUint8(i uint8) Decimal {
	return NewFromUint64(uint64(i))
}

func NewFromUint16(i uint16) Decimal {
	return NewFromUint64(uint64(i))
}

func NewFromUint32(i uint32) Decimal {
	return NewFromUint64(uint64(i))
}

func NewFromUint64(i uint64) Decimal {
	d := decimal.New(0, 0)
	d.SetUint64(i)
	return Decimal{*d}
}

func NewFromString(s string) (Decimal, error) {
	if !isDecimal(s) {
		return Decimal{}, errors.New("Invalid decimal")
	}
	d := decimal.New(0, 0)
	_, ok := d.SetString(s)
	if !ok {
		return Decimal{}, errors.New("Invalid decimal")
	}
	return Decimal{*d}, nil
}
func NewFromFloat(f float64) Decimal {
	return NewFromFloat64(f)
}
func NewFromFloat32(f float32) Decimal {
	return NewFromFloat64(float64(f))
}
func NewFromFloat64(f float64) Decimal {
	d := decimal.New(0, 0)
	d.SetFloat64(f)
	return Decimal{
		*d,
	}
}

func (d Decimal) String() string {
	return d.native.String()
}

func (d Decimal) ToInt64() int64 {
	v, _ := d.native.Int64()
	return v
}

func (d Decimal) ToFloat64() float64 {
	v, _ := d.native.Float64()
	return v
}

// IntPart is an alias for ToInt64 that supports the old decimal api
func (d Decimal) IntPart() int64 {
	return d.ToInt64()
}

func (d Decimal) Add(x Decimal) Decimal {
	tmp := decimal.New(0, 0)
	tmp.Add(&d.native, &x.native)
	return Decimal{*tmp}
}

func (d Decimal) Sub(x Decimal) Decimal {
	tmp := decimal.New(0, 0)
	tmp.Sub(&d.native, &x.native)
	return Decimal{*tmp}
}

func (d Decimal) Cmp(x Decimal) int {
	return d.native.Cmp(&x.native)
}
func (d Decimal) Equals(x Decimal) bool {
	return d.Cmp(x) == 0
}

func (d Decimal) Div(x Decimal) Decimal {
	tmp := decimal.New(0, 0)
	tmp.Quo(&d.native, &x.native)
	return Decimal{*tmp}
}

func (d Decimal) Mul(x Decimal) Decimal {
	tmp := decimal.New(0, 0)
	tmp.Mul(&d.native, &x.native)
	return Decimal{*tmp}
}

func (d Decimal) Mod(x Decimal) Decimal {
	tmp := decimal.New(0, 0)
	tmp.Rem(&d.native, &x.native)
	return Decimal{*tmp}
}

func (d Decimal) Floor() Decimal {
	tmp := decimal.New(0, 0)
	math.Floor(tmp, &d.native)
	return Decimal{*tmp}
}

func (d Decimal) Ceil() Decimal {
	tmp := decimal.New(0, 0)
	math.Ceil(tmp, &d.native)
	return Decimal{*tmp}
}

func (d Decimal) Round(p int) Decimal {
	tmp := decimal.New(0, 0)
	tmp.Copy(&d.native)
	tmp.Round(p)
	return Decimal{*tmp}
}
func (d Decimal) Truncate(p int) Decimal {
	parts := strings.SplitN(d.native.String(), ".", 2)
	if len(parts) <= 1 {
		v, _ := NewFromString(parts[0])
		return v
	}
	if p > len(parts[1])-1 {
		p = len(parts[1])
	}
	v, _ := NewFromString(parts[0] + "." + parts[1][:p])
	return v
}

func isDecimal(s string) bool {
	if len(s) <= 0 {
		return false
	}
	runes := []rune(s)

	i := 0
	if runes[0] == '+' || runes[0] == '-' {
		i++
	}

	gotDot := false
	for ; i < len(runes); i++ {
		if runes[i] == '.' {
			if gotDot {
				return false
			}
			gotDot = true
			continue
		}
		if !unicode.IsNumber(runes[i]) {
			return false
		}
	}

	return true
}

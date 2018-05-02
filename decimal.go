package decimal

import (
	"github.com/shopspring/decimal"
)

type Decimal struct {
	native decimal.Decimal
}

var Zero = decimal.Zero

func New(value int64, scale int32) Decimal {
	return Decimal{
		decimal.New(value, scale),
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
		decimal.New(i, 0),
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
	return Decimal{
		decimal.New(int64(i), 0),
	}
}

func NewFromString(s string) (Decimal, error) {
	d, err := decimal.NewFromString(s)
	return Decimal{
		d,
	}, err
}
func NewFromFloat(f float64) Decimal {
	return NewFromFloat64(f)
}
func NewFromFloat32(f float32) Decimal {
	return NewFromFloat64(float64(f))
}
func NewFromFloat64(f float64) Decimal {
	return Decimal{
		decimal.NewFromFloat(f),
	}
}

func (d Decimal) String() string {
	return d.native.String()
}

func (d Decimal) ToInt64() int64 {
	return d.IntPart()
}

func (d Decimal) IntPart() int64 {
	return d.native.IntPart()
}

func (d Decimal) Add(x Decimal) Decimal {
	return Decimal{d.native.Add(x.native)}
}

func (d Decimal) Sub(x Decimal) Decimal {
	return Decimal{d.native.Sub(x.native)}
}

func (d Decimal) Cmp(x Decimal) int {
	return d.native.Cmp(x.native)
}
func (d Decimal) Equals(x Decimal) bool {
	return d.native.Cmp(x.native) == 0
}

func (d Decimal) Div(x Decimal) Decimal {
	return Decimal{d.native.Div(x.native)}
}

func (d Decimal) Mul(x Decimal) Decimal {
	return Decimal{d.native.Mul(x.native)}
}

func (d Decimal) Mod(x Decimal) Decimal {
	return Decimal{d.native.Mod(x.native)}
}

func (d Decimal) Floor() Decimal {
	return Decimal{d.native.Floor()}
}

func (d Decimal) Ceil() Decimal {
	return Decimal{d.native.Ceil()}
}

func (d Decimal) Round(p int) Decimal {
	return Decimal{d.native.Round(int32(p))}
}
func (d Decimal) Truncate(p int) Decimal {
	return Decimal{d.native.Truncate(int32(p))}
}

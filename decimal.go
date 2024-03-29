package decimal

import (
	"errors"
	"fmt"

	"github.com/ericlagergren/decimal"
)

var zero Decimal = Zero()

type Decimal struct {
	nat *decimal.Big
}

func (d *Decimal) native() *decimal.Big {
	if d.nat == nil {
		d.nat = decimal.New(0, 0)
	}
	return d.nat
}

func Zero() Decimal {
	return New(0, 0)
}

func New(value int64, scale int32) Decimal {
	return Decimal{
		decimal.New(value, int(scale)),
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
	d := decimal.New(0, 0)
	d.SetUint64(i)
	return Decimal{d}
}

func NewFromFloat32(f float32) Decimal {
	return NewFromFloat64(float64(f))
}

func NewFromFloat64(f float64) Decimal {
	d := decimal.New(0, 0)
	d.SetFloat64(f)
	return Decimal{d}
}

func NewFromString(s string) (Decimal, error) {
	d := decimal.New(0, 0)
	_, ok := d.SetString(s)
	if !ok {
		return Decimal{}, errors.New("Invalid decimal")
	}
	if d.IsNaN(0) {
		return Decimal{}, errors.New("Invalid decimal")
	}
	return Decimal{d}, nil
}

func MustNewFromString(s string) Decimal {
	d, err := NewFromString(s)
	if err != nil {
		panic(err)
	}
	return d
}

func NewFromDecimal(d Decimal) Decimal {
	if d.nat == nil {
		return Zero()
	}
	cpy := decimal.New(0, 0)
	cpy.Copy(d.nat)
	return Decimal{cpy}
}

func NewFromInterface(value interface{}) (Decimal, error) {
	switch v := value.(type) {
	case float32:
		return NewFromFloat32(v), nil
	case float64:
		return NewFromFloat64(v), nil
	case int:
		return NewFromInt(v), nil
	case int8:
		return NewFromInt8(v), nil
	case int16:
		return NewFromInt16(v), nil
	case int32:
		return NewFromInt32(v), nil
	case int64:
		return NewFromInt64(v), nil
	case uint:
		return NewFromUint(v), nil
	case uint8:
		return NewFromUint8(v), nil
	case uint16:
		return NewFromUint16(v), nil
	case uint32:
		return NewFromUint32(v), nil
	case uint64:
		return NewFromUint64(v), nil
	case string:
		return NewFromString(v)
	case []byte:
		return NewFromString(string(v))
	case Decimal:
		return NewFromDecimal(v), nil
	case *Decimal:
		return NewFromDecimal(*v), nil
	default:
		if buf, ok := v.([]byte); ok {
			return NewFromString(string(buf))
		}
		tmp, err := NewFromString(fmt.Sprintf("%v", v))
		if err != nil {
			return Decimal{}, fmt.Errorf("Unable to create decimal from value type %T: %v", v, err)
		}
		return Decimal{tmp.native()}, nil
	}
}

func MustNewFromInterface(value interface{}) Decimal {
	d, err := NewFromInterface(value)
	if err != nil {
		panic(err)
	}
	return d
}

func (d Decimal) Int8() (int8, error) {
	i, ok := d.native().Int64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an int8", d.String())
	}
	return int8(i), nil
}

func (d Decimal) MustInt8() int8 {
	v, err := d.Int8()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Int16() (int16, error) {
	i, ok := d.native().Int64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an int16", d.String())
	}
	return int16(i), nil
}

func (d Decimal) MustInt16() int16 {
	v, err := d.Int16()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Int32() (int32, error) {
	i, ok := d.native().Int64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an int32", d.String())
	}
	return int32(i), nil
}

func (d Decimal) MustInt32() int32 {
	v, err := d.Int32()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Int64() (int64, error) {
	i, ok := d.native().Int64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an int64", d.String())
	}
	return i, nil
}

func (d Decimal) MustInt64() int64 {
	v, err := d.Int64()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Uint8() (uint8, error) {
	i, ok := d.native().Uint64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an uint8", d.String())
	}
	return uint8(i), nil
}

func (d Decimal) MustUint8() uint8 {
	v, err := d.Uint8()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Uint16() (uint16, error) {
	i, ok := d.native().Uint64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an uint16", d.String())
	}
	return uint16(i), nil
}

func (d Decimal) MustUint16() uint16 {
	v, err := d.Uint16()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Uint32() (uint32, error) {
	i, ok := d.native().Uint64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an uint32", d.String())
	}
	return uint32(i), nil
}

func (d Decimal) MustUint32() uint32 {
	v, err := d.Uint32()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Uint64() (uint64, error) {
	i, ok := d.native().Uint64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an uint64", d.String())
	}
	return i, nil
}

func (d Decimal) MustUint64() uint64 {
	v, err := d.Uint64()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Int() (int, error) {
	i, ok := d.native().Int64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an int", d.String())
	}
	return int(i), nil
}

func (d Decimal) MustInt() int {
	v, err := d.Int()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Uint() (uint, error) {
	i, ok := d.native().Uint64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an uint", d.String())
	}
	return uint(i), nil
}

func (d Decimal) MustUint() uint {
	v, err := d.Uint()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Float32() (float32, error) {
	i, ok := d.native().Float64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an float32", d.String())
	}
	return float32(i), nil
}

func (d Decimal) MustFloat32() float32 {
	v, err := d.Float32()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) Float64() (float64, error) {
	i, ok := d.native().Float64()
	if !ok {
		return 0, fmt.Errorf("`%s' not an float64", d.String())
	}
	return i, nil
}

func (d Decimal) MustFloat64() float64 {
	v, err := d.Float64()
	if err != nil {
		panic(err)
	}
	return v
}

func (d Decimal) String() string {
	if d.native() == nil || d.Equals(zero) {
		return "0"
	}
	return fmt.Sprintf("%f", d)
}

func (d Decimal) Bytes() []byte {
	return []byte(d.String())
}

// Format implements the fmt.Formatter interface.
func (d Decimal) Format(s fmt.State, c rune) {
	d.native().Format(s, c)
}

// IsNaN is a method that wraps IsNaN method of the Big type. Big type method expects an integer argument called quiet. Here's a breakdown of the input values and what they mean:
// - quiet > 0: The function will return true if the Big value x is a quiet NaN.
// - quiet < 0: The function will return true if the Big value x is a signaling NaN.
// - quiet == 0: The function will return true if the Big value x is either a quiet or signaling NaN.
// We are passing 0 as the argument to the IsNaN method, which means that we are checking if the Big value x is either a quiet or signaling NaN.
// For more information, see https://stackoverflow.com/questions/18118408/what-is-the-difference-between-quiet-nan-and-signaling-nan
func (d Decimal) IsNaN() bool {
	return d.native().IsNaN(0)
}

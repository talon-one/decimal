package decimal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZero(t *testing.T) {
	Zero().Add(NewFromInt(1))
	require.NotEqual(t, "1", Zero().String())
}

func TestNewFromDecimal(t *testing.T) {
	d0 := NewFromInt(0)
	d1 := NewFromInt(2)
	d2 := NewFromDecimal(d1)
	require.True(t, d1.Equals(d2))
	d1.native().SetUint64(0)
	require.True(t, d1.Equals(d0))
	require.False(t, d1.Equals(d2))
}

func TestNewFromString(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    string
	}{
		{
			"1", "1", "",
		},
		{
			"1.0", "1.0", "",
		},
		{
			"-1", "-1", "",
		},
		{
			"-1.0", "-1.0", "",
		},
		{
			"1.", "1", "",
		},
		{
			"-1.", "-1", "",
		},
		{
			"", "", "Invalid decimal",
		},
		{
			"1.2.3", "", "Invalid decimal",
		},
		{
			"ABC", "", "Invalid decimal",
		},
	}

	for _, test := range tests {
		d, err := NewFromString(test.input)
		if test.err != "" {
			require.EqualError(t, err, test.err)
			continue
		} else if err != nil {
			require.NoError(t, err)
		}
		require.Equal(t, test.output, d.String())
	}
}

type testStruct struct {
	str string
}

func (a testStruct) String() string {
	return a.str
}

func TestNewFromInterface(t *testing.T) {
	require.Equal(t, float32(1.0), MustNewFromInterface(float32(1.0)).MustFloat32())
	require.Equal(t, float64(1.0), MustNewFromInterface(float64(1.0)).MustFloat64())
	require.Equal(t, int(1), MustNewFromInterface(int(1)).MustInt())
	require.Equal(t, int8(1), MustNewFromInterface(int8(1)).MustInt8())
	require.Equal(t, int16(1), MustNewFromInterface(int16(1)).MustInt16())
	require.Equal(t, int32(1), MustNewFromInterface(int32(1)).MustInt32())
	require.Equal(t, int64(1), MustNewFromInterface(int64(1)).MustInt64())
	require.Equal(t, uint(1), MustNewFromInterface(uint(1)).MustUint())
	require.Equal(t, uint8(1), MustNewFromInterface(uint8(1)).MustUint8())
	require.Equal(t, uint16(1), MustNewFromInterface(uint16(1)).MustUint16())
	require.Equal(t, uint32(1), MustNewFromInterface(uint32(1)).MustUint32())
	require.Equal(t, uint64(1), MustNewFromInterface(uint64(1)).MustUint64())
	require.Equal(t, "1", MustNewFromInterface("1").String())
	require.EqualValues(t, []byte("1"), MustNewFromInterface([]byte("1")).Bytes())
	require.True(t, NewFromInt(1).Equals(MustNewFromInterface(NewFromInt(1))))

	require.Equal(t, "123", MustNewFromInterface(&testStruct{"123"}).String())

	_, err := NewFromInterface(&testStruct{"ABC"})
	require.EqualError(t, err, "Unable to create decimal from value type *decimal.testStruct: Invalid decimal")
}

func TestNilDecimal(t *testing.T) {
	var d Decimal
	require.Equal(t, "0", d.String())
}

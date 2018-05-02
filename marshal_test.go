package decimal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalJSON(t *testing.T) {
	d, err := NewFromString("123.456")
	require.NoError(t, err)
	buf, err := d.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, `123.456`, string(buf))
}

func TestUnmarshalJSON(t *testing.T) {
	var d Decimal
	require.NoError(t, d.UnmarshalJSON([]byte(`"123.456"`)))
	require.Equal(t, "123.456", d.String())
}

func TestMarshalText(t *testing.T) {
	d, err := NewFromString("123.456")
	require.NoError(t, err)
	buf, err := d.MarshalText()
	require.NoError(t, err)
	require.Equal(t, `123.456`, string(buf))
}

func TestUnmarshalText(t *testing.T) {
	var d Decimal
	require.NoError(t, d.UnmarshalJSON([]byte(`123.456`)))
	require.Equal(t, "123.456", d.String())
}

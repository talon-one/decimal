package decimal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToInt64(t *testing.T) {
	d, err := NewFromString("123.456")
	require.NoError(t, err)
	i := d.ToInt64()
	require.Equal(t, "123", fmt.Sprintf("%v", i))
}

func TestToFloat64(t *testing.T) {
	d, err := NewFromString("123.456")
	require.NoError(t, err)
	f := d.ToFloat64()
	require.Equal(t, "123.456", fmt.Sprintf("%v", f))
}

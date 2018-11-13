package decimal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/talon-one/decimal"
)

type TestData struct {
	Decimals              []decimal.Decimal
	StringRepresentations []string
}

func setup(values ...string) (data TestData) {
	data.Decimals = make([]decimal.Decimal, len(values))
	data.StringRepresentations = make([]string, len(values))
	for i, v := range values {
		var err error
		data.Decimals[i], err = decimal.NewFromString(v)
		if err != nil {
			panic(err)
		}
		data.StringRepresentations[i] = data.Decimals[i].String()
	}
	return data
}

func (data *TestData) VerifyIntegrity(t *testing.T) {
	for i := range data.Decimals {
		require.Equal(t, data.StringRepresentations[i], data.Decimals[i].String(), "Decimal on position %d was modified", i)
	}
}

func TestAdd(t *testing.T) {
	data := setup("2", "3")
	require.Equal(t, "5", decimal.Add(data.Decimals[0], data.Decimals[1]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "5", data.Decimals[0].Add(data.Decimals[1]).String())
	data.StringRepresentations[0] = "5"
	data.VerifyIntegrity(t)
}

func TestSub(t *testing.T) {
	data := setup("2", "3")
	require.Equal(t, "-1", decimal.Sub(data.Decimals[0], data.Decimals[1]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "-1", data.Decimals[0].Sub(data.Decimals[1]).String())
	data.StringRepresentations[0] = "-1"
	data.VerifyIntegrity(t)
}

func TestCmp(t *testing.T) {
	data := setup("2", "3")
	require.Equal(t, -1, decimal.Cmp(data.Decimals[0], data.Decimals[1]))
	require.Equal(t, 1, decimal.Cmp(data.Decimals[1], data.Decimals[0]))
	data.VerifyIntegrity(t)

	require.Equal(t, -1, data.Decimals[0].Cmp(data.Decimals[1]))
	require.Equal(t, 1, data.Decimals[1].Cmp(data.Decimals[0]))
	data.VerifyIntegrity(t)
}

func TestEquals(t *testing.T) {
	data := setup("2", "3")
	require.Equal(t, false, decimal.Equals(data.Decimals[0], data.Decimals[1]))
	require.Equal(t, false, decimal.Equals(data.Decimals[1], data.Decimals[0]))
	data.VerifyIntegrity(t)

	require.Equal(t, false, data.Decimals[0].Equals(data.Decimals[1]))
	require.Equal(t, false, data.Decimals[1].Equals(data.Decimals[0]))
	data.VerifyIntegrity(t)

	data = setup("3", "3")
	require.Equal(t, true, decimal.Equals(data.Decimals[0], data.Decimals[1]))
	require.Equal(t, true, decimal.Equals(data.Decimals[1], data.Decimals[0]))
	data.VerifyIntegrity(t)

	require.Equal(t, true, data.Decimals[0].Equals(data.Decimals[1]))
	require.Equal(t, true, data.Decimals[1].Equals(data.Decimals[0]))
	data.VerifyIntegrity(t)
}

func TestEqualsInterface(t *testing.T) {
	data := setup("2", "3")
	require.False(t, decimal.EqualsInterface(data.Decimals[0], data.Decimals[1]))
	require.False(t, decimal.EqualsInterface(data.Decimals[1], data.Decimals[0]))
	require.False(t, decimal.EqualsInterface(data.Decimals[0], &data.Decimals[1]))
	require.False(t, decimal.EqualsInterface(data.Decimals[1], &data.Decimals[0]))
	require.False(t, decimal.EqualsInterface(data.Decimals[1], nil))
	data.VerifyIntegrity(t)

	require.False(t, data.Decimals[0].EqualsInterface(data.Decimals[1]))
	require.False(t, data.Decimals[1].EqualsInterface(data.Decimals[0]))
	require.False(t, data.Decimals[0].EqualsInterface(&data.Decimals[1]))
	require.False(t, data.Decimals[1].EqualsInterface(&data.Decimals[0]))
	require.False(t, data.Decimals[1].EqualsInterface(nil))
	data.VerifyIntegrity(t)

	data = setup("3", "3")
	require.True(t, decimal.EqualsInterface(data.Decimals[0], data.Decimals[1]))
	require.True(t, decimal.EqualsInterface(data.Decimals[1], data.Decimals[0]))
	require.True(t, decimal.EqualsInterface(data.Decimals[0], &data.Decimals[1]))
	require.True(t, decimal.EqualsInterface(data.Decimals[1], &data.Decimals[0]))
	data.VerifyIntegrity(t)

	require.True(t, data.Decimals[0].EqualsInterface(data.Decimals[1]))
	require.True(t, data.Decimals[1].EqualsInterface(data.Decimals[0]))
	require.True(t, data.Decimals[0].EqualsInterface(&data.Decimals[1]))
	require.True(t, data.Decimals[1].EqualsInterface(&data.Decimals[0]))
	data.VerifyIntegrity(t)
}

func TestMul(t *testing.T) {
	data := setup("2", "3")
	require.Equal(t, "6", decimal.Mul(data.Decimals[0], data.Decimals[1]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "6", data.Decimals[0].Mul(data.Decimals[1]).String())
	data.StringRepresentations[0] = "6"
	data.VerifyIntegrity(t)
}

func TestDiv(t *testing.T) {
	data := setup("6", "2")
	require.Equal(t, "3", decimal.Div(data.Decimals[0], data.Decimals[1]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "3", data.Decimals[0].Div(data.Decimals[1]).String())
	data.StringRepresentations[0] = "3"
	data.VerifyIntegrity(t)
}

func TestMod(t *testing.T) {
	data := setup("6", "2")
	require.Equal(t, "0", decimal.Mod(data.Decimals[0], data.Decimals[1]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "0", data.Decimals[0].Mod(data.Decimals[1]).String())
	data.StringRepresentations[0] = "0"
	data.VerifyIntegrity(t)

	data = setup("6", "4")
	require.Equal(t, "2", decimal.Mod(data.Decimals[0], data.Decimals[1]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "2", data.Decimals[0].Mod(data.Decimals[1]).String())
	data.StringRepresentations[0] = "2"
	data.VerifyIntegrity(t)
}

func TestFloor(t *testing.T) {
	data := setup("6.5")
	require.Equal(t, "6", decimal.Floor(data.Decimals[0]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "6", data.Decimals[0].Floor().String())
	data.StringRepresentations[0] = "6"
	data.VerifyIntegrity(t)
}

func TestCeil(t *testing.T) {
	data := setup("6.5")
	require.Equal(t, "7", decimal.Ceil(data.Decimals[0]).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "7", data.Decimals[0].Ceil().String())
	data.StringRepresentations[0] = "7"
	data.VerifyIntegrity(t)
}

func TestRound(t *testing.T) {
	data := setup("6.56")
	require.Equal(t, "6.6", decimal.Round(data.Decimals[0], 2).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "6.6", data.Decimals[0].Round(2).String())
	data.StringRepresentations[0] = "6.6"
	data.VerifyIntegrity(t)
}

func TestTruncate(t *testing.T) {
	data := setup("6.556")
	require.Equal(t, "6.55", decimal.Truncate(data.Decimals[0], 2).String())
	data.VerifyIntegrity(t)

	require.Equal(t, "6.55", data.Decimals[0].Truncate(2).String())
	data.StringRepresentations[0] = "6.55"
	data.VerifyIntegrity(t)

	data = setup("6.55")
	require.Equal(t, "6.55", decimal.Truncate(data.Decimals[0], 4).String())
	data.VerifyIntegrity(t)

	data = setup("6")
	require.Equal(t, "6", decimal.Truncate(data.Decimals[0], 2).String())
	data.VerifyIntegrity(t)
}

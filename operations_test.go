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
	testData := []struct {
		input    string
		digits   int
		expected string
	}{
		{input: "6.56", digits: 2, expected: "6.6"},
		{input: "6.6", digits: 2, expected: "6.6"},
		{input: "-5.684341886E-14", digits: 10, expected: "-0.00000000000005684341886"},
	}
	for i, j := range testData {
		data := setup(j.input)
		output := decimal.Round(data.Decimals[0], j.digits).String()
		require.Equal(t, j.expected, output, "At %d: %s ≠ %s", i, j.expected, output)
		data.VerifyIntegrity(t)
	}
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

func TestRoundToInt(t *testing.T) {
	testData := []struct {
		input    string
		expected string
	}{
		{input: "6.1", expected: "6"},
		{input: "1.0", expected: "1"},
		{input: "10", expected: "10"},
		{input: "1.5", expected: "2"},
	}
	for i, j := range testData {
		data := setup(j.input)
		output := decimal.RoundToInt(data.Decimals[0]).String()
		require.Equal(t, j.expected, output, "At %d: %s ≠ %s", i, j.input, output)
		data.VerifyIntegrity(t)
	}
}

func TestPrecisionAndScale(t *testing.T) {
	testData := []struct {
		input     string
		precision int
		scale     int
	}{
		{input: "6.1", precision: 2, scale: 1},
		{input: "1.0", precision: 2, scale: 1},
		{input: "10", precision: 2, scale: 0},
		{input: "1.5", precision: 2, scale: 1},
		{input: "0.1", precision: 1, scale: 1},
		{input: "1E-10", precision: 1, scale: 10},
		{input: "1.0123456789E10", precision: 11, scale: 0},
		{input: "1.0123456789E9", precision: 11, scale: 1},
		{input: "01.0123456789E9", precision: 11, scale: 1},
		{input: "01.012345678900E9", precision: 13, scale: 3},
		{input: "0.00001", precision: 1, scale: 5},
		{input: "123.45", precision: 5, scale: 2},
	}
	for i, j := range testData {
		data := setup(j.input)
		precision := data.Decimals[0].Precision()
		scale := data.Decimals[0].Scale()
		require.Equal(t, j.precision, precision, "Wrong precision at %d: %d ≠ %d for input %s", i, j.precision, precision, j.input)
		require.Equal(t, j.scale, scale, "Wrong scale at %d: %d ≠ %d for input %s", i, j.scale, scale, j.input)
		data.VerifyIntegrity(t)
	}
}

func TestQuantize(t *testing.T) {
	testData := []struct {
		input    string
		digits   int
		expected string
	}{
		{input: "6.1", digits: 0, expected: "6"},
		{input: "1.0", digits: 2, expected: "1.00"},
		{input: "10", digits: -1, expected: "10"},
		{input: "1.56", digits: 1, expected: "1.6"},
		{input: "6E-2", digits: 2, expected: "0.06"},
		{input: "6E-2", digits: 1, expected: "0.1"},
		{input: "0.00001", digits: 1, expected: "0"},
		{input: "123.44", digits: 1, expected: "123.4"},
		{input: "-5.684341886E-14", digits: 10, expected: "0"},
	}
	for i, j := range testData {
		data := setup(j.input)
		output := decimal.Quantize(data.Decimals[0], j.digits).String()
		require.Equal(t, j.expected, output, "At %d: %s ≠ %s", i, j.expected, output)
		data.VerifyIntegrity(t)
	}
}

func TestRoundToDigits(t *testing.T) {
	testData := []struct {
		input    string
		digits   int
		expected string
	}{
		{input: "6.1", digits: 0, expected: "6"},
		{input: "1.0", digits: 2, expected: "1.0"},
		{input: "1.0", digits: 3, expected: "1.00"},
		{input: "10", digits: -1, expected: "10"},
		{input: "1.56", digits: 1, expected: "2"},
		{input: "6E-2", digits: 2, expected: "0.06"},
		{input: "6E-2", digits: 1, expected: "0.1"},
		{input: "0.00001", digits: 1, expected: "0"},
		{input: "0.00001", digits: 5, expected: "0.00001"},
		{input: "123.44", digits: 1, expected: "123"},
		{input: "123.5", digits: 1, expected: "124"},
		{input: "6.56", digits: 2, expected: "6.6"},
		{input: "6.5", digits: 2, expected: "6.5"},
		{input: "123", digits: 2, expected: "123"},
		{input: "6.2E-3", digits: 5, expected: "0.00620"},
		{input: "6.2E-3", digits: 4, expected: "0.0062"},
		{input: "6.2E-3", digits: 3, expected: "0.006"},
		{input: "6.2E-3", digits: 2, expected: "0.01"},
		{input: "6.2E-3", digits: 1, expected: "0"},
		{input: "0.1", digits: 3, expected: "0.100"},
		{input: "0.1", digits: 2, expected: "0.10"},
		{input: "0.1", digits: 1, expected: "0.1"},
	}
	for i, j := range testData {
		data := setup(j.input)
		output := decimal.RoundToDigits(data.Decimals[0], j.digits).String()
		require.Equal(t, j.expected, output, "At %d: %s ≠ %s", i, j.expected, output)
		data.VerifyIntegrity(t)
	}
}

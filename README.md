# Decimal

The `decimal` package provides an arbitrary-precision floating-point decimal implementation in Go. It represents a number as a value and scale, allowing for precise calculations without losing decimal points.

## Features
- Representation of floating-point numbers with arbitrary precision.
- Support for basic operations such as addition, subtraction, multiplication, etc. (if those methods are defined in your package).
- Utility functions to marshal and unmarshal decimal values.
- Utility functions to serialize and deserialize decimal values.

## Installation

To install the `decimal` package, run the following command:

```bash
go get github.com/talon-one/decimal
```

## Usage

Here's how you can use the decimal package:

Creating a Decimal:

```go
import "github.com/talon-one/decimal"

// Create a new decimal with a value and scale
d := decimal.New(12345, -3) // Represents 12345000

// Create a decimal from an integer
d := decimal.NewFromInt(42) // Represents 42
```

⚠️ Immutability is not supported — changes to a value will affect all its copies.
If an independent copy is required, use `NewFromDecimal()` to create a new value.

```go
d1 := decimal.NewFromInt(100)
d2 := d1
d3 := decimal.NewFromDecimal(d1)

d1.Sub(decimal.NewFromInt(1))

d1.String() // 99 (changed)
d2.String() // 99 (affected)
d3.String() // 100 (independent, not affected)
```

## Credits

Thanks to [ericlagergren/decimal](https://github.com/ericlagergren/decimal) for the underlying decimal representation.
package decimal

import (
	"database/sql/driver"
	"fmt"
)

// Value implements the driver.Valuer interface for database serialization.
func (d Decimal) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan implements the sql.Scanner interface for database deserialization
func (d *Decimal) Scan(value interface{}) error {
	switch v := value.(type) {
	case float32:
		*d = NewFromFloat32(v)
		return nil
	case float64:
		*d = NewFromFloat64(v)
		return nil
	case int:
		*d = NewFromInt(v)
		return nil
	case int8:
		*d = NewFromInt8(v)
		return nil
	case int16:
		*d = NewFromInt16(v)
		return nil
	case int64:
		*d = NewFromInt64(v)
		return nil
	case uint:
		*d = NewFromUint(v)
		return nil
	case uint8:
		*d = NewFromUint8(v)
		return nil
	case uint16:
		*d = NewFromUint16(v)
		return nil
	case uint64:
		*d = NewFromUint64(v)
		return nil
	case string:
		var err error
		*d, err = NewFromString(v)
		return err
	case []byte:
		var err error
		*d, err = NewFromString(string(v))
		return err
	default:
		if buf, ok := v.([]byte); ok {
			var err error
			*d, err = NewFromString(string(buf))
			return err
		}
		tmp, err := NewFromString(fmt.Sprintf("%v", v))
		if err != nil {
			return fmt.Errorf("Unable to scan value type %T: %v", v, err)
		}
		d.native = tmp.native
		return err
	}
}

package decimal

func (d Decimal) Equal(v interface{}) bool {
	switch x := v.(type) {
	case Decimal:
		return d.Cmp(x) == 0
	case *Decimal:
		return d.Cmp(*x) == 0
	}
	return false
}

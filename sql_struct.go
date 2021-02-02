package modules

import (
	"database/sql/driver"
	"github.com/filecoin-project/lotus/chain/types"
)
//自定义结构体，配合lotus转换大数
type NullBigInt struct {
	BigInt     types.BigInt
	Valid      bool
}

func NewNullBigInt(int types.BigInt) NullBigInt {
	return NullBigInt{
		BigInt: int,
		Valid: true,
	}
}

func ZeroNullBigInt() NullBigInt {
	return NullBigInt{
		BigInt: types.NewInt(0),
		Valid: true,
	}
}

// Scan implements the Scanner interface.
func (n *NullBigInt) Scan(value interface{}) error {
	if value == nil {
		n.BigInt, n.Valid = types.NewInt(0), false
		return nil
	}
	n.Valid = true

	//strInt, err := types.BigFromString(fmt.Sprintf("%v", value))

	strInt, err := types.BigFromString(string(value.(string)))
	if err != nil {
		return err
	}
	n.BigInt = strInt
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBigInt) Value() (driver.Value, error) {
	if !n.Valid {
		return "0", nil
	}
	return n.BigInt.String(), nil
}

func (n NullBigInt) String() string {
	if n.Valid {
		return n.BigInt.String()
	}

	return "N/A"
}
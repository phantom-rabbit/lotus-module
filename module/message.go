package module

type Message  struct {
	Common
	Cid        string     `gorm:"index;not null"`

	From       string     `gorm:"index;not null"`
	To         string     `gorm:"index;not null"`
	Nonce      uint64     `gorm:"not null"`
	Value      NullBigInt `gorm:"type:numeric"`
	GasFeeCap  NullBigInt `gorm:"type:numeric"`
	GasPremium NullBigInt `gorm:"type:numeric"`
	GasLimit   int64
	Method     string     `gorm:"index;not null"`
	Version    uint64

	Parameter    []byte

	HasReceipt         bool `gorm:"default:false;"`
	Exit               string
	Error              string
	GasUsed            NullBigInt `gorm:"type:numeric"`
	BaseFeeBurn        NullBigInt `gorm:"type:numeric"`
	OverEstimationBurn NullBigInt `gorm:"type:numeric"`
	MinerPenalty       NullBigInt `gorm:"type:numeric"`
	MinerTip           NullBigInt `gorm:"type:numeric"`
	Refund             NullBigInt `gorm:"type:numeric"`
	TotalCost          NullBigInt `gorm:"type:numeric"`

	PackMiner          string
}
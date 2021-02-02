package module

type Block struct {
	Common

	Cid  	     string `gorm:"uniqueIndex;not null"`
	Parents      []byte

	Miner        string `gorm:"index;not null"`
	WinCount     int64
	MessageNum   int64

	MinedReward  NullBigInt `gorm:"type:numeric"`
	MinedPenalty NullBigInt `gorm:"type:numeric"`
	BaseFeeBurn  NullBigInt `gorm:"type:numeric"`

	ParentStateRoot       string `gorm:"index;not null"`
	ParentWeight          NullBigInt `gorm:"type:numeric"`
	ParentBaseFee         NullBigInt `gorm:"type:numeric"`

	Serialize          []byte
	Size               int64
	ForkSignaling      uint64
}

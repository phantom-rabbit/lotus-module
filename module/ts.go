package module

type Ts struct {
	ID                 uint `gorm:"primarykey"`
	Common
	StateRoot          string

	CirculatingFil     NullBigInt `gorm:"type:numeric"`
	VestedFil          NullBigInt `gorm:"type:numeric"`
	MinedFil           NullBigInt `gorm:"type:numeric"`
	BurntFil           NullBigInt `gorm:"type:numeric"`
	LockedFil          NullBigInt `gorm:"type:numeric"`

	TotalRawBytesPower         NullBigInt `gorm:"type:numeric"`
	TotalQaBytesPower          NullBigInt `gorm:"type:numeric"`

	MinerCount                    uint64
	MinimumConsensusMinerCount    int64
	MsgNum                        int64

	ThisEpochReward        NullBigInt `gorm:"type:numeric"`
	TotalMinedReward       NullBigInt `gorm:"type:numeric"`

	ParentBaseFee          NullBigInt `gorm:"type:numeric"`
}
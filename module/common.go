package module

type Common struct {
	Height             int64 `gorm:"index"`
	Timestamp          uint64
}
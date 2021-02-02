package module

type BlockMessage struct {
	Common
	BlockCid  	 string `gorm:"index;not null"`
	MessageCid   string `gorm:"index;not null"`
}
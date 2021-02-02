package lotus_module

type ActorRecord struct {
	Common

	ActorId      string

	Nonce        uint64
	Balance      NullBigInt `gorm:"type:numeric"`
}
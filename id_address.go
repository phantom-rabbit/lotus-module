package lotus_module

type IdAddressTable struct {
	AddrID            string `gorm:"index;not null;"`
	Address           string `gorm:"index;not null;"`
	AddrType          string

	AddrTag           string
}
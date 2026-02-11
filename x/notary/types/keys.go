package types

const (
	// ModuleName defines the module name
	ModuleName = "notary"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_notary"
)

var (
	ParamsKey = []byte("p_notary")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "jeff"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_jeff"
)

var (
	ParamsKey = []byte("p_jeff")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

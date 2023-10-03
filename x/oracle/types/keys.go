package types

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"

	ReportsKey = "Reports-value-"

	CommitReportKey = "CommitReport-value-"

	ReporterStoreKey = "Reporter-value-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

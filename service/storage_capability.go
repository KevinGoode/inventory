package service

//StorageCapability defines storage capability eg fixcd capacity or dedupe ratio
type StorageCapability struct {
	CapabilityType     string  `cql:"capability_type" json:"capability_type"`
	FixedCapacityBytes int     `cql:"fixed_capacity_bytes" json:"fixed_capacity_bytes"`
	DedupeRatio        float32 `cql:"dedupe_ratio" json:"dedupe_ratio"`
}

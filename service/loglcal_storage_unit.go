package service

//LogicalStorageUnit has comment
type LogicalStorageUnit struct {
	ID              string              `cql:"id" json:"id"`
	StorageDeviceID string              `cql:"storage_device_id" json:"storage_device_id"`
	AssetID         string              `cql:"asset_id" json:"asset_id"`
	Name            string              `cql:"name" json:"name"`
	Status          Status              `cql:"status" json:"status"`
	Tags            []string            `cql:"tags" json:"tags"`
	Capabilities    []StorageCapability `cql:"capabilities" json:"capabilities"`
	UsedBytes       int                 `cql:"used_bytes" json:"used_bytes"`
	LocalManagerIds []LocalManagerID    `cql:"local_manager_ids" json:"local_manager_ids"`
	ChangeRate      int                 `cql:"change_rate" json:"change_rate"`
}

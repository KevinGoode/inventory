package service

//AssetHost has comment
type AssetHost struct {
	ID                    string                `cql:"id" json:"id"`
	Name                  string                `cql:"name" json:"name"`
	SerialNumber          string                `cql:"serial_number" json:"serial_number"`
	OsName                string                `cql:"os_name" json:"os_name"`
	OsVersion             string                `cql:"os_version" json:"os_version"`
	Status                Status                `cql:"status" json:"status"`
	Tags                  []string              `cql:"tags" json:"tags"`
	Capabilities          []AssetHostCapability `cql:"capabilities" json:"capabilities"`
	LocalManagerIds       []LocalManagerID      `cql:"local_manager_ids" json:"local_manager_ids"`
	ConnectedStorageUnits []LogicalStorageUnit  `cql:"connected_storage_units" json:"connected_storage_units"`
}


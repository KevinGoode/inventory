package service

//StorageDevice is container struct for primary and secondary storage devices
type StorageDevice struct {
	ID              string               `cql:"id" json:"id"`
	Name            string               `cql:"name" json:"name"`
	Status          Status               `cql:"status" json:"status"`
	Tags            []string             `cql:"tags" json:"tags"`
	Connections     []NetworkConnection  `cql:"connections" json:"connections"`
	SerialNumber    string               `cql:"serial_number" json:"serial_number"`
	LocalManagerIds []LocalManagerID     `cql:"local_manager_ids" json:"local_manager_ids"`
	AllStorageUnits []LogicalStorageUnit `cql:"all_storage_units" json:"all_storage_units"`
}

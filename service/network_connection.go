package service

//NetworkConnection contains speed and connection type information
type NetworkConnection struct {
	StorageDeviceID string `cql:"storage_device_id" json:"storage_device_id"`
	Speed           int    `cql:"speed" json:"speed"`
	ConnectionType  string `cql:"connection_type" json:"connection_type"`
}

package service

//Asset is an application asset eg sql db
type Asset struct {
	ID                     string               `cql:"id" json:"id"`
	AssetHostID            string               `cql:"asset_host_id" json:"asset_host_id"`
	Name                   string               `cql:"name" json:"name"`
	PolicyID               string               `cql:"policy_id" json:"policy_id"`
	Status                 []Status             `cql:"status" json:"status"`
	Tags                   []string             `cql:"tags" json:"tags"`
	LocalManagerIds        []LocalManagerID     `cql:"local_manager_ids" json:"local_manager_ids"`
	ApplicationVersion     string               `cql:"application_version" json:"application_version"`
	AssociatedStorageUnits []LogicalStorageUnit `cql:"associated_storage_units" json:"associated_storage_units"`
}

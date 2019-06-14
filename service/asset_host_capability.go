package service

//AssetHostCapability is software (namne,version,status) that is installed on asset host
type AssetHostCapability struct {
	Name    string `cql:"name" json:"name"`
	Version string `cql:"version" json:"version"`
	Status  Status `cql:"status" json:"status"`
}

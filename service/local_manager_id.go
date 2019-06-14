package service

//LocalManagerID contains
type LocalManagerID struct {
	ManagerID      string `cql:"manager_id" json:"manager_id"`
	ID             string `cql:"id" json:"id"`
	IsConfigured   bool   `cql:"is_configured" json:"is_configured"`
	IsConfigurable bool   `cql:"is_configurable" json:"is_configurable"`
}

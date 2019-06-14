package service

//Status contains a name value and criticality. NB need tags for data binding to work in parent structs
type Status struct {
	Name        string `cql:"name" json:"name"`
	Value       string `cql:"value" json:"value"`
	Criticality string `cql:"criticality" json:"criticality"`
}

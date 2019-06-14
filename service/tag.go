package service

//Tag tags an entity
type Tag struct {
	Name        string   `cql:"name" json:"name"`
	Description string   `cql:"description" json:"description"`
	Parents     []string `cql:"parents" json:"parents"`
	Scope       []string `cql:"scope" json:"scope"`
}

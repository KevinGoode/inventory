package service

var factory DatabaseFactoryAPI

//DatabaseFactory is resonsible for getting/creating database
type DatabaseFactory struct {
	dbDictionary map[string]DatabaseAPI
	creater      DatabaseCreaterAPI
}

//GetDatabase returns a database api
func (fac *DatabaseFactory) GetDatabase(customerID string) DatabaseAPI {
	var ok = false
	var api DatabaseAPI
	api, ok = fac.dbDictionary[customerID]
	if ok != true {
		api = fac.creater.CreateDatabase(customerID)
		fac.dbDictionary[customerID] = api
	}
	return api
}

//DisposeAll disposes all databases
func (fac *DatabaseFactory) DisposeAll() {
	for _, db := range fac.dbDictionary {
		db.Dispose()
	}
	//Delete old cache
	fac.dbDictionary = make(map[string]DatabaseAPI)
}

//NewDatabaseFactory constructs singleton database factory
func NewDatabaseFactory(createrAPI DatabaseCreaterAPI) DatabaseFactoryAPI {
	if factory == nil {
		concreteFac := DatabaseFactory{}
		concreteFac.dbDictionary = make(map[string]DatabaseAPI)
		concreteFac.creater = createrAPI
		factory = &concreteFac
	}
	return factory
}

package service

//DatabaseFactoryAPI is main factory api
type DatabaseFactoryAPI interface {
	GetDatabase(customerID string) DatabaseAPI
	DisposeAll()
}

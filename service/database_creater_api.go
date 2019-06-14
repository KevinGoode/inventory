package service

//DatabaseCreaterAPI is main database api to execute command
type DatabaseCreaterAPI interface {
	CreateDatabase(CustomerID string) DatabaseAPI
}

package service

//ParameterReaderAPI defines api to get input params
type ParameterReaderAPI interface {
	Process()
	GetReadParameter() string
	GetWriteParameter() string
	GetCustomerParameter() string
	GetWriteParameters() string
}

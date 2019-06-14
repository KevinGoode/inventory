package service

//DummyParameterReader is container for input params
type DummyParameterReader struct {
	Read       string
	Write      string
	Customer   string
	Parameters string
}

//Process processes input params
func (reader *DummyParameterReader) Process() {

}

//GetReadParameter gets read parameter. (Call Process first)
func (reader DummyParameterReader) GetReadParameter() string {
	return reader.Read
}

//GetWriteParameter gets write parameter. (Call Process first)
func (reader DummyParameterReader) GetWriteParameter() string {
	return reader.Write
}

//GetCustomerParameter customer read parameter. (Call Process first)
func (reader DummyParameterReader) GetCustomerParameter() string {
	return reader.Customer
}

//GetWriteParameters gets write parameters. (Call Process first)
func (reader DummyParameterReader) GetWriteParameters() string {
	return reader.Parameters
}

//NewDummyReader creates a new parameter reader
func NewDummyReader() ParameterReaderAPI {
	reader := DummyParameterReader{}
	return &reader
}

package service

import (
	"flag"
)

//ParameterReader is container for input params
type ParameterReader struct {
	read       string
	write      string
	customer   string
	parameters string
}

//Process processes input params
func (reader *ParameterReader) Process() {
	flag.StringVar(&reader.read, "read", "", ReadHelp)
	flag.StringVar(&reader.read, "r", "", ReadHelp)
	flag.StringVar(&reader.write, "write", "", WriteHelp)
	flag.StringVar(&reader.write, "w", "", WriteHelp)
	flag.StringVar(&reader.customer, "customer", "", "Customer Id")
	flag.StringVar(&reader.customer, "c", "", "Customer Id")
	flag.StringVar(&reader.parameters, "parameters", "", ParametersHelp)
	flag.StringVar(&reader.parameters, "p", "", ParametersHelp)
	flag.Parse()
}

//GetReadParameter gets read parameter. (Call Process first)
func (reader ParameterReader) GetReadParameter() string {
	return reader.read
}

//GetWriteParameter gets write parameter. (Call Process first)
func (reader ParameterReader) GetWriteParameter() string {
	return reader.write
}

//GetCustomerParameter customer read parameter. (Call Process first)
func (reader ParameterReader) GetCustomerParameter() string {
	return reader.customer
}

//GetWriteParameters gets write parameters. (Call Process first)
func (reader ParameterReader) GetWriteParameters() string {
	return reader.parameters
}

//NewParameterReader creates a new parameter reader
func NewParameterReader() ParameterReaderAPI {
	reader := ParameterReader{}
	return &reader
}

package service

//ParameterValidatorAPI defines api to validate input params
type ParameterValidatorAPI interface {
	Validate() string
}

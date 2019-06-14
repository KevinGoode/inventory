package service

//ParameterValidator is container for params validator
type ParameterValidator struct {
	reader ParameterReaderAPI
}

//Validate returns error message or empty string
func (validator *ParameterValidator) Validate() string {
	var returnString = ""

	customer := validator.reader.GetCustomerParameter()
	read := validator.reader.GetReadParameter()
	write := validator.reader.GetWriteParameter()
	parameters := validator.reader.GetWriteParameter()
	if customer == "" {
		returnString = "Must specify customer id on every command.\n"
	}
	if read != "" && write != "" {
		returnString = "Cannot read and write at the same time.\n"
	}
	if read != "" && read != "tags" && read != "storage_devices" && read != "assets" && read != "asset_hosts" {
		returnString = "Invalid read argument: " + read + "\n"
		returnString += ReadHelp + "\n"
	}
	if write != "" && write != "tag" && write != "storage_device" && write != "asset" && write != "asset_host" && parameters == "" {
		returnString = "Invalid write argument: " + write
		returnString += WriteHelp + "\n"
		returnString += ParametersHelp + "\n"
	}
	return returnString
}

//NewParameterValidator creates a new parameter reader
func NewParameterValidator(reader ParameterReaderAPI) ParameterValidatorAPI {
	validator := ParameterValidator{}
	validator.reader = reader
	return &validator
}

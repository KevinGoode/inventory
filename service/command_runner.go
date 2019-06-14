package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

//CommandRunner dummy storage strcut
type CommandRunner struct {
	paramReader ParameterReaderAPI
	validator   ParameterValidatorAPI
	dbFactory   DatabaseFactoryAPI
}

//Run runs command
func (runner CommandRunner) Run() (int, string) {
	parameters := runner.paramReader
	dbFactory := runner.dbFactory
	var exitCode = 0
	//defer dbFactory.DisposeAll()
	//Read and validate args
	parameters.Process()
	outputMessage := runner.validator.Validate()
	if outputMessage == "" {
		outputBuilder := strings.Builder{}
		customerDb := dbFactory.GetDatabase(parameters.GetCustomerParameter())
		if parameters.GetReadParameter() != "" {
			switch parameters.GetReadParameter() {
			case "tags":
				tags := customerDb.GetAllTags()
				for _, tag := range tags {
					Write(tag, &outputBuilder)
				}
			case "storage_devices":
				devices := customerDb.GetAllStorageDevices()
				for _, device := range devices {
					Write(device, &outputBuilder)
				}
			case "assets":
				assets := customerDb.GetAllAssets()
				for _, asset := range assets {
					Write(asset, &outputBuilder)
				}
			case "asset_hosts":
				hosts := customerDb.GetAllAssetHosts()
				for _, host := range hosts {
					Write(host, &outputBuilder)
				}
			}
			outputMessage = outputBuilder.String()
		}
		if parameters.GetWriteParameter() != "" {
			var err error
			var bytes []byte
			switch parameters.GetWriteParameter() {
			case "tag":
				var tag Tag
				bytes, err = json.Marshal(tag)
				err = json.Unmarshal([]byte(parameters.GetWriteParameters()), &tag)
				if err == nil {
					customerDb.CreateTag(tag)
				}
			case "asset_host":
				var host AssetHost
				bytes, err = json.Marshal(host)
				err = json.Unmarshal([]byte(parameters.GetWriteParameters()), &host)
				if err == nil {
					err = customerDb.CreateAssetHost(host)
				}
			case "storage_device":
				var device StorageDevice
				bytes, err = json.Marshal(device)
				err = json.Unmarshal([]byte(parameters.GetWriteParameters()), &device)
				if err == nil {
					err = customerDb.CreateStorageDevice(device)
				}
			case "asset":
				var asset Asset
				bytes, err = json.Marshal(asset)
				err = json.Unmarshal([]byte(parameters.GetWriteParameters()), &asset)
				if err == nil {
					err = customerDb.CreateAsset(asset)
				}
			}
			if err != nil {
				outputMessage = fmt.Sprintf("Couldn't set parameters. Error '%s'. Parameters should be : '%s'", err, string(bytes))
			}
		}
	} else {
		exitCode = -1
	}
	return exitCode, outputMessage
}

//NewCommandRunner creates new command runner
func NewCommandRunner(parameters ParameterReaderAPI, validator ParameterValidatorAPI, dbFactory DatabaseFactoryAPI) CommandRunnerAPI {
	runner := CommandRunner{}
	runner.validator = validator
	runner.paramReader = parameters
	runner.dbFactory = dbFactory
	return &runner
}

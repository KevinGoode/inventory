package service

import (
	"strings"
	"testing"

	"go.uber.org/dig"
)

var Reader *DummyParameterReader
var Runner CommandRunnerAPI
var Fac DatabaseFactoryAPI

func initDIAndSetAPIS() {
	container := dig.New()
	//Stub database
	container.Provide(NewDummyDbCreator)
	container.Provide(NewDummyReader)
	//Real classes
	container.Provide(NewDatabaseFactory)
	container.Provide(NewParameterValidator)
	container.Provide(NewCommandRunner)
	container.Invoke(func(reader ParameterReaderAPI, runner CommandRunnerAPI, fac DatabaseFactoryAPI) {
		var parameterReader = reader.(*DummyParameterReader)
		Reader = parameterReader
		Reader.Customer = "inventory_demo"
		Runner = runner
		Fac = fac
	})
}
func TestStorageDeviceRW(t *testing.T) {
	initDIAndSetAPIS()
	defer Fac.DisposeAll()
	//Perform simple device write
	Reader.Write = "storage_device"
	var device = StorageDevice{ID: "id1", Name: "StoreOnce1", Status: Status{Name: "online", Value: "ok", Criticality: "ok"},
		Tags: nil, Connections: nil, SerialNumber: "SN", LocalManagerIds: nil, AllStorageUnits: nil}
	outputBuilder := strings.Builder{}
	Write(device, &outputBuilder)
	Reader.Parameters = outputBuilder.String()
	err, message := Runner.Run()

	if err != 0 {
		t.Error("Expected write to run successfully")
	}
	if message != "" {
		t.Error("Expected message to be empty")
	}
	//Now perform read
	Reader.Write = ""
	Reader.Parameters = ""
	Reader.Read = "storage_devices"
	err, message = Runner.Run()
	if err != 0 {
		t.Error("Expected read to run successfully")
	}
	if message == "" {
		t.Error("Expected message to contain some data")
	}
}

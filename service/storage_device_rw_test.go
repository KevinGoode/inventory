package service

import (
	"testing"
)

func TestSingleDeviceSingleDbRW(t *testing.T) {
	//Perform simple device write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createDevice(db, "id1", "StoreOnce1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	devices := db.GetAllStorageDevices()
	if len(devices) != 1 {
		t.Error("Expected one device.")
	}
}
func TestMultipleDeviceSingleDbRW(t *testing.T) {
	//Perform simple device write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createDevice(db, "id1", "StoreOnce1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	err = createDevice(db, "id2", "3Par1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	devices := db.GetAllStorageDevices()
	if len(devices) != 2 {
		t.Error("Expected two devices.")
	}
}
func TestMultipleDeviceMulitpleDbRW(t *testing.T) {
	//Perform simple device write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createDevice(db, "id1", "StoreOnce1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	db2 := fac.GetDatabase("customer2")
	err = createDevice(db2, "id2", "3Par1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	devices := db.GetAllStorageDevices()
	if len(devices) != 1 {
		t.Error("Expected one device.")
	}
	devices = db2.GetAllStorageDevices()
	if len(devices) != 1 {
		t.Error("Expected one device.")
	}
}
func createDevice(db DatabaseAPI, id string, name string) error {
	var device = StorageDevice{ID: id, Name: name, Status: Status{Name: "online", Value: "ok", Criticality: "ok"},
		SerialNumber: "SN", AllStorageUnits: nil}
	return db.CreateStorageDevice(device)
}

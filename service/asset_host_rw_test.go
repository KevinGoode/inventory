package service

import (
	"testing"
)

func TestSingleAssetHostSingleDbRW(t *testing.T) {
	//Perform simple AssetHost write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createAssetHost(db, "id1", "win1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	assets := db.GetAllAssetHosts()
	if len(assets) != 1 {
		t.Error("Expected one asset host")
	}
}
func TestMultipleAssetHostSingleDbRW(t *testing.T) {
	//Perform simple AssetHost write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createAssetHost(db, "id1", "win1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	err = createAssetHost(db, "id2", "ubunutu1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	assets := db.GetAllAssetHosts()
	if len(assets) != 2 {
		t.Error("Expected two asset host.")
	}
}
func TestMultipleAssetHostMulitpleDbRW(t *testing.T) {
	//Perform simple AssetHost write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createAssetHost(db, "id1", "win1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	db2 := fac.GetDatabase("customer2")
	err = createAssetHost(db2, "id2", "ubuntu1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	assets := db.GetAllAssetHosts()
	if len(assets) != 1 {
		t.Error("Expected one asset host")
	}
	assets = db2.GetAllAssetHosts()
	if len(assets) != 1 {
		t.Error("Expected one asset host")
	}
}
func createAssetHost(db DatabaseAPI, id string, name string) error {
	var asset = AssetHost{ID: id, Name: name}
	return db.CreateAssetHost(asset)
}

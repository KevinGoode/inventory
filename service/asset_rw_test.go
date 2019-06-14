package service

import (
	"testing"
)

func TestSingleAssetSingleDbRW(t *testing.T) {
	//Perform simple Asset write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createAsset(db, "id1", "Oracledb1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	assets := db.GetAllAssets()
	if len(assets) != 1 {
		t.Error("Expected one asset")
	}
}
func TestMultipleAssetSingleDbRW(t *testing.T) {
	//Perform simple Asset write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createAsset(db, "id1", "Oracledb1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	err = createAsset(db, "id2", "SQL1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	assets := db.GetAllAssets()
	if len(assets) != 2 {
		t.Error("Expected two asset.")
	}
}
func TestMultipleAssetMulitpleDbRW(t *testing.T) {
	//Perform simple Asset write
	fac := NewDatabaseFactory(NewDummyDbCreator())
	defer fac.DisposeAll()
	db := fac.GetDatabase("customer1")
	err := createAsset(db, "id1", "Oracledb1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	db2 := fac.GetDatabase("customer2")
	err = createAsset(db2, "id2", "SQL1")
	if err != nil {
		t.Error("Expected write to run successfully")
	}
	//Now perform read
	assets := db.GetAllAssets()
	if len(assets) != 1 {
		t.Error("Expected one asset")
	}
	assets = db2.GetAllAssets()
	if len(assets) != 1 {
		t.Error("Expected one asset")
	}
}
func createAsset(db DatabaseAPI, id string, name string) error {
	var asset = Asset{ID: id, Name: name}
	return db.CreateAsset(asset)
}

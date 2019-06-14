package service

import (
	"reflect"
)

//DummyDb is a dummy in memory database
type DummyDb struct {
	StorageDevices []*StorageDevice
	Assets         []*Asset
	AssetHosts     []*AssetHost
	Tags           []*Tag
}

//Dispose does nothing
func (api DummyDb) Dispose() {

}

//CreateDatabase checks to see if database (aka keyspace) exists and if it dosn't then it is created
//A session is created and the database migrated to the most recent version then the session is stored.
//NOTE: Only one session is created per keyspace
func (api *DummyDb) CreateDatabase(CustomerID string) DatabaseAPI {
	db := DummyDb{}
	return &db
}

//GetAllStorageDevices gets all primary and secondary stroage servers
func (api DummyDb) GetAllStorageDevices() []*StorageDevice {
	return api.StorageDevices
}

// GetAllAssetHosts gets all asset hosts. IE servers (windows/linux etc) that host supported applications
func (api DummyDb) GetAllAssetHosts() []*AssetHost {
	return api.AssetHosts
}

//GetAllAssets gets all application assets eg databases
func (api DummyDb) GetAllAssets() []*Asset {
	return api.Assets
}

//GetAllTags gets all tags
func (api DummyDb) GetAllTags() []*Tag {
	return api.Tags
}

//CreateTag creates a tag
func (api *DummyDb) CreateTag(tag Tag) error {
	api.Tags = append(api.Tags, &tag)
	return nil
}

//CreateAssetHost creates an asset host
func (api *DummyDb) CreateAssetHost(host AssetHost) error {
	api.AssetHosts = append(api.AssetHosts, &host)
	return nil
}

//CreateStorageDevice creates a storage device
func (api *DummyDb) CreateStorageDevice(device StorageDevice) error {
	api.StorageDevices = append(api.StorageDevices, &device)
	return nil
}

//CreateAsset creates an asset
func (api *DummyDb) CreateAsset(asset Asset) error {
	api.Assets = append(api.Assets, &asset)
	return nil
}

//GetTableColumnNames gets column names
func (api DummyDb) GetTableColumnNames(table interface{}) []string {
	var columnNames []string
	t := reflect.TypeOf(table)
	for i := 0; i < t.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		columnName, ok := f.Tag.Lookup("cql")
		if ok == true {
			columnNames = append(columnNames, columnName)
		}
	}
	return columnNames
}

//NewDummyDbCreator creates a new dummy database
func NewDummyDbCreator() DatabaseCreaterAPI {
	db := DummyDb{}
	return &db
}

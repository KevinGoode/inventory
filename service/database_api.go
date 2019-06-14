package service

//DatabaseAPI is main database api to execute command
type DatabaseAPI interface {
	GetTableColumnNames(table interface{}) []string
	GetAllStorageDevices() []*StorageDevice
	GetAllAssetHosts() []*AssetHost
	GetAllAssets() []*Asset
	GetAllTags() []*Tag

	CreateTag(tag Tag) error
	CreateAssetHost(host AssetHost) error
	CreateAsset(asset Asset) error
	CreateStorageDevice(device StorageDevice) error
	Dispose()
}

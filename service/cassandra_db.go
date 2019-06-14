package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/migrate"
	"github.com/scylladb/gocqlx/qb"
)

//CassandraDb is a wrapper class for cassandra dn
type CassandraDb struct {
	session *gocql.Session
}

//Dispose closes session
func (api CassandraDb) Dispose() {
	api.session.Close()
}

//CreateDatabase checks to see if database (aka keyspace) exists and if it dosn't then it is created
//A session is created and the database migrated to the most recent version then the session is stored.
//NOTE: Only one session is created per keyspace
func (api *CassandraDb) CreateDatabase(CustomerID string) DatabaseAPI {
	if !api.doesKeySpaceExist(CustomerID) {
		if !api.createKeySpace(CustomerID) {
			fmt.Println("Failed to create keyspace")
			os.Exit(-1)
		}
	}
	api.createSession(CustomerID)
	api.upgrade()
	return api
}
func (api *CassandraDb) upgrade() {
	ctx := context.Background()
	if err := migrate.Migrate(ctx, api.session, "./versions"); err != nil {
		panic(err)
	}
}
func (api *CassandraDb) createSession(CustomerID string) {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = CustomerID
	cluster.Consistency = gocql.Quorum
	api.session, _ = cluster.CreateSession()
}
func (api *CassandraDb) createKeySpace(CustomerID string) bool {
	var rf = 1
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()
	var query = fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class' : 'SimpleStrategy','replication_factor' : %d}`, CustomerID, rf)
	err := session.Query(query).Exec()
	if err != nil {
		return false
	}
	return true
}
func (api *CassandraDb) doesKeySpaceExist(CustomerID string) bool {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()
	iter := session.Query("SELECT KEYSPACE_NAME FROM SYSTEM_SCHEMA.KEYSPACES").Iter()
	var keyspaceName string
	for iter.Scan(&keyspaceName) {
		if keyspaceName == CustomerID {
			return true
		}
	}
	return false
}

//GetAllStorageDevices gets all primary and secondary stroage servers
func (api CassandraDb) GetAllStorageDevices() []*StorageDevice {
	var devices []*StorageDevice
	stmt, _ := qb.Select("storage_devices").ToCql()
	err := gocqlx.Select(&devices, api.session.Query(stmt))
	if err != nil {
		log.Fatal(err)
	}
	return devices
}

//GetAllLogicalUnits gets all logical stroage units (Volumes, cloud stores, catlyst stores etc) for a given storage device
func (api CassandraDb) GetAllLogicalUnits(StorageDeviceID string) []*LogicalStorageUnit {
	var units []*LogicalStorageUnit
	//TODO
	return units
}

// GetAllAssetHosts gets all asset hosts. IE servers (windows/linux etc) that host supported applications
func (api CassandraDb) GetAllAssetHosts() []*AssetHost {
	var hosts []*AssetHost
	stmt, _ := qb.Select("asset_hosts").ToCql()
	err := gocqlx.Select(&hosts, api.session.Query(stmt))
	if err != nil {
		log.Fatal(err)
	}
	return hosts
}

//GetAllAssets gets all application assets eg databases
func (api CassandraDb) GetAllAssets() []*Asset {
	var assets []*Asset
	stmt, _ := qb.Select("assets").ToCql()
	err := gocqlx.Select(&assets, api.session.Query(stmt))
	if err != nil {
		log.Fatal(err)
	}
	return assets
}

//GetAllTags gets all tags
func (api CassandraDb) GetAllTags() []*Tag {
	var tags []*Tag
	stmt, _ := qb.Select("tags").ToCql()
	err := gocqlx.Select(&tags, api.session.Query(stmt))
	if err != nil {
		log.Fatal(err)
	}
	return tags
}

//CreateTag creates a tag
func (api CassandraDb) CreateTag(tag Tag) error {
	names := api.GetTableColumnNames(tag)
	stmt, names := qb.Insert("tags").Columns(names...).ToCql()
	q := gocqlx.Query(api.session.Query(stmt), names).BindStruct(tag)
	err := q.ExecRelease()
	return err
}

//CreateAssetHost creates an asset host
func (api CassandraDb) CreateAssetHost(host AssetHost) error {
	names := api.GetTableColumnNames(host)
	stmt, names := qb.Insert("asset_hosts").Columns(names...).ToCql()
	q := gocqlx.Query(api.session.Query(stmt), names).BindStruct(host)
	err := q.ExecRelease()
	return err
}

//CreateStorageDevice creates a storage device
func (api CassandraDb) CreateStorageDevice(device StorageDevice) error {
	names := api.GetTableColumnNames(device)
	stmt, names := qb.Insert("storage_devices").Columns(names...).ToCql()
	q := gocqlx.Query(api.session.Query(stmt), names).BindStruct(device)
	err := q.ExecRelease()
	return err
}

//CreateAsset creates an asset
func (api CassandraDb) CreateAsset(asset Asset) error {
	names := api.GetTableColumnNames(asset)
	stmt, names := qb.Insert("assets").Columns(names...).ToCql()
	q := gocqlx.Query(api.session.Query(stmt), names).BindStruct(asset)
	err := q.ExecRelease()
	return err
}

//GetTableColumnNames gets column names
func (api CassandraDb) GetTableColumnNames(table interface{}) []string {
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

//NewCassandraDbCreater creates a new cassandra database
func NewCassandraDbCreater() DatabaseCreaterAPI {
	db := CassandraDb{}
	return &db
}

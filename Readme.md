# Inventory

This is technology demonstrator of go apis: gocql and gocqlx  that wrap
the cassandra nosql  db. Application is a simple CLI to read and write records to a cassandra database on localhost.

The example uses a storage device inventory as an example db and demonstrates the data binding capabilities of gocqlx . 

Code automatically creates a keyspace(db) per customer so should work on localhost with cassandra installed. Inspect storage structs to see cql and json data bindings. The CLI accepts json bindings as parameter '-p'


# Example Commands
eg Write and write some records into inventory_demo db.\n
go run main.go -c inventory_demo -w storage_device -p {\"id\":\"1\"}\n
go run main.go -c inventory_demo -r storage_devices\n
go run main.go -c inventory_demo -w asset_host {\"id\":\"2\"}\n
go run main.go -c inventory_demo -r asset_hosts\n
go run main.go -c inventory_demo -w asset -p {\"id\":\"3\"}\n
go run main.go -c inventory_demo -r assets\n

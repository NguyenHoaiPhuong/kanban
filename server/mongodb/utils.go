package mongodb

// MongoWritingMethod : InsertMany or BulkWrite
type MongoWritingMethod string

const (
	// INSERTMANY : using InsertMany function
	INSERTMANY MongoWritingMethod = "InsertMany"
	// BULKWRITE : using BulkWrite function
	BULKWRITE MongoWritingMethod = "BulkWrite"
)

func generateMongoConnectionURI(serverHost, serverPort, username, password, dbName string) string {
	connectionURI := serverHost
	if username != "" && password != "" {
		connectionURI = username + ":" + password + "@" + serverHost
	}
	connectionURI = "mongodb://" + connectionURI + ":" + serverPort + "/" + dbName + "?authSource=admin"
	return connectionURI
}

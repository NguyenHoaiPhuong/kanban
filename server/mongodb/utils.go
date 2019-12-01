package mongodb

import "reflect"

// MongoWritingMethod : InsertMany or BulkWrite
type MongoWritingMethod string

const (
	// INSERTMANY : using InsertMany function
	INSERTMANY MongoWritingMethod = "InsertMany"
	// BULKWRITE : using BulkWrite function
	BULKWRITE MongoWritingMethod = "BulkWrite"
)

// GenerateMongoConnectionURI : generate connecting url with authentication
func GenerateMongoConnectionURI(serverHost, serverPort, username, password, dbName string) string {
	connectionURI := serverHost
	if username != "" && password != "" {
		connectionURI = username + ":" + password + "@" + serverHost
	}
	connectionURI = "mongodb://" + connectionURI + ":" + serverPort
	if dbName != "" {
		connectionURI = connectionURI + "/" + dbName + "?authSource=admin"
	}
	return connectionURI
}

// https://stackoverflow.com/questions/23555241/golang-reflection-how-to-get-zero-value-of-a-field-type
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	if v.CanInterface() {
		return v.Interface() == z.Interface()
	} else {
		return false
	}
	return true
}

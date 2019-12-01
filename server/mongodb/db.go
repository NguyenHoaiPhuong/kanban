package mongodb

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	a "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// DropDatabase : drop the specific database from server host
func DropDatabase(ctx context.Context, client *mongo.Client, dbName string) {
	db := client.Database(dbName)
	db.Drop(ctx)
}

// DropCollection : drop the specific collection from database
func DropCollection(ctx context.Context, db *mongo.Database, colName string) error {
	if CheckCollectionExist(ctx, db, colName) {
		err := db.Collection(colName).Drop(ctx)
		return err
	}
	return nil
}

// GetDBWithSubname : get all database names in the server host which contain subname
func GetDBWithSubname(ctx context.Context, client *mongo.Client, subName string) ([]string, error) {
	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if subName == "" {
		return dbNames, nil
	}

	dbNamesContainingSubname := make([]string, 0)
	for _, name := range dbNames {
		if strings.Contains(name, subName) {
			dbNamesContainingSubname = append(dbNamesContainingSubname, name)
		}
	}
	return dbNamesContainingSubname, nil
}

// CreateClient returns client respective to the given server host and port
// Refer to following link for more details of authentication
// https://docs.mongodb.com/manual/reference/connection-string/
func CreateClient(ctx context.Context, serverHost, serverPort, username, password, dbName string) (*mongo.Client, error) {
	connMsg := GenerateMongoConnectionURI(serverHost, serverPort, username, password, dbName)
	clientOptions := options.Client().ApplyURI(connMsg)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ConnectToDB returns db respective to the given server host, port and db name
func ConnectToDB(ctx context.Context, serverHost, serverPort, dbName string) (*mongo.Database, error) {
	client, err := CreateClient(ctx, serverHost, serverPort, "", "", "")
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

// CheckCollectionExist will check if the given collection name is present in the database names set
func CheckCollectionExist(ctx context.Context, db *mongo.Database, collectionName string) bool {
	names, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return false
	}
	for _, v := range names {
		if v == collectionName {
			return true
		}
	}
	return false
}

// CheckCollectionHasIndex will check if the given filed in given collection is indexed or not
func CheckCollectionHasIndex(ctx context.Context, db *mongo.Database, collectionName string, indexedFieldName string) bool {
	indexView := db.Collection(collectionName).Indexes()
	curIndex, err := indexView.List(ctx)
	defer curIndex.Close(ctx)
	if err != nil {
		return false
	}
	for curIndex.Next(ctx) {
		var result bson.D
		curIndex.Decode(&result)
		for _, field := range result {
			if field.Key == "key" {
				value := field.Value.(bson.D)
				for _, f := range value {
					if f.Key == indexedFieldName {
						return true
					}
				}
			}
		}

	}

	return false
}

// RemoveIndex : drop index
func RemoveIndex(ctx context.Context, db *mongo.Database, collectionName string, indexedFieldName string) error {
	indexView := db.Collection(collectionName).Indexes()
	_, err := indexView.DropOne(ctx, indexedFieldName)
	return err
}

// CheckIndexExistOrCreateIt will check for database index existing, output a message if not and then proceed to creating it
func CheckIndexExistOrCreateIt(ctx context.Context, db *mongo.Database, collectionName string,
	indexedFieldName string, msgIfNotExist a.Value) error {
	if CheckCollectionExist(ctx, db, collectionName) && CheckCollectionHasIndex(ctx, db, collectionName, indexedFieldName) == false {
		fmt.Println(msgIfNotExist)
		opts := options.CreateIndexes()
		keys := bsonx.Doc{{Key: indexedFieldName, Value: bsonx.Int32(int32(1))}}
		index := mongo.IndexModel{}
		index.Keys = keys
		_, err := db.Collection(collectionName).Indexes().CreateOne(ctx, index, opts)
		return err
	}
	return nil
}

// WriteToDB : save data onto mongodb
func WriteToDB(ctx context.Context, db *mongo.Database, colName string,
	object interface{}, batchLimit int, writingMethod MongoWritingMethod) error {
	var err error
	if CheckCollectionExist(ctx, db, colName) {
		err = DropCollection(ctx, db, colName)
		if err != nil {
			return err
		}
	}
	bulk := &BulkCollection{items: object, collectionName: colName}
	switch writingMethod {
	case INSERTMANY:
		err = bulk.InsertManyTo(db, batchLimit)
	case BULKWRITE:
		err = bulk.BulkWriteTo(db, batchLimit)
	}

	return err
}

// InsertToDB : Insert object into the given collection
func InsertToDB(ctx context.Context, db *mongo.Database, colName string,
	object interface{}, batchLimit int, writingMethod MongoWritingMethod) error {
	var err error
	bulk := &BulkCollection{items: object, collectionName: colName}
	switch writingMethod {
	case INSERTMANY:
		err = bulk.InsertManyTo(db, batchLimit)
	case BULKWRITE:
		err = bulk.BulkWriteTo(db, batchLimit)
	}
	return err
}

// CollectionIsSubsetOf : compare data in col1 to data in col2.
// If all data in col1 can be found in col2, the function returns true.
func CollectionIsSubsetOf(col1, col2 *mongo.Collection) bool {
	var expectedResults []bson.D
	ctx := context.Background()
	opts := options.Find()
	opts.SetSort(bson.D{{"Date", 1}})
	cursor, err := col1.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return false
	}
	cursor.All(ctx, &expectedResults)
	cursor.Close(ctx)
	for _, expectedResultEntry := range expectedResults {
		var finalFieldsQuery bson.D
		var expectedSubArrayHC bson.D
		for _, expectedResultEntryField := range expectedResultEntry {
			v := reflect.ValueOf(expectedResultEntryField.Value)
			t := reflect.TypeOf(expectedResultEntryField.Value)
			expectedResultsFieldType := v.Kind()
			if expectedResultEntryField.Value != nil && !isZero(v) {
				fmt.Println("expected result field :", expectedResultEntryField)
				if expectedResultEntryField.Key != "_id" {
					fmt.Println(expectedResultEntryField.Key, "--> Used")
					switch expectedResultsFieldType {
					case reflect.Func, reflect.Map:
						fmt.Println("skiping unsupported expected data field :", expectedResultEntryField.Key)
					case reflect.Array, reflect.Slice:
						isEmbeddedArray := false
						embeddedArrayLength := 0
						underlyingType := t.Elem()
						underlyingTypeName := underlyingType.Name()
						if underlyingTypeName == "E" {
							finalFieldsQuery = append(finalFieldsQuery, bson.E{
								expectedResultEntryField.Key, expectedResultEntryField.Value,
							})
						} else {
							isEmbeddedArray = true
							for i := 0; i < v.Len(); i++ {
								embeddedArrayLength++
								var expectedSubArray bson.D
								arrayEntryV := v.Index(i)
								if !isZero(arrayEntryV) {
									arrayEntryEl := arrayEntryV.Elem()

									kind := arrayEntryEl.Kind()
									if kind == reflect.Slice {
										arrayEntryEl := arrayEntryV.Elem()
										for j := 0; j < arrayEntryEl.Len(); j++ {
											elField := arrayEntryEl.Index(j).Interface()
											arrayField := elField.(bson.E)
											expectedSubArray = append(expectedSubArray, arrayField)
											expectedSubArrayHC = append(expectedSubArrayHC, arrayField)
										}
										finalFieldsQuery = append(finalFieldsQuery, bson.E{
											expectedResultEntryField.Key, bson.M{
												"$elemMatch": expectedSubArray,
											},
										})
									} else {
										fmt.Println("unsupported embedded array element, support only structure (not simple types)")
									}
								}
							}
						}

						if isEmbeddedArray {
							finalFieldsQuery = append(finalFieldsQuery, bson.E{
								expectedResultEntryField.Key, bson.M{
									"$size": embeddedArrayLength,
								},
							})
						}

					case reflect.Struct:
						fmt.Println("skiping unsupported expected data field : ", expectedResultEntryField.Key)
					default:
						finalFieldsQuery = append(finalFieldsQuery, expectedResultEntryField)
					}
				}
			}
		}
		if len(finalFieldsQuery) == 0 {
			continue
		}

		fmt.Println("Filter : ", finalFieldsQuery)
		var foundResults []bson.D

		cur, err := col2.Find(ctx, finalFieldsQuery)
		if err != nil {
			return false
		}
		cur.All(ctx, &foundResults)
		if len(foundResults) == 0 {
			fmt.Printf("Could not find expected results for collection %s in collection %s.\n", col1.Name(), col2.Name())
			fmt.Println("Details of the unfound expected result : ")
			fmt.Println(expectedResultEntry)
			return false
		}
		// fmt.Println(foundResults)
	}

	// If we arrive here, it means all test succeeded
	return true
}

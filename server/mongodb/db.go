package mongodb

package simcel

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	a "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// DropDatabase : drop the specific database from server host
func DropDatabase(ctx *context, client *mongo.Client, dbName string) {
	db := client.Database(dbName)
	db.Drop(ctx)
}

// DropCollection : drop the specific collection from database
func DropCollection(ctx *context, db *mongo.Database, colName string) error {
	if mongoCheckCollectionExist(db, colName) {
		err := db.Collection(colName).Drop(ctx)
		return err
	}
	return nil
}

// GetDBWithSubname : get all database names in the server host which contain subname
func GetDBWithSubname(ctx *context, client *mongo.Client, subName string) ([]string, error) {
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
func CreateClient(ctx *context, serverHost, serverPort, username, password, dbName string) (*mongo.Client, error) {
	connMsg := generateMongoConnectionURI(serverHost, serverPort, username, password, dbName)
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
func ConnectToDB(ctx *context, serverHost, serverPort, dbName string) (*mongo.Database, error) {
	client, err := CreateClient(ctx, serverHost, serverPort, "", "", "")
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

// CheckCollectionExist will check if the given collection name is present in the database names set
func CheckCollectionExist(ctx *context, db *mongo.Database, collectionName string) bool {
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
func CheckCollectionHasIndex(ctx *context, db *mongo.Database, collectionName string, indexedFieldName string) bool {
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
func RemoveIndex(ctx *context, db *mongo.Database, collectionName string, indexedFieldName string) error {
	indexView := db.Collection(collectionName).Indexes()
	_, err := indexView.DropOne(ctx, indexedFieldName)
	return err
}

// CheckIndexExistOrCreateIt will check for database index existing, output a message if not and then proceed to creating it
func CheckIndexExistOrCreateIt(ctx *context, db *mongo.Database, collectionName string, 
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
func WriteToDB(ctx *context, db *mongo.Database, colName string, object interface{}, writingMethod MongoWritingMethod) error {
	if CheckCollectionExist(ctx, db, colName) {
		err := DropCollection(ctx, db, colName)
		if err != nil {
			return err
		}
	}
	bulk := &BulkCollection{items: object, collectionName: colName}
	switch writingMethod {
	case INSERTMANY:
		bulk.mongoInsertManyTo(db)
	case BULKWRITE:
		bulk.mongoBulkWriteTo(db)
	}
}

func mongoInsertToDB(db *mongo.Database, colName string, object interface{}, writingMethod MongoWritingMethod) {
	bulk := &BulkCollection{items: object, collectionName: colName}
	switch writingMethod {
	case INSERTMANY:
		bulk.mongoInsertManyTo(db)
	case BULKWRITE:
		bulk.mongoBulkWriteTo(db)
	}
}

func mongoInitDatabase(duplicatedDb bool, configDbName string, client *mongo.Client) (usedDb string) {
	fmt.Println("Using network configuration database : ", configDbName)
	biggestRunNbr := 0
	dbExist := false

	options := options.Find()
	options.SetSort(bson.D{{"DB_ID", -1}})
	options.SetLimit(1)
	result := bson.M{}
	ctx := context.Background()
	runResults := client.Database("simulation_config").Collection("run_result_dbs")
	cursor, err := runResults.Find(ctx, bson.M{"configDbName": configDbName}, options)
	CheckErr(err)
	defer cursor.Close(ctx)
	cursor.Next(ctx)
	err = cursor.Decode(&result)
	if err != nil {
		fmt.Println("could not retrieve latest simulation result database, will create default one")
	} else {
		if casted, ok := result["DB_ID"]; ok {
			dbExist = true
			num := casted.(int32)
			biggestRunNbr = int(num)
			fmt.Println("Latest run result database ID : ", biggestRunNbr)
		} else {
			fmt.Println(casted, ok)
			fmt.Println(result)
			log.Fatal(ok)
		}
	}

	if duplicatedDb || !dbExist {
		biggestRunNbr++
		runResults.InsertOne(context.Background(), bson.M{"configDbName": configDbName, "DB_ID": biggestRunNbr})
	}

	usedDbName := configDbName + RESULT_DB_SUFFIX + strconv.Itoa(biggestRunNbr)
	// We should create a new DB either if there is no DB or if asked by params
	eraseCurrentDb := dbExist && !duplicatedDb

	if eraseCurrentDb {
		fmt.Println("Delete database", usedDbName)
		mongoDropDatabase(client, usedDbName)
	}

	return usedDbName
}

// installDatabse should be run once before all simulation. It will install the initial network database, along with validation (and import customer data )
func mongoInstallDatabase(cfg *SimcelConfig) {
	ctx := context.Background()

	host := *cfg.mongodbServerHost
	port := *cfg.mongodbServerPort
	dbName := "simulation_config"
	db := mongoConnectToDB(host, port, dbName)

	colName := "runs"
	if !mongoCheckCollectionExist(db, colName) {
		res := db.RunCommand(ctx, bson.M{
			"create": "runs",
		}, nil)

		CheckErr(res.Err())
	}

	dbName = "simulation_runs"
	db = mongoConnectToDB(host, port, dbName)
	if !mongoCheckCollectionExist(db, colName) {
		res := db.RunCommand(ctx, bson.M{
			"create": "runs",
		}, nil)

		CheckErr(res.Err())
	}
}

func (scDesc *SCModelDescriptor) mongoSaveToDatabase(db *mongo.Database, collectionName string, SCModelsMaps map[string]map[ISCModel]ISCModel, writingMethod MongoWritingMethod) {
	typeName := reflect.TypeOf(scDesc.mod).Elem().Name()

	total := len(SCModelsMaps[typeName])

	toInsert := make([]interface{}, total)
	i := 0
	for _, scModel := range SCModelsMaps[typeName] {
		toInsert[i] = scModel
		i++
	}
	bulk := &BulkCollection{items: toInsert, collectionName: collectionName}
	switch writingMethod {
	case INSERTMANY:
		bulk.mongoInsertManyTo(db)
	case BULKWRITE:
		bulk.mongoBulkWriteTo(db)
	}
}

func mongoCheckAllTransactionnalCollectionHaveIndexes(client *mongo.Client, dbName string) {
	msg := a.BrightYellow("|||||||||||| InventoryAdjustment database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	err := mongoCheckIndexExistOrCreateIt(client.Database(dbName), string(InventoryAdjustmentCol), "Date", msg)
	CheckErr(err)
	msg = a.BrightYellow("|||||||||||| CustomerDemand database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	err = mongoCheckIndexExistOrCreateIt(client.Database(dbName), string(CustomerDemandCol), "Date", msg)
	CheckErr(err)
	msg = a.BrightYellow("|||||||||||| ReplayDeliveryOrders database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	err = mongoCheckIndexExistOrCreateIt(client.Database(dbName), string(ReplayDeliveryOrderCol), "Date", msg)
	CheckErr(err)
	msg = a.BrightYellow("|||||||||||| ReplayReplenishSOs database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	err = mongoCheckIndexExistOrCreateIt(client.Database(dbName), string(ReplayReplenishSOCol), "Date", msg)
	CheckErr(err)
	msg = a.BrightYellow("|||||||||||| ReplayProduction database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	err = mongoCheckIndexExistOrCreateIt(client.Database(dbName), string(ReplayProductionCol), "Date", msg)
	CheckErr(err)
}

// mongoCollectionIsSubsetOf : compare data in col1 to data in col2.
// If all data in col1 can be found in col2, the function returns true.
func mongoCollectionIsSubsetOf(col1, col2 *mongo.Collection) bool {
	var expectedResults []bson.D
	ctx := context.Background()
	opts := options.Find()
	opts.SetSort(bson_mongo.D{{"Date", 1}})
	cursor, err := col1.Find(ctx, bson.D{{}}, opts)
	CheckErr(err)
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
				printIfVerbose("expected result field :", expectedResultEntryField)
				if expectedResultEntryField.Key != "_id" &&
					expectedResultEntryField.Key != "ID" &&
					expectedResultEntryField.Key != "appliedsegments" &&
					expectedResultEntryField.Key != "kpis" &&
					expectedResultEntryField.Key != "inputmodel" &&
					expectedResultEntryField.Key != "BatchID" {
					printIfVerbose(expectedResultEntryField.Key, "--> Used")
					switch expectedResultsFieldType {
					case reflect.Func, reflect.Map:
						printIfVerbose("skiping unsupported expected data field : ", expectedResultEntryField.Key)
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
										panic("unsupported embedded array element, support only structure (not simple types)")
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
						printIfVerbose("skiping unsupported expected data field : ", expectedResultEntryField.Key)
					default:
						finalFieldsQuery = append(finalFieldsQuery, expectedResultEntryField)
					}
				}
			}
		}
		if len(finalFieldsQuery) == 0 {
			continue
		}

		printIfVerbose("Filter : ", finalFieldsQuery)
		var foundResults []bson.D

		cur, err := col2.Find(ctx, finalFieldsQuery)
		CheckErr(err)
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

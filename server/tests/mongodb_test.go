package test

import (
	"context"
	"log"
	"testing"

	"github.com/NguyenHoaiPhuong/kanban/server/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoTestDB(dbName string) *mongo.Database {
	serverHost := "localhost"
	serverPort := "27017"
	ctx := context.Background()
	db, err := mongodb.ConnectToDB(ctx, serverHost, serverPort, dbName)
	if err != nil {
		log.Fatalln("Cannot connect to database")
	}

	return db
}

func TestDropCollection(t *testing.T) {
	dbName := "random_test_1"
	db := initMongoTestDB(dbName)

	colName := "FacilityMaster"

	ctx := context.Background()
	mongodb.DropCollection(ctx, db, colName)
	if mongodb.CheckCollectionExist(ctx, db, colName) {
		t.Errorf("Error: Collection named %s in database %s WASN'T deleted", colName, dbName)
	}
}

func TestDropDatabase(t *testing.T) {
	serverHost := "localhost"
	serverPort := "27017"
	client := mongoCreateClient(serverHost, serverPort)

	subName := "test_2"
	dbNames := mongoGetDBWithSubname(client, subName)
	for _, dbName := range dbNames {
		mongoDropDatabase(client, dbName)
	}
	dbNames = mongoGetDBWithSubname(client, subName)
	if len(dbNames) > 0 {
		for _, dbName := range dbNames {
			t.Errorf("Error: database %s WASN'T deleted", dbName)
		}
	}
}

func TestCheckCollectionExist(t *testing.T) {
	dbName := "random_test_1"
	db := initMongoTestDB(dbName)
	colName := "CustomerDemand"
	if !mongoCheckCollectionExist(db, colName) {
		t.Errorf("Error: Collection %s exists in the db %s\n", colName, dbName)
	}

	colName = "CustomerDemand11"
	if mongoCheckCollectionExist(db, colName) {
		t.Errorf("Error: Collection %s DOES NOT exist in the db %s\n", colName, dbName)
	}
}

func TestCheckIndexExistOrCreateIt(t *testing.T) {
	dbName := "random_test_1"
	db := initMongoTestDB(dbName)
	colName := "CustomerDemand"
	fieldName := "Date"
	mongoRemoveIndex(db, colName, fieldName)
	msg := a.BrightYellow("|||||||||||| CustomerDemand database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	if err := mongoCheckIndexExistOrCreateIt(db, colName, fieldName, msg); err != nil {
		t.Errorf("Error: Cannot create index for field %s in collection %s in the db %s\n", fieldName, colName, dbName)
	}
}

func TestCheckCollectionHasIndex(t *testing.T) {
	dbName := "random_test_1"
	db := initMongoTestDB(dbName)

	colName := "CustomerDemand"
	fieldName := "Date"
	if !mongoCheckCollectionHasIndex(db, colName, fieldName) {
		t.Errorf("Error: Collection %s in the db %s has indexed field %s\n", colName, dbName, fieldName)
	}

	fieldName = "CustomerRef"
	if mongoCheckCollectionHasIndex(db, colName, fieldName) {
		t.Errorf("Error: Collection %s in the db %s has NO indexed field %s\n", colName, dbName, fieldName)
	}
}

func TestReadWriteDatabase(t *testing.T) {
	dbName := "random_test_1"
	db := initMongoTestDB(dbName)
	colName := "CustomerDemand"
	col := db.Collection(colName)

	/********** Test reading data **********/
	ctx := context.Background()
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetSort(bson.D{{"Date", 1}})
	demands := make(Demands, 0)
	cursor, err := col.Find(ctx, filter, opts)
	CheckErr(err)
	cursor.All(ctx, &demands)

	currentTime := demands[0].Date.UTC()
	for _, demand := range demands {
		// fmt.Println(demand)
		if currentTime.After(demand.Date.UTC()) {
			t.Errorf("Error: sorting customer demand based on Date WRONGLY. Date %v is after date %v.", currentTime, demand.Date.UTC())
		}
		currentTime = demand.Date.UTC()
	}

	/********** Test writing data **********/
	newDBName := dbName + "_save"
	db = initMongoTestDB(newDBName)
	newColName := colName + "_BULKWRITE"
	mongoWriteToDB(db, newColName, demands, BULKWRITE)
	col1 := db.Collection(newColName)
	if !mongoCollectionIsSubsetOf(col, col1) || !mongoCollectionIsSubsetOf(col1, col) {
		t.Errorf("mongo-driver BULKWRITE error.")
	}
	newColName = colName + "_INSERTMANY"
	mongoWriteToDB(db, newColName, demands, INSERTMANY)
	col1 = db.Collection(newColName)
	if !mongoCollectionIsSubsetOf(col, col1) || !mongoCollectionIsSubsetOf(col1, col) {
		t.Errorf("mongo-driver INSERTMANY error.")
	}
}

package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/NguyenHoaiPhuong/kanban/server/mongodb"
	a "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Demand struct
type Demand struct {
	Date        time.Time `json:"Date" bson:"Date" required:"true"`
	CustomerRef string    `json:"CustomerRef" bson:"CustomerRef" required:"true"`
	ProductID   string    `json:"ProductID" bson:"ProductID" required:"true"`
	Quantity    float64   `json:"Quantity" bson:"Quantity" required:"true"`
	Unit        string    `json:"Unit" bson:"Unit" required:"true"`
}

// Demands : slice of demands
type Demands []*Demand

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
	ctx := context.Background()

	serverHost := "localhost"
	serverPort := "27017"
	client, err := mongodb.CreateClient(ctx, serverHost, serverPort, "", "", "")
	if err != nil {
		t.Error("CreateClient Error:", err)
	}

	subName := "test_2"
	dbNames, err := mongodb.GetDBWithSubname(ctx, client, subName)
	if err != nil {
		t.Error("GetDBWithSubname Error:", err)
	}
	for _, dbName := range dbNames {
		mongodb.DropDatabase(ctx, client, dbName)
	}
	dbNames, err = mongodb.GetDBWithSubname(ctx, client, subName)
	if err != nil {
		t.Error("GetDBWithSubname Error:", err)
	}
	if len(dbNames) > 0 {
		for _, dbName := range dbNames {
			t.Errorf("Error: database %s WASN'T deleted", dbName)
		}
	}
}

func TestCheckCollectionExist(t *testing.T) {
	ctx := context.Background()

	dbName := "random_test_1"
	db := initMongoTestDB(dbName)
	colName := "CustomerDemand"
	if !mongodb.CheckCollectionExist(ctx, db, colName) {
		t.Errorf("Error: Collection %s exists in the db %s\n", colName, dbName)
	}

	colName = "CustomerDemand11"
	if mongodb.CheckCollectionExist(ctx, db, colName) {
		t.Errorf("Error: Collection %s DOES NOT exist in the db %s\n", colName, dbName)
	}
}

func TestCheckIndexExistOrCreateIt(t *testing.T) {
	ctx := context.Background()

	dbName := "random_test_1"
	db := initMongoTestDB(dbName)
	colName := "CustomerDemand"
	fieldName := "Date"
	mongodb.RemoveIndex(ctx, db, colName, fieldName)
	msg := a.BrightYellow("|||||||||||| CustomerDemand database has no index on the date, creating index, this may take several minutes  ||||||||||||").Bold().BgBrightRed()
	if err := mongodb.CheckIndexExistOrCreateIt(ctx, db, colName, fieldName, msg); err != nil {
		t.Errorf("Error: Cannot create index for field %s in collection %s in the db %s\n", fieldName, colName, dbName)
	}
}

func TestCheckCollectionHasIndex(t *testing.T) {
	ctx := context.Background()

	dbName := "random_test_1"
	db := initMongoTestDB(dbName)

	colName := "CustomerDemand"
	fieldName := "Date"
	if !mongodb.CheckCollectionHasIndex(ctx, db, colName, fieldName) {
		t.Errorf("Error: Collection %s in the db %s has indexed field %s\n", colName, dbName, fieldName)
	}

	fieldName = "CustomerRef"
	if mongodb.CheckCollectionHasIndex(ctx, db, colName, fieldName) {
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
	if err != nil {
		t.Error(err)
	}
	cursor.All(ctx, &demands)

	currentTime := demands[0].Date.UTC()
	for _, demand := range demands {
		// log.Println(demand)
		if currentTime.After(demand.Date.UTC()) {
			t.Errorf("Error: sorting customer demand based on Date WRONGLY. Date %v is after date %v.", currentTime, demand.Date.UTC())
		}
		currentTime = demand.Date.UTC()
	}

	/********** Test writing data **********/
	newDBName := dbName + "_save"
	db = initMongoTestDB(newDBName)
	newColName := colName + "_BULKWRITE"
	mongodb.WriteToDB(ctx, db, newColName, demands, 100, mongodb.BULKWRITE)
	col1 := db.Collection(newColName)
	if !mongodb.CollectionIsSubsetOf(col, col1) || !mongodb.CollectionIsSubsetOf(col1, col) {
		t.Errorf("mongo-driver BULKWRITE error.")
	}
	newColName = colName + "_INSERTMANY"
	mongodb.WriteToDB(ctx, db, newColName, demands, 100, mongodb.INSERTMANY)
	col1 = db.Collection(newColName)
	if !mongodb.CollectionIsSubsetOf(col, col1) || !mongodb.CollectionIsSubsetOf(col1, col) {
		t.Errorf("mongo-driver INSERTMANY error.")
	}
}

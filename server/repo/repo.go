package repo

import (
	"context"
	"log"
	"reflect"

	"github.com/NguyenHoaiPhuong/kanban/server/models"
	"github.com/NguyenHoaiPhuong/kanban/server/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB struct
type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// Init : initialize MongoDB
func (mdb *MongoDB) Init(host, port, dbName string) {
	ctx := context.Background()
	client, err := mongodb.CreateClient(ctx, host, port, "", "", "")
	if err != nil {
		log.Fatalln("Init MongoDB Error:", err)
	}
	mdb.Client = client
	mdb.DB = client.Database(dbName)
}

// GetAllDocuments : get all documents in the given DB and Collection
func (mdb *MongoDB) GetAllDocuments(ctx context.Context, colName string, inMod models.IModel) (models.IModels, error) {
	// Using reflection to create a slice of the required type
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(inMod)), 0, 0)
	// Using reflection to create a pointer to this slice (required arguement for mgo.All())
	slicePtr := reflect.New(slice.Type())
	slicePtr.Elem().Set(slice)

	opts := options.Find()
	cursor, err := mdb.DB.Collection(colName).Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, slicePtr.Interface())
	cursor.Close(ctx)

	finalSlice := slicePtr.Elem()
	mods := make(models.IModels, finalSlice.Len())
	for i := 0; i < finalSlice.Len(); i++ {
		elem := finalSlice.Index(i).Interface()
		mods[i] = elem.(models.IModel)
	}

	return mods, nil
}

// GetDocumentByKey gets document by given key and respective value
func (mdb *MongoDB) GetDocumentByKey(ctx context.Context, colName string, key string, value interface{}) (models.IModel, error) {
	var doc models.IModel

	opts := options.FindOne()
	err := mdb.DB.Collection(colName).FindOne(ctx, bson.M{key: value}, opts).Decode(&doc)

	return doc, err
}

// AddDocument adds new document
func (mdb *MongoDB) AddDocument(ctx context.Context, colName string, doc interface{}) error {
	opts := options.InsertOne()
	col := mdb.DB.Collection(colName)
	_, err := col.InsertOne(ctx, doc, opts)
	return err
}

// UpdateDocument update specific document
func (mdb *MongoDB) UpdateDocument(ctx context.Context, colName string, mod models.IModel) error {
	col := mdb.DB.Collection(colName)
	_, err := col.UpdateOne(ctx, bson.M{"ID": mod.GetID()}, mod)
	return err
}

// DeleteDocumentByKey deletes document by given key and respective value
func (mdb *MongoDB) DeleteDocumentByKey(ctx context.Context, colName string, key string, value interface{}) error {
	col := mdb.DB.Collection(colName)
	_, err := col.DeleteOne(ctx, bson.M{key: value})
	return err
}

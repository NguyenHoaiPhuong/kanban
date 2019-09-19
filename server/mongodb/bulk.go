package mongodb

import (
	"context"
	"log"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// A BulkCollection wrap an entity collection to support doing bulk insert in an elegant way
type BulkCollection struct {
	items          interface{}
	collectionName string
}

// InsertManyTo : Insert items by writing method InsertMany
func (b *BulkCollection) InsertManyTo(db *mongo.Database, batchLimit int) error {
	list := reflect.ValueOf(b.items)

	batchStart := 0
	length := list.Len()
	var toInsert []interface{}
	if batchLimit < length {
		toInsert = make([]interface{}, batchLimit)
	} else {
		toInsert = make([]interface{}, length)
	}

	if length == 0 {
		log.Println("Could not write collection", b.collectionName, "with zero item")
		return nil
	}

	if b.collectionName == "" {
		b.collectionName = reflect.TypeOf(list.Index(0).Interface()).Elem().Name()
	}

	for i := 0; i < length; i++ {
		toInsert[batchStart] = list.Index(i).Interface()

		if (batchStart < batchLimit-1) && (i < length-1) {
			// If batch counter do not reach limit and data length
			batchStart++
			continue
		}

		uc := db.Collection(b.collectionName)
		res, err := uc.InsertMany(context.Background(), toInsert)
		if err != nil {
			return err
		}

		checkErrWritingData(batchStart+1, len(res.InsertedIDs), b.collectionName)

		batchStart = 0
		if i+batchLimit+1 > length {
			toInsert = make([]interface{}, length-i-1)
		} else {
			toInsert = make([]interface{}, batchLimit)
		}
	}

	return nil
}

// BulkWriteTo : Save items by writing method BulkWrite
func (b *BulkCollection) BulkWriteTo(db *mongo.Database, batchLimit int) error {
	list := reflect.ValueOf(b.items)

	batchStart := 0
	length := list.Len()
	var toInsert []mongo.WriteModel
	if batchLimit < length {
		toInsert = make([]mongo.WriteModel, batchLimit)
	} else {
		toInsert = make([]mongo.WriteModel, length)
	}
	opts := options.BulkWrite()
	opts.SetOrdered(false)

	if length == 0 {
		log.Println("Could not write collection ", b.collectionName, " with zero item")
		return nil
	}

	if b.collectionName == "" {
		b.collectionName = reflect.TypeOf(list.Index(0).Interface()).Elem().Name()
	}

	for i := 0; i < length; i++ {
		doc := list.Index(i).Interface()
		iom := mongo.NewInsertOneModel()
		iom = iom.SetDocument(doc)
		toInsert[batchStart] = iom

		if (batchStart < batchLimit-1) && (i < length-1) {
			// If batch counter do not reach limit and data length
			batchStart++
			continue
		}

		uc := db.Collection(b.collectionName)
		res, err := uc.BulkWrite(context.Background(), toInsert, opts)
		if err != nil {
			return err
		}

		checkErrWritingData(batchStart+1, int(res.InsertedCount), b.collectionName)

		batchStart = 0
		if i+batchLimit+1 > length {
			toInsert = make([]mongo.WriteModel, length-i-1)
		} else {
			toInsert = make([]mongo.WriteModel, batchLimit)
		}
	}

	return nil
}

// checkErrWritingData : Check whether losing data while writing onto MongoDB
func checkErrWritingData(sentCount int, writtenCount int, colName string) {
	if sentCount != writtenCount {
		msg := "Number of " + string(colName) + " items written onto MongoDB: " + strconv.Itoa(writtenCount) +
			". Number of " + string(colName) + " items expected to write onto MongoDB:" + strconv.Itoa(sentCount)
		log.Println(msg)
	}
}

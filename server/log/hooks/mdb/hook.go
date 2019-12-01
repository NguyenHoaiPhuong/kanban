package mdb

import (
	"context"
	"log"

	"github.com/NguyenHoaiPhuong/warehouse/server/log/hooks"
	"github.com/NguyenHoaiPhuong/warehouse/server/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Hook is responsible for writing all logging messages onto collection named LogMessages in output DB
type Hook struct {
	cfg        *MongoConfig
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	hooks.Hook
}

// Init initializes mongodb hook by the given MongoDBConfig
func (hook *Hook) Init(cfg *MongoConfig, level int) {
	hook.cfg = cfg.Copy()
	hook.initDatabase()

	// IMPORTANT: below commented codes demonstrate that it is possible to assign specific logging levels for a hook
	// For example: assigning (PanicLevel, ErrorLevel, DebugLevel) --> the hook will handle only logging messages respective to those lelvels
	// hook.LogLevels = []logrus.Level{
	// 	logrus.PanicLevel,
	// 	logrus.FatalLevel,
	// 	logrus.ErrorLevel,
	// 	logrus.WarnLevel,
	// 	logrus.InfoLevel,
	// 	logrus.DebugLevel,
	// 	logrus.TraceLevel,
	// }
	hook.SetLevel(level)
}

func (hook *Hook) initDatabase() {
	ctx := context.Background()
	client, err := mongodb.CreateClient(ctx, *hook.cfg.Host, *hook.cfg.Port, *hook.cfg.UserName, *hook.cfg.Password, *hook.cfg.DBName)
	if err != nil {
		log.Fatalln("Init MongoDB Error:", err)
	}
	hook.client = client
	hook.database = client.Database(*hook.cfg.DBName)
	// IMPORTANT : hard code the collection name of logging message here
	colName := "LogMessages"
	hook.collection = hook.database.Collection(colName)
}

// Fire implements
func (hook *Hook) Fire(e *logrus.Entry) error {
	msg := make(map[string]interface{})
	msg["TimeStamp"] = e.Time
	msg["LoggingLevel"] = e.Level
	msg["Message"] = e.Message

	_, err := hook.collection.InsertOne(context.Background(), msg)

	return err
}

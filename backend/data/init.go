// data/init.go
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	context "context"
	fmt "fmt"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	readpref "go.mongodb.org/mongo-driver/mongo/readpref"
	log "log"
	time "time"
)

// Database container
type Database struct {
	// Database logger
	logger *log.Logger

	// Mongo Client
	client *mongo.Client
}

// Create a new database client object
func CreateDatabaseClient(l *log.Logger) *Database {
	return &Database{logger: l, client: nil}
}

// Connect to the database
func (db *Database) Connect(username, password, host, collection string) error {
	var err error

	log.Println("Not using username and password for authentication.")

	// Create a new database client object
	db.client, err = mongo.NewClient(options.Client().ApplyURI(
		//fmt.Sprintf("mongodb://%s:%s@%s:27017/%s?retryWrites=true&w=majority",
		//	username, password, host, collection,
		//),
		fmt.Sprintf("mongodb://%s:27017/%s?retryWrites=true&w=majority",
			host, collection,
		),
	))

	if err != nil {
		db.logger.Println(err)
		return err
	}

	// Create a new context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.client.Connect(ctx)
	if err != nil {
		db.logger.Println(err)
		return err
	}

	// Attempt to ping the connection
	if err = db.client.Ping(ctx, readpref.Primary()); err != nil {
		db.logger.Println(err)
		return err
	}

	return nil
}

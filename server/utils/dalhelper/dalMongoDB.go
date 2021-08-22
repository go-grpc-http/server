package dalhelper

import (
	"context"
	"gin-solid-template/util/confighelper"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *mongo.Client
var sessionError error
var once sync.Once

//GetMongoConnection creates Mongo Client.
//Returns client if connection established else error
func GetMongoConnection() (*mongo.Client, error) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//init client options with uri
		clientOptions := options.Client().ApplyURI(confighelper.GetConfig("MongoDSN"))

		//set maximum number of idle connections to handle
		clientOptions.SetMaxConnIdleTime(100)

		//set maximum number of open connections to handle
		clientOptions.SetMaxPoolSize(1000)

		/* clientOptions.SetRetryReads(true)
		clientOptions.SetRetryWrites(true) */

		//max connection idle time
		clientOptions.SetMaxConnIdleTime(4 * time.Hour)

		//Initiate connection
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			sessionError = err
			return
		}

		//Ping to check connection status
		err = client.Ping(ctx, nil)
		if err != nil {
			sessionError = err
			return
		}
		instance = client
	})
	return instance, sessionError
}

//GetMongoConnectionTest connects to the mongo server with provided DSN
//Mainly used when establishing connections for testing
func GetMongoConnectionTest(mongoDSN string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDSN))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

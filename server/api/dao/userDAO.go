package dao

import (
	"context"
	"fmt"
	models "go-gin-boilerplate/server/api/models"
	"go-gin-boilerplate/server/utils/confighelper"
	"go-gin-boilerplate/server/utils/dalhelper"
	"go-gin-boilerplate/server/utils/logginghelper"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// SaveTodoToDB - save user to MongoDB
func SaveUserToDB(ctx context.Context, user models.User) error {
	logginghelper.LogDebug("IN : SaveUserToDB", user)
	uuidVar, _ := uuid.NewV4() // uuid
	user.Id = uuidVar.String()
	d, err := getMongoDBSession()
	if err != nil {
		logginghelper.LogError("ERROR : in MongoDB collection!  ---> ", err)
		return err
	}
	_, err = d.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserFromDB(userRequest models.Login) (models.User, bool, error) {
	fmt.Printf("%+v", userRequest)
	var user models.User
	d, err := getMongoDBSession()
	if err != nil {
		logginghelper.LogError("ERROR : in MongoDB collection!  ---> ", err)
		return user, false, err
	}
	filter := bson.M{"userName": userRequest.UserName}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = d.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func getMongoDBSession() (*mongo.Collection, error) {
	var collection *mongo.Collection
	dbName := confighelper.GetConfig("MongoDBName")
	collectionName := confighelper.GetConfig("MongoDBCollectionName")

	session, err := dalhelper.GetMongoConnection()
	if err != nil {
		logginghelper.LogError("ERROR : Error occurred while SaveUserToDB  ---> ", err)
		logginghelper.LogDebug("OUT : SaveUserToDB")
		return collection, err
	}
	d := session.Database(dbName).Collection(collectionName)
	return d, nil
}

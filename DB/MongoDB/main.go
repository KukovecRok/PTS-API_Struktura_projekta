package MongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"todorok/DataStructures"
)

type MongoDB struct {
	Client        *mongo.Client
	User          string
	Pass          string
	IP            string
	Port          int
	Database      string
	AuthDB        string
	AuthMechanism string
}


const (
	// Timeout operations after N seconds.
	connectionStringTemplate = "mongodb://%s:%s@%s:%d/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=%s&authMechanism=%s"
)

func (dbo *MongoDB) Init(ctx context.Context) (err error) {
	connectionURI := fmt.Sprintf(connectionStringTemplate, dbo.User, dbo.Pass, dbo.IP, dbo.Port, dbo.AuthDB, dbo.AuthMechanism)

	dbo.Client, err = mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		return
	}
	err = dbo.Client.Connect(ctx)
	if err != nil {

		return
	}
	err = dbo.DoInit(ctx)

	return
}

func (dbo *MongoDB) DoInit(ctx context.Context) (err error) {
	// Če ni vpisanih vlog in uporabnikov, dodaj default admin
	if count, err := dbo.Client.Database(dbo.Database).Collection("roles").CountDocuments(ctx, bson.M{}); count == 0 {
		if err != nil {
			return err
		}
		if count, err := dbo.Client.Database(dbo.Database).Collection("users").CountDocuments(ctx, bson.M{}); count == 0 {

			if err != nil {
				return err
			}
			_, err = dbo.Client.Database(dbo.Database).Collection("users").InsertOne(ctx, DataStructures.User{
				Password: "$2y$12$0AkUUBh4626wsJSd5u.FQ.KwCrLxqfDQDDYVcwpCtBWCgj9.GtAFC", // demo
				Email:    "admin@it-tim.si",
			})
			if err != nil {
				return err
			}
		}

	}

	// Kreiranje indeksov je idempotentno
	_, err = dbo.Client.Database(dbo.Database).Collection("users").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	// Kreiranje indeksov je idempotentno
	_, err = dbo.Client.Database(dbo.Database).Collection("steklenicas").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		return err
	}

	err = dbo.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

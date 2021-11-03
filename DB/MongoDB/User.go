package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (dbo *MongoDB) GetUsers(ctx context.Context) (users []DataStructures.User, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	users = make([]DataStructures.User, 0)

	err = cursor.All(ctx, &users)
	if err != nil {
		return
	}
	return
}

func (dbo *MongoDB) GetUserById(ctx context.Context, id primitive.ObjectID) (user DataStructures.User, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("users").FindOne(ctx, bson.M{"_id":id}).Decode(&user)
	return
}

func (dbo *MongoDB) AddUser(ctx context.Context, user DataStructures.User) (err error){

	_, err = dbo.Client.Database(dbo.Database).Collection("users").InsertOne(ctx, user)
	if err != nil {
		return
	}

	return
}
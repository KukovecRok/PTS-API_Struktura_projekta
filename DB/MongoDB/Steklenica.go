package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (dbo *MongoDB) GetSteklenica(ctx context.Context) (steklenica []DataStructures.Steklenica, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("steklenica").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	steklenica = make([]DataStructures.Steklenica, 0)

	err = cursor.All(ctx, &steklenica)
	if err != nil {
		return
	}
	return
}

func (dbo *MongoDB) GetSteklenicaById(ctx context.Context, id primitive.ObjectID) (steklenica DataStructures.Steklenica, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("steklenica").FindOne(ctx, bson.M{"_id": id}).Decode(&steklenica)
	if err != nil {
		return
	}

	return
}
func (dbo *MongoDB) InsertSteklenica(ctx context.Context, steklenica DataStructures.Steklenica) (err error){

	_, err = dbo.Client.Database(dbo.Database).Collection("steklenica").InsertOne(ctx, steklenica)
	if err != nil {
		return
	}

	return
}

func (dbo *MongoDB) RemoveSteklenica(ctx context.Context, SteklenicaId primitive.ObjectID) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("steklenica").DeleteOne(ctx, bson.M{"_id": SteklenicaId})
	if err != nil {
		return
	}

	return
}
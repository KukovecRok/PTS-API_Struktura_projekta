package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (dbo *MongoDB) GetOcenjevalec(ctx context.Context) (ocenjevalec []DataStructures.Ocenjevalec, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("ocenjevalec").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	ocenjevalec = make([]DataStructures.Ocenjevalec, 0)

	err = cursor.All(ctx, &ocenjevalec)
	if err != nil {
		return
	}
	return
}

func (dbo *MongoDB) GetOcenjevalecById(ctx context.Context, id primitive.ObjectID) (ocenjevalec DataStructures.Ocenjevalec, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("ocenjevalec").FindOne(ctx, bson.M{"_id": id}).Decode(&ocenjevalec)
	if err != nil {
		return
	}

	return
}
func (dbo *MongoDB) InsertOcenjevalec(ctx context.Context, ocenjevalec DataStructures.Ocenjevalec) (err error){

	_, err = dbo.Client.Database(dbo.Database).Collection("ocenjevalec").InsertOne(ctx, ocenjevalec)
	if err != nil {
		return
	}

	return
}

func (dbo *MongoDB) RemoveOcenjevalec(ctx context.Context, ocenjevalecId primitive.ObjectID) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("ocenjevalec").DeleteOne(ctx, bson.M{"_id": ocenjevalecId})
	if err != nil {
		return
	}

	return
}
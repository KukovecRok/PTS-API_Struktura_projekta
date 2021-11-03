package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (dbo *MongoDB) GetProizvajalec(ctx context.Context) (proizvajalec []DataStructures.Proizvajalec, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("proizvajalec").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	proizvajalec = make([]DataStructures.Proizvajalec, 0)

	err = cursor.All(ctx, &proizvajalec)
	if err != nil {
		return
	}
	return
}

func (dbo *MongoDB) GetProizvajalecById(ctx context.Context, id primitive.ObjectID) (proizvajalec DataStructures.Proizvajalec, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("proizvajalec").FindOne(ctx, bson.M{"_id": id}).Decode(&proizvajalec)
	if err != nil {
		return
	}

	return
}
func (dbo *MongoDB) InsertProizvajalec(ctx context.Context, proizvajalec DataStructures.Proizvajalec) (err error){

	_, err = dbo.Client.Database(dbo.Database).Collection("proizvajalec").InsertOne(ctx, proizvajalec)
	if err != nil {
		return
	}

	return
}

func (dbo *MongoDB) RemoveProizvajalec(ctx context.Context, proizvajalecId primitive.ObjectID) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("proizvajalec").DeleteOne(ctx, bson.M{"_id": proizvajalecId})
	if err != nil {
		return
	}

	return
}
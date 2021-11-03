package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (dbo *MongoDB) GetLetnoOcenjevanje(ctx context.Context) (letnoOcenjevanje []DataStructures.LetnoOcenjevanje, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("letnoOcenjevanje").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	letnoOcenjevanje = make([]DataStructures.LetnoOcenjevanje, 0)

	err = cursor.All(ctx, &letnoOcenjevanje)
	if err != nil {
		return
	}
	return
}

func (dbo *MongoDB) GetLetnoOcenjevanjeByLeto(ctx context.Context, leto int) (letnoOcenjevanje DataStructures.LetnoOcenjevanje, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("letnoOcenjevanje").FindOne(ctx, bson.M{"leto": leto}).Decode(&letnoOcenjevanje)
	if err != nil {
		return
	}

	return
}
func (dbo *MongoDB) InsertLetnoOcenjevanje(ctx context.Context, letnoOcenjevanje DataStructures.LetnoOcenjevanje) (err error){

	_, err = dbo.Client.Database(dbo.Database).Collection("letnoOcenjevanje").InsertOne(ctx, letnoOcenjevanje)
	if err != nil {
		return
	}

	return
}

func (dbo *MongoDB) RemoveLetnoOcenjevanje(ctx context.Context, letnoOcenjevanjeId primitive.ObjectID) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("letnoOcenjevanje").DeleteOne(ctx, bson.M{"_id": letnoOcenjevanjeId})
	if err != nil {
		return
	}

	return
}
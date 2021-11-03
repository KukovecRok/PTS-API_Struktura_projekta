package Logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (c *Controller) GetOcenjevalec(ctx context.Context) (ocenjevalec []DataStructures.Ocenjevalec, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetOcenjevalec(ctx)
}

func (c *Controller) GetOcenjevalecById(ctx context.Context, id primitive.ObjectID) (ocenjevalec DataStructures.Ocenjevalec, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetOcenjevalecById(ctx, id)
}
func (c *Controller) InsertOcenjevalec(ctx context.Context, ocenjevalec DataStructures.Ocenjevalec) error {
	return c.db.InsertOcenjevalec(ctx, ocenjevalec)
}
func (c *Controller) RemoveOcenjevalec(ctx context.Context, ocenjevalec primitive.ObjectID) error {
	return c.db.RemoveOcenjevalec(ctx, ocenjevalec)
}
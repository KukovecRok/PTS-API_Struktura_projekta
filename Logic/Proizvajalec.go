package Logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (c *Controller) GetProizvajalec(ctx context.Context) (proizvajalec []DataStructures.Proizvajalec, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetProizvajalec(ctx)
}

func (c *Controller) GetProizvajalecById(ctx context.Context, proizajalecId primitive.ObjectID) (proizvajalec DataStructures.Proizvajalec, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetProizvajalecById(ctx, proizajalecId)
}
func (c *Controller) InsertProizvajalec(ctx context.Context, proizvajalec DataStructures.Proizvajalec) error {
	return c.db.InsertProizvajalec(ctx, proizvajalec)
}
func (c *Controller) RemoveProizvajalec(ctx context.Context, proizvajalec primitive.ObjectID) error {
	return c.db.RemoveProizvajalec(ctx, proizvajalec)
}
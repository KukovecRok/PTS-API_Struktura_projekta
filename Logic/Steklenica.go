package Logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (c *Controller) GetSteklenica(ctx context.Context) (steklenica []DataStructures.Steklenica, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetSteklenica(ctx)
}

func (c *Controller) GetSteklenicaById(ctx context.Context, SteklenicaId primitive.ObjectID) (steklenica DataStructures.Steklenica, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetSteklenicaById(ctx, SteklenicaId)
}
func (c *Controller) InsertSteklenica(ctx context.Context, steklenica DataStructures.Steklenica) error {
	return c.db.InsertSteklenica(ctx, steklenica)
}
func (c *Controller) RemoveSteklenica(ctx context.Context, steklenica primitive.ObjectID) error {
	return c.db.RemoveSteklenica(ctx, steklenica)
}
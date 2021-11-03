package Logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (c *Controller) GetLetnoOcenjevanje(ctx context.Context) (letnoOcenjevanje []DataStructures.LetnoOcenjevanje, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetLetnoOcenjevanje(ctx)
}

func (c *Controller) GetLetnoOcenjevanjeByLeto(ctx context.Context, leto int) (letnoOcenjevanje DataStructures.LetnoOcenjevanje, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetLetnoOcenjevanjeByLeto(ctx, leto)
}
func (c *Controller) InsertLetnoOcenjevanje(ctx context.Context, letnoOcenjevanje DataStructures.LetnoOcenjevanje) error {
	return c.db.InsertLetnoOcenjevanje(ctx, letnoOcenjevanje)
}
func (c *Controller) RemoveLetnoOcenjevanje(ctx context.Context, letnoOcenjevanje primitive.ObjectID) error {
	return c.db.RemoveLetnoOcenjevanje(ctx, letnoOcenjevanje)
}
package Logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)

func (c *Controller) GetUsers(ctx context.Context) (users []DataStructures.User, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetUsers(ctx)

}

func (c *Controller) GetUserById(ctx context.Context, userId primitive.ObjectID) (user DataStructures.User, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetUserById(ctx, userId)
}

func (c *Controller) InsertUser(ctx context.Context, user DataStructures.User) error {
	return c.db.AddUser(ctx, user)
}


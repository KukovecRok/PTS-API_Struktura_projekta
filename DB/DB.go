package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorok/DataStructures"
)



type DB interface {
	Init(ctx context.Context) (err error)

	GetUsers(ctx context.Context) (users []DataStructures.User, err error)
	GetUserById(ctx context.Context, id primitive.ObjectID) (user DataStructures.User, err error)
	AddUser(ctx context.Context, user DataStructures.User) (err error)

	GetLetnoOcenjevanje(ctx context.Context) ([]DataStructures.LetnoOcenjevanje, error)
	GetLetnoOcenjevanjeByLeto(ctx context.Context, leto int) (letnoOcenjevanje DataStructures.LetnoOcenjevanje, err error)
	InsertLetnoOcenjevanje(ctx context.Context, letnoOcenjevanje DataStructures.LetnoOcenjevanje) (err error)
	RemoveLetnoOcenjevanje(ctx context.Context, letnoOcenjevanjeId primitive.ObjectID) (err error)

	GetOcenjevalec(ctx context.Context) ([]DataStructures.Ocenjevalec, error)
	GetOcenjevalecById(ctx context.Context, id primitive.ObjectID) (DataStructures.Ocenjevalec, error)
	InsertOcenjevalec(ctx context.Context, ocenjevalec DataStructures.Ocenjevalec) (err error)
	RemoveOcenjevalec(ctx context.Context, ocenjevalec primitive.ObjectID) (err error)

	GetProizvajalec(ctx context.Context) ([]DataStructures.Proizvajalec, error)
	GetProizvajalecById(ctx context.Context, id primitive.ObjectID) (DataStructures.Proizvajalec, error)
	InsertProizvajalec(ctx context.Context, proizvajalec DataStructures.Proizvajalec) (err error)
	RemoveProizvajalec(ctx context.Context, proizvajalec primitive.ObjectID) (err error)

	GetSteklenica(ctx context.Context) ([]DataStructures.Steklenica, error)
	GetSteklenicaById(ctx context.Context, id primitive.ObjectID) (DataStructures.Steklenica, error)
	InsertSteklenica(ctx context.Context, steklenica DataStructures.Steklenica) (err error)
	RemoveSteklenica(ctx context.Context, steklenica primitive.ObjectID) (err error)
}

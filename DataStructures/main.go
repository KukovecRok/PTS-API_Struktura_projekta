package DataStructures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type MyTime struct {
	time.Time
}

func (self *MyTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	// Get rid of the quotes "" around the value.
	// A second option would be to include them
	// in the date format string instead, like so below:
	//   time.Parse(`"`+time.RFC3339Nano+`"`, s)
	s = s[1 : len(s)-1]

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05.999999999Z0700", s)
	}
	self.Time = t
	return
}

type Ocenjevanje struct {
	IdOcenjevalca primitive.ObjectID `json:"id_ocenjevalca"`
	Ocena         string             `json:"ocena"`
	Cistost       string             `json:"cistost"`
	Bistrost      string             `json:"bistrost"`
	Okus          string             `json:"okus"`
	Harmonicnost  string             `json:"harmonicnost"`
	Vonj          string             `json:"vonj"`
}
type LetoOcenjevanja struct {
	IdLetoOcenjevanja primitive.ObjectID `json:"id"`
	Leto              int                `json:"leto"`
}
type ProizvajalecSteklenica struct {
	IdProizvajalec primitive.ObjectID `json:"id"`
	Naziv          string             `json:"naziv"`
}

type Steklenica struct {
	Id              primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Naziv           string                 `json:"naziv"`
	SortaVina       string                 `json:"sorta_vina"`
	Letnik          int                    `json:"letnik"`
	Proizvajalec    ProizvajalecSteklenica `json:"Proizvajalec"`
	Ocene           []Ocenjevanje          `json:"ocenjevanje"`
	MestoVArhivu    string                 `json:"mesto_v_arhivu"`
	LetoProdaje     int                    `json:"leto_prodaje"`
	LetoOcenjevanja LetoOcenjevanja        `json:"leto_ocenjevanja"`
}
type LetnoOcenjevanje struct {
	Id                     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Leto                   int                `json:"leto"`
	Kraj                   string             `json:"kraj"`
	DatumZacetekTekmovanja time.Time          `json:"datumZacetekTekmovanja"` //ERROR
	DatumKonecTekmovanja   time.Time          `json:"datumKonecTekmovanja"`   //ERROR
}
type Ocenjevalec struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Ime        string             `json:"ime"`
	Priimek    string             `json:"priimek"`
	Email      string             `json:"email"`
	Naslov     string             `json:"naslov"`
	Kraj       string             `json:"kraj"`
	Drzava     string             `json:"drzava"`
	Telefonska string             `json:"telefonska"`
}
type Proizvajalec struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Naziv      string             `json:"naziv"`
	Email      string             `json:"email"`
	Naslov     string             `json:"naslov"`
	Kraj       string             `json:"kraj"`
	Drzava     string             `json:"drzava"`
	Telefonska string             `json:"telefonska"`
	StTock     int                `json:"st_tock"`
}

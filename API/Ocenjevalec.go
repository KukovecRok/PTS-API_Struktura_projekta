package API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todorok/DataStructures"
)

func (a *Controller) GetOcenjevalec(c *gin.Context) {

	//Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	ocenjevalec, err := a.c.GetOcenjevalec(c.Request.Context())
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, ocenjevalec)
}
func (a *Controller) GetOcenjevalecById(c *gin.Context) {

	//Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	ocenjevalec, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	OcenjevalecId, err := a.c.GetOcenjevalecById(c.Request.Context(), ocenjevalec)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, OcenjevalecId)
}

func (a *Controller) InsertOcenjevalec(c *gin.Context) {

	var ocenjevalec DataStructures.Ocenjevalec

	err := c.BindJSON(&ocenjevalec)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.InsertOcenjevalec(c.Request.Context(), ocenjevalec)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "vstavljam ocenjevalca")
}

func (a *Controller) RemoveOcenjevalec(c *gin.Context) {

	ocenjevalec, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = a.c.RemoveOcenjevalec(c.Request.Context(), ocenjevalec)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "odstranjujem ocenjevalca")
}
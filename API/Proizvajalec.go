package API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todorok/DataStructures"
)

func (a *Controller) GetProizvajalec(c *gin.Context) {

	//Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	proizvajalec, err := a.c.GetProizvajalec(c.Request.Context())
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, proizvajalec)
}
func (a *Controller) GetProizvajalecById(c *gin.Context) {

	//Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	proizvajalec, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	proizvajalecId, err := a.c.GetProizvajalecById(c.Request.Context(), proizvajalec)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, proizvajalecId)
}

func (a *Controller) InsertProizvajalec(c *gin.Context) {

	var proizvajalec DataStructures.Proizvajalec

	err := c.BindJSON(&proizvajalec)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.InsertProizvajalec(c.Request.Context(), proizvajalec)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "vstavljam proizvajalca")
}

func (a *Controller) RemoveProizvajalec(c *gin.Context) {

	proizvajalec, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = a.c.RemoveProizvajalec(c.Request.Context(), proizvajalec)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "odstranjujem proizvajalca")
}
package API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todorok/DataStructures"
)

func (a *Controller) GetSteklenica(c *gin.Context) {

	//Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	steklenica, err := a.c.GetSteklenica(c.Request.Context())
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, steklenica)
}
func (a *Controller) GetSteklenicaById(c *gin.Context) {

	//Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	steklenica, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	SteklenicaId, err := a.c.GetSteklenicaById(c.Request.Context(), steklenica)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, SteklenicaId)
}

func (a *Controller) InsertSteklenica(c *gin.Context) {

	var steklenica DataStructures.Steklenica

	err := c.BindJSON(&steklenica)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.InsertSteklenica(c.Request.Context(), steklenica)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "vstavljam steklenico")
}

func (a *Controller) RemoveSteklenica(c *gin.Context) {

	steklenica, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}


	err = a.c.RemoveSteklenica(c.Request.Context(), steklenica)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "odstranjujem steklenico")
}
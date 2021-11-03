package API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"todorok/DataStructures"
)

func (a *Controller) GetLetnoOcenjevanje(c *gin.Context) {

	//Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	letnoOcenjevanje, err := a.c.GetLetnoOcenjevanje(c.Request.Context())
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, letnoOcenjevanje)
}
func (a *Controller) GetLetnoOcenjevanjeByLeto(c *gin.Context) {

	//Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	letnoOcenjevanjeLeto, err := strconv.Atoi(c.Param("leto"))
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	letnoOcenjevanje, err := a.c.GetLetnoOcenjevanjeByLeto(c.Request.Context(), letnoOcenjevanjeLeto)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, letnoOcenjevanje)
}

func (a *Controller) InsertLetnoOcenjevanje(c *gin.Context) {

	var letnoOcenjevanje DataStructures.LetnoOcenjevanje

	err := c.BindJSON(&letnoOcenjevanje)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.InsertLetnoOcenjevanje(c.Request.Context(), letnoOcenjevanje)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "vstavljeno letno ocenjevanje")
}

func (a *Controller) RemoveLetnoOcenjevanje(c *gin.Context) {

	letnoOcenjevanje, err := primitive.ObjectIDFromHex(c.Param("_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = a.c.RemoveLetnoOcenjevanje(c.Request.Context(), letnoOcenjevanje)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "odstranjujem letno ocenjevanje")
}
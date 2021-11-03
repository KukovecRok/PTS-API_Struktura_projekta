package API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todorok/DataStructures"
)

func (a *Controller) GetUsers(c *gin.Context) {

	//Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	users, err := a.c.GetUsers(c.Request.Context())
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, users)
}

func (a *Controller) GetUserById(c *gin.Context) {

	//Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	userId, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	user, err := a.c.GetUserById(c.Request.Context(), userId)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, user)
}
func (a *Controller) InsertUser(c *gin.Context) {

	var user DataStructures.User

	err := c.BindJSON(&user)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.InsertUser(c.Request.Context(), user)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "vstavljen user")
}
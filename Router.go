package main

import (
	"github.com/gin-gonic/gin"
	"todorok/API"
)

//Naredimo Router objekt da na njega obesimo funkcije
type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1
	api := r.engine.Group("/api/v1")

	//Pot /api/v1/user
	user := api.Group("/user")
	r.registerUserRoutes(user)

	letnoOcenjevanje := api.Group("/letnoOcenjevanje")
	r.registerLetnoOcenjevanjeRoutes(letnoOcenjevanje)

	ocenjevalec := api.Group("/ocenjevalec")
	r.registerOcenjevalecRoutes(ocenjevalec)

	proizvajalec := api.Group("/proizvajalec")
	r.registerProizvajalecRoutes(proizvajalec)

	steklenica := api.Group("/steklenica")
	r.registerSteklenicaRoutes(steklenica)

	return

}

func (r *Router) registerUserRoutes(user *gin.RouterGroup) {

	//Pot /api/v1/user
	user.GET("/", r.api.GetUsers)

	//Pot /api/v1/user/:user_id
	// : ni del routea ampak parameter
	user.GET("/:user_id", r.api.GetUserById)

	user.POST("/", r.api.InsertUser)

}


func (r *Router) registerLetnoOcenjevanjeRoutes(letnoOcenjevanje *gin.RouterGroup) {
	letnoOcenjevanje.GET("/", r.api.GetLetnoOcenjevanje)

	letnoOcenjevanje.GET("/:leto", r.api.GetLetnoOcenjevanjeByLeto)

	letnoOcenjevanje.POST("/", r.api.InsertLetnoOcenjevanje)

	letnoOcenjevanje.DELETE("/", r.api.RemoveLetnoOcenjevanje)
}
func (r *Router) registerOcenjevalecRoutes(ocenjevalec *gin.RouterGroup) {
	ocenjevalec.GET("/", r.api.GetOcenjevalec)

	ocenjevalec.GET("/:id", r.api.GetOcenjevalecById)

	ocenjevalec.POST("/", r.api.InsertOcenjevalec)

	ocenjevalec.DELETE("/:id", r.api.RemoveOcenjevalec)
}

func (r *Router) registerProizvajalecRoutes(proizvajalec *gin.RouterGroup) {
	proizvajalec.GET("/", r.api.GetProizvajalec)

	proizvajalec.GET("/:id", r.api.GetProizvajalecById)

	proizvajalec.POST("/", r.api.InsertProizvajalec)

	proizvajalec.DELETE("/:id", r.api.RemoveProizvajalec)
}

func (r *Router) registerSteklenicaRoutes(steklenica *gin.RouterGroup) {
	steklenica.GET("/", r.api.GetSteklenica)

	steklenica.GET("/:id", r.api.GetSteklenicaById)

	steklenica.POST("/", r.api.InsertSteklenica)

	steklenica.DELETE("/:id", r.api.RemoveSteklenica)
}
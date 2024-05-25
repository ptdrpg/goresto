package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/app"
	"github.com/ptdrpg/resto/controller"
	"github.com/ptdrpg/resto/repository"
	"github.com/ptdrpg/resto/router"
)

func main() {
	restgo := gin.Default()
	app.DBconnexion()
	db := app.DB
	repo := repository.NewRepository(db)
	controller := controller.NewController(db, repo)
	r := router.NewRouter(restgo, controller)
	r.RegisterRouter()
	r.R.Run(":4400") // listen and serve on 0.0.0.0:8080
}

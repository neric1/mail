package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neric1/mail/controller"
)

func NewRouter(
	usersController *controller.UsersController,

) *gin.Engine {
	router := gin.Default()
	// add swagger

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	router.POST("/sendEmail", usersController.SendEmail)

	return router
}

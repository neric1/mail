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
	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Next()
	})
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Mailer: welcome home")
	})

	router.POST("/sendEmail", usersController.SendEmail)

	return router
}

package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neric1/mail/controller"
)

func NewRouter(
	usersController *controller.UsersController,

) *gin.Engine {
	router := gin.Default()
	// add swagger
	// CORS middleware
	// router.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	// 	c.Next()
	// })
	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://idsr.afro.who.int"},     // Allow only this origin
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},        // Allow only these HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Allow these headers
		AllowCredentials: true,                                      // Allow credentials (cookies, authorization headers, etc.)
	}

	router.Use(cors.New(corsConfig))
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Mailer: welcome home")
	})

	router.POST("/sendEmail", usersController.SendEmail)

	return router
}

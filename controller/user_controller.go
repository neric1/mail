package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neric1/mail/data/request"
	"github.com/neric1/mail/data/response"
	"github.com/neric1/mail/helper"
	"github.com/rs/zerolog/log"
	gomail "gopkg.in/gomail.v2"
)

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (controller *UsersController) SendEmail(ctx *gin.Context) {
	log.Info().Msg("login Users")
	emailRequest := request.EmailRequestBody{}
	err := ctx.ShouldBindJSON(&emailRequest)
	helper.ErrorPanic(err)
	senderEmail := os.Getenv("APP_EMAIL") // Use your Gmail address
	password := os.Getenv("APP_PASSWORD") // Use an App Password, not your Gmail password
	// recipientEmail := emailRequest.Recipients

	// SMTP server configuration
	smtpHost := os.Getenv("HOST") // Use your SMTP server address
	smtpPort := 587

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", emailRequest.Recipients...) // Change this to recipient's email //"HTML Email Test"
	m.SetHeader("Subject", emailRequest.Subject)  //`<h1>Hello, World!</h1><p>This is an <strong>HTML</strong> email.</p>`
	m.SetBody("text/html", emailRequest.Body)

	// Send the email
	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, password)
	if err := d.DialAndSend(m); err != nil {
		// fmt.Println("", err)
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error sending email: " + err.Error(),
			Data:   d.LocalName,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusFound, webResponse)
		return
	}
	// fmt.Println("Email sent successfully!")
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Email sent successfully",
		Data:   d.LocalName,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusFound, webResponse)

}

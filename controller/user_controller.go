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
	m.SetAddressHeader("From", senderEmail, "WHO-IDSR PLATFORM message [No reply]")
	// m.SetHeader("From", senderEmail)
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
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UsersController) SendEmailWithAttachement(ctx *gin.Context) {
	log.Info().Msg("SendEmail endpoint called")

	// Parse text fields
	recipients := ctx.PostFormArray("recipients")
	subject := ctx.PostForm("subject")
	body := ctx.PostForm("body")

	// Basic validation
	if len(recipients) == 0 || subject == "" || body == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Missing required fields",
			Data:   nil,
		})
		return
	}

	// Load config from environment
	senderEmail := os.Getenv("APP_EMAIL")
	password := os.Getenv("APP_PASSWORD")
	smtpHost := os.Getenv("HOST")
	smtpPort := 587

	// Construct the email message
	m := gomail.NewMessage()
	m.SetAddressHeader("From", senderEmail, "WHO-IDSR PLATFORM Cholera Weekly update")
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Handle uploaded files
	form, err := ctx.MultipartForm()
	if err == nil && form.File != nil {
		files := form.File["attachments"] // input field name should be attachments[]
		for _, fileHeader := range files {
			// Optionally save file locally or attach directly from memory
			tempPath := "./tmp/" + fileHeader.Filename
			if err := ctx.SaveUploadedFile(fileHeader, tempPath); err != nil {
				log.Error().Err(err).Str("file", fileHeader.Filename).Msg("Failed to save uploaded file")
				continue
			}
			m.Attach(tempPath)
		}
	}

	// Send email
	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, password)
	if err := d.DialAndSend(m); err != nil {
		log.Error().Err(err).Msg("Failed to send email")
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed to send email",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Email sent successfully",
		Data:   nil,
	})
}

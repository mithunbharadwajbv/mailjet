package http

import (
	"fmt"
	"golang-docker/domain/mailjet"
	"golang-docker/servivce/mailservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MailHandler interface {
	SendMail(*gin.Context)
}

type mailHandler struct {
	service mailservice.Mailservice
}

func GetNewMailHandler(service mailservice.Mailservice) MailHandler {
	return &mailHandler{
		service: service,
	}
}

func (handler *mailHandler) SendMail(c *gin.Context) {
	to := c.Param("mailid")
	res, err := handler.service.SendMail(mailjet.GetDummyData(to))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("succesfully sent mail to %s \n with response: %v", to, res))
}

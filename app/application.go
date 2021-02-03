package app

import (
	client "golang-docker/client/jetmail"
	"golang-docker/http"
	"golang-docker/servivce/mailservice"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	var mailJetClient = client.GetNewMailJetClient()
	var mailService = mailservice.GetNewMailService(mailJetClient)
	var mailHandler = http.GetNewMailHandler(mailService)

	router.GET("/SendMail/:mailid", mailHandler.SendMail)

	router.Run(":8080")
}

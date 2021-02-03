package mailservice

import (
	client "golang-docker/client/jetmail"

	"github.com/mailjet/mailjet-apiv3-go"
)

type Mailservice interface {
	SendMail(*mailjet.MessagesV31) (*mailjet.ResultsV31, error)
}

type mailService struct {
	client client.MailjetClient
}

func GetNewMailService(client client.MailjetClient) Mailservice {
	return &mailService{
		client: client,
	}
}

func (service *mailService) SendMail(message *mailjet.MessagesV31) (*mailjet.ResultsV31, error) {
	return service.client.SendMail(message)
}

package client

import (
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

var (
	clientval *mailjet.Client

	apiKeyPublic  = "xxx"
	apiKeyPrivate = "xxx"
)

func init() {
	clientval = mailjet.NewMailjetClient(apiKeyPublic, apiKeyPrivate)
}

type MailjetClient interface {
	SendMail(*mailjet.MessagesV31) (*mailjet.ResultsV31, error)
}

type mainJetClient struct {
}

func GetNewMailJetClient() MailjetClient {
	return &mainJetClient{}
}

func (client *mainJetClient) SendMail(message *mailjet.MessagesV31) (*mailjet.ResultsV31, error) {
	res, err := clientval.SendMailV31(message)
	if err != nil {
		fmt.Println("error while sending mail")
		return nil, err
	}
	return res, nil
}

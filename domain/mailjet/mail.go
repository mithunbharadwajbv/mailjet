package mailjet

import (
	"github.com/mailjet/mailjet-apiv3-go"
)

var (
	fromeMail = "xxx@gmail.com"
	fromName  = "xxx"
)

func GetDummyData(toAddress string) *mailjet.MessagesV31 {
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: fromeMail,
				Name:  fromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toAddress,
					Name:  toAddress,
				},
			},
			Subject:  "Greetings from Mailjet.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	return &messages
}

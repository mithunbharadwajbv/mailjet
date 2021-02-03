have take hexagonal architecture to structure my code, code is hightly modular 
and easy to replace with any other component

├── app
│   └── application.go
├── client
│   └── jetmail
│       └── mailjetcliet.go
├── domain
│   └── mailjet
│       └── mail.go
├── go.mod
├── go.sum
├── http
│   └── mail_http.go
├── main.go
├── resource.md
└── servivce
    └── mailservice
        └── mailservice.go


# configration : 

1. should change public and private key of your own in : client/jetmail/mailjetcliet.go
apiKeyPublic  = "xxx"
apiKeyPrivate = "xxx"


2. only to address can be configured throught api call, so configure from address in path : domain/mailjet/mail.go
fromeMail = "xxx@gmail.com"
fromName  = "xxx"

3. run application by command [ go run main.go ]

# api call to run

http://localhost:8080/SendMail/gagansau@gmail.com
give ur mail to to send at last

have attached some screen shotes to demontrate application is running as expected


# comments 

i have not configures loggers as middleware i have all configrations and would love to discuss over a phonecall
not configred jwt (i have xp but need more time)


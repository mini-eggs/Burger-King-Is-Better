# Burger King is Better, Twitter Bot
Just a Twitter bot fighting fake news. Query recent tweets for a hashtag and reply.

#### Prerequisites
1) The Go language and GOPATH/GOROOT are installed and set correctly.
2) Heroku CLI is installed.
3) A Twitter account and associated developer account.

#### Setup
Open terminal/command prompt.

`$ cd go/src/`

`$ git clone https://github.com/mini-eggs/Burger-King-Is-Better.git`

`$ cd Burger-King-Is-Better`

`$ go install`

Create an `.env` file in the projects root. This file will need:
1) Twitter credentials, i.e. `consumerKey`, `consumerSecret`, `accessToken`, and `accessSecret`.
2) MySQL database credentials, i.e. `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_NAME`, and `DB_PORT`.

#### Build
`$ go get`

#### Run CLI examples
`$ heroku local:run bk-is-better --type tweet --text hello`

`$ heroku local:run bk-is-better --type query --hashtag mcdonalds --infinite true`

#### Run server
`$ heroku local web -p 5000`

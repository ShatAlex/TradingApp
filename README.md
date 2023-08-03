# Trading Application

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ShatAlex/TradingApp)
![Static Badge](https://img.shields.io/badge/gin-v1.9.1-brightgreen)
![Static Badge](https://img.shields.io/badge/sqlx-v1.3.5-brown)
![Static Badge](https://img.shields.io/badge/polygon_io-v1.13.1-purple)
![Static Badge](https://img.shields.io/badge/swagger-v1.16.1-orange)



## :sparkles: Project Description
REST API designed according to the rules of Clean Architecture with JWT-based authentication system.

The project focuses on comfortable tracking of securities trades with the integration of a third-party API [(polygon.io)](https://github.com/polygon-io) to obtain the current exchange rates.
___

## :clipboard: Usage
To Run the application use:
```
make run
```
If the application is running for the first time, you need to apply migrations to the db:
```
make migrate
```
___

## :pushpin: API Endpoints

### AUTH
This group of endpoints is intended for user registration and authentication. 

Further functionality is not available for non-authorized users.
##### POST - /auth/sign-up
Example Input:
```
{
    "name": "Alex",
	"username": "user",
	"password": "qwerty"
} 
```
###### POST - /auth/sign-in
Example Input:
```
{
	"username": "user",
	"password": "qwerty"
} 
```
Example Response:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTEwMzY4ODgsImlhdCI6MTY5MTAwODA4OCwidXNlcl9pZCI6MX0.3IKGHaYyDsLIILwMSl6u0Or4lzaOLVD4Zj0MLqehhns"
} 
```
### TYPES
This group of endpoints is designed to perform CRUD operations on all available types of trades.

Since they are the same for all users, this group is available only to the super-user.
##### POST - /api/v1/types/
Example Input:
```
{
  "trade_type": "Payment of commission"
}
```
##### GET - /api/v1/types/ or /api/v1/types/{id}
Example Input:
```
{
  "data": [
    {
      "id": 1,
      "trade_type": "Purchase of securities
    },
    {
      "id": 2,
      "trade_type": "Sale of securities"
    }
  ]
}
```
### TRADES
This group of endpoints is designed to perform CRUD operations on trades, except for buying and selling securities (this functionality is provided in the endpoints group of the portfolio)

##### POST - /api/v1/trades/
Example Input:
```
{
  "amount": 4,
  "price": 10,
  "ticker": "AAPL",
  "type_id": 3
}
```
##### GET - api/v1/trades/ or api/v1/trades/{id}
Example Response:
```
{
  "data": [
    {
      "id": 1,
      "ticker": "AAPL",
      "user_id": 1,
      "type_id": 3,
      "price": 10,
      "amount": 4
    }
  ]
}
```
### PORTFOLIO
In this group, you can buy and sell securities, view the contents of the investment portfolio and receive income information in accordance with the current rates.
##### POST - /api/v1/portfolio/buy or /api/v1/portfolio/sell
Example Input:
```
{
  "amount": 4,
  "ticker": "AAPL"
}
```
##### GET - api/v1/portfolio
Example Response:
```
[
  {
    "id": 1,
    "user_id": 1,
    "ticker": "AAPL",
    "amount": 4
  }
]
```
##### GET - api/v1/portfolio/detail
Example Input:
```
{
  "ticker": "AAPL"
}
```
Example Response:
```
[
  {
    "id": 1,
    "user_id": 1,
    "ticker": "AAPL",
    "amount": 4
  }
]
```
##### GET - api/v1/portfolio/income
Example Response:
```
{
  "total_income": 782.42
}
```
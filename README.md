
# Bitcoin Wallet 

System responsible for managing funds and assets of an account or digital wallet.





## Running locally

Clone the project (or download it manually)

```bash
  git clone https://github.com/mourajj/maxxer.git
```

Create an .env file and set your environment variables according to your MYSQL Workbench configuration:

```
DB_USUARIO=user
DB_SENHA=password
DB_NOME=maxxer

API_PORT=5000
```

Install the dependencies according to the go.mod file

```bash
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
go get github.com/joho/godotenv
```

After that you must be able to run the API by running the command and the API will be listening on the port 5000

```bash
go run main.go
```
## Endpoints 

### /deposit 
Add funds to an account.

```http
  POST /deposit
```

```JSON
{
  "user":"pedro",
  "currency":"BTC",
  "amount": 1
}
```

| Parameter   | Type       | Description                          |
| :---------- | :--------- | :---------------------------------- |
| `user` | `string` | **Mandatory** -  Name of the user |
| `currency` | `string` | **Mandatory** -  Limited to BTC,ETH,ADA,DOGE as requested |
| `amount` | `float` | **Mandatory** -  Amount |


### /withdraw
Remove funds from an account.

```http
  PUT /withdraw
```

```JSON
{
  "user":"pedro",
  "currency":"BTC",
  "amount": 1
}
```

| Parameter   | Type       | Description                          |
| :---------- | :--------- | :---------------------------------- |
| `user` | `string` | **Mandatory** -  Name of the user |
| `currency` | `string` | **Mandatory** -  Limited to BTC,ETH,ADA,DOGE as requested |
| `amount` | `float` | **Mandatory** -  Amount |

### /history
List of all deposits and withdrawals a user made in the last {minutes}

```http
  GET /history/{user}
```

```JSON
{
  "minutes":10
}
```

| Parameter   | Type       | Description                          |
| :---------- | :--------- | :---------------------------------- |
| `user` | `string` | **Mandatory** -  Name of the user |
| `minutes` | `int` | **Mandatory** - Minutes |


### /balance
Shows the current state of an user account

```http
  GET /balance/{user}
```


| Parameter   | Type       | Description                          |
| :---------- | :--------- | :---------------------------------- |
| `user` | `string` | **Mandatory** -  Name of the user |


### /richestperson -- Extra endpoint
Shows the person that has the greater amount of money according to the sum of all of his purchased cryptocurrencies

```http
  GET /richestperson/{user}
```


| Parameter   | Type       | Description                          |
| :---------- | :--------- | :---------------------------------- |
| `user` | `string` | **Mandatory** -  Name of the user |


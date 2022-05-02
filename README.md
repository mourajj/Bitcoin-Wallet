
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

## BONUS / REQUIREMENTS:

### Load Testing - "How many requests per second your endpoints can make?"

There is a folder in this repository called "loadTesting", it contains 5 .js files, each file is related to an endpoint which you can perform load tests by using this following command:

```bash
  k6 run deposit.js
```
To run these tests, you need to install an [open-source load testing tool called K6](https://k6.io/docs/getting-started/installation/) (It's really easy to install).
 
You can specify the amount of Virtual Users and also the time that the load test will run, feel free to check it out =)

### "We don’t like empty databases =)" 

Neither do I, therefore, there is folder called "sql" where you can see 2 files, one of them is the command to create the database and all of its tables, and the other is to insert some data on it.

### "If something wrong happens, how does your system let everybody know it?"

For each unexpected situation, there is HTTP return code and also an error message, feel free to try out!

### "How can we measure the latency of the external APIs calls?"

This can also be done with the K6 tool, if you want to test the external api delay, please refer to the balance.js file and take a look in my comment.

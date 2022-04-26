package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"maxxer/src/models"
	"net/http"
	"strconv"
	"strings"
	_ "time"
)

type JSONData struct {
	Data Data `json:"data"`
}

type Data struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Time struct {
	Data struct {
		Iso   string `json:"iso"`
		Epoch int    `json:"epoch"`
	} `json:"data"`
}

type Funds struct {
	db *sql.DB
}

func NewFundsRepository(db *sql.DB) *Funds {
	return &Funds{db}
}

func (repository Funds) InsertHistory(funds models.Funds) error {
	statement, erro := repository.db.Prepare(
		"insert into funds_history (nickname, currency, amount, transaction_type) values (?, ?, ?, ?);",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	statement.Exec(funds.User, funds.Currency, funds.Amount, "deposit")
	return nil
}

func (repository Funds) Insert(funds models.Funds) error {
	statement, erro := repository.db.Prepare(
		"insert into funds (nickname, currency, amount) values (?, ?, ?) ON DUPLICATE KEY UPDATE amount = amount + ?;",
	)
	if erro != nil {
		fmt.Println(erro)
	}
	defer statement.Close()

	statement.Exec(funds.User, funds.Currency, funds.Amount, funds.Amount)
	return nil
}

func (repository Funds) Withdraw(funds models.Funds) (error, bool) {
	statement, erro := repository.db.Prepare(
		"UPDATE funds SET amount = amount - ? WHERE nickname = ? and currency = ?",
	)
	if erro != nil {
		return erro, false
	}
	defer statement.Close()

	resultado, erro := statement.Exec(funds.Amount, funds.User, funds.Currency)
	if erro != nil {
		return erro, false
	}

	rows, erro := resultado.RowsAffected()
	if rows <= 0 {
		fmt.Println("This user doesn't exist yet or he doesn't have this amount")
		return erro, true
	}
	return erro, false
}

func (repository Funds) CleanDB() {
	statement, erro := repository.db.Prepare(
		"delete from funds where amount <= 0",
	)
	if erro != nil {
		fmt.Println(erro)
	}
	defer statement.Close()
	statement.Exec()
}

func (repository Funds) WithdrawHistory(funds models.Funds) error {
	statement, erro := repository.db.Prepare(
		"insert into funds_history (nickname, currency, amount, transaction_type) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	statement.Exec(funds.User, funds.Currency, funds.Amount, "withdraw")
	return nil
}

func (repository Funds) GetHistory(nickname string, minutos int) ([]models.FundsHistory, error, int) {
	linhas, erro := repository.db.Query(`
		select * from funds_history 
		where nickname = ? and date >= NOW() - INTERVAL ? MINUTE`,
		nickname, minutos,
	)
	if erro != nil {
		return nil, erro, 0
	} else if minutos == 0 {
		return nil, erro, 1
	}

	linhas.Err()
	defer linhas.Close()

	var funds []models.FundsHistory

	for linhas.Next() {
		var fund models.FundsHistory

		if erro = linhas.Scan(
			&fund.ID,
			&fund.User,
			&fund.Currency,
			&fund.TransactionType,
			&fund.Amount,
			&fund.Date,
		); erro != nil {
			return nil, erro, 0
		}

		dados := strings.Split(fund.Date, "T")
		dados[1] = dados[1][:len(dados[1])-6]
		fund.Date = dados[0] + " " + dados[1]
		funds = append(funds, fund)
	}

	if len(funds) == 0 {
		return nil, erro, 2
	}

	return funds, nil, 0
}

func (repository Funds) GetBalance(nickname string) ([]models.Balance, error) {
	linhas, erro := repository.db.Query(`
	SELECT currency, amount FROM funds where nickname = ?`,
		nickname,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var balances []models.Balance

	for linhas.Next() {
		var balance models.Balance

		if erro = linhas.Scan(
			&balance.Currency,
			&balance.Amount,
		); erro != nil {
			return nil, erro
		}

		balances = append(balances, balance)
	}

	// ------ Getting the data from the EXTERNAL API

	stringUSD := "https://api.coinbase.com/v2/prices/?-USD/spot"
	stringEUR := "https://api.coinbase.com/v2/prices/?-EUR/spot"

	var TotalDol float64 = 0
	var TotalEur float64 = 0
	for i := range balances {

		stringUSD = strings.Replace(stringUSD, "?", balances[i].Currency, 1)
		stringEUR = strings.Replace(stringEUR, "?", balances[i].Currency, 1)

		var retornoUSD JSONData
		response, err := http.Get(stringUSD)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(data, &retornoUSD)
		}

		data, _ := strconv.ParseFloat(retornoUSD.Data.Amount, 32)
		balances[i].PriceInDollars = data
		balances[i].TotalDollars = (balances[i].Amount * data)
		TotalDol = TotalDol + balances[i].TotalDollars

		var retornoEUR JSONData
		response, err = http.Get(stringEUR)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(data, &retornoEUR)
		}
		data, _ = strconv.ParseFloat(retornoEUR.Data.Amount, 32)
		balances[i].PriceInEuros = data
		balances[i].TotalEuros = (balances[i].Amount * data)
		TotalEur = TotalEur + balances[i].TotalEuros

		var tempo Time
		response, err = http.Get("https://api.coinbase.com/v2/time")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(data, &tempo)
		}
		tempo.Data.Iso = tempo.Data.Iso[:len(tempo.Data.Iso)-1]
		dados := strings.Split(tempo.Data.Iso, "T")
		balances[i].TimeOfRateUsed = dados[0] + " " + dados[1]

		stringUSD = "https://api.coinbase.com/v2/prices/?-USD/spot"
		stringEUR = "https://api.coinbase.com/v2/prices/?-EUR/spot"
	}

	var balance models.Balance
	balance.TotalDollars = TotalDol
	balance.TotalEuros = TotalEur

	balances = append(balances, balance)

	return balances, nil
}

func (repository Funds) GetUsers() ([]string, error) {
	linhas, erro := repository.db.Query(`
	select distinct nickname from funds`,
	)
	if erro != nil {
		return nil, erro
	}

	linhas.Err()
	defer linhas.Close()

	users := []string{}

	for linhas.Next() {
		var user string

		if erro = linhas.Scan(
			&user,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

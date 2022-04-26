package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"maxxer/src/banco"
	"maxxer/src/models"
	"maxxer/src/repository"
	"maxxer/src/responses"
	"net/http"

	"github.com/gorilla/mux"
)

// localhost:5000/deposit
func Deposit(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		responses.JSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	var funds models.Funds
	if erro = json.Unmarshal(corpoRequisicao, &funds); erro != nil {
		fmt.Println(erro)
		responses.JSON(w, http.StatusBadRequest, funds)
		return
	}

	availablecoins := []string{"btc", "eth", "ada", "doge", "BTC", "ETH", "ADA", "DOGE"}
	exist := false

	if funds.User == "" {
		responses.JSON(w, http.StatusBadRequest, "The user is not specified")
		return
	} else if funds.Amount <= 0 {
		responses.JSON(w, http.StatusBadRequest, "The amount is not specified or is less than 0")
		return
	}

	for _, x := range availablecoins {
		if funds.Currency == x {
			exist = true
		}
	}

	if !exist {
		responses.JSON(w, http.StatusBadRequest, "The currency doesn't exist in our list or is not specified correctly.")
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		fmt.Println(erro)
		responses.JSON(w, http.StatusBadRequest, funds)
		return
	}
	defer db.Close()

	repository := repository.NewFundsRepository(db)
	repository.Insert(funds)
	repository.InsertHistory(funds)

	responses.JSON(w, http.StatusCreated, funds)
}

// localhost:5000/withdraw
func Withdraw(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		responses.JSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	var funds models.Funds

	if erro = json.Unmarshal(corpoRequisicao, &funds); erro != nil {
		responses.JSON(w, http.StatusBadRequest, "The parameters are not correct.")
		return
	} else if funds.Amount <= 0 {
		responses.JSON(w, http.StatusBadRequest, "The amount is not specified or is less than 0")
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer db.Close()

	var contain bool

	repository := repository.NewFundsRepository(db)
	erro, contain = repository.Withdraw(funds)
	if erro != nil || contain {
		responses.JSON(w, http.StatusBadRequest, "This user doesn't exist in our database yet or he doesn't have this amount")
		fmt.Println(erro)
		return
	}
	repository.WithdrawHistory(funds)
	repository.CleanDB()
	responses.JSON(w, http.StatusOK, funds)
}

// localhost:5000/balance/user
func Balance(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		responses.JSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	parametros := mux.Vars(r)
	nickname := parametros["user"]

	db, erro := banco.Connect()
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer db.Close()

	repository := repository.NewFundsRepository(db)
	history, erro := repository.GetBalance(nickname)
	if erro != nil {
		fmt.Println(erro)
		responses.JSON(w, http.StatusNotFound, history)
		return
	} else if len(history) < 2 {
		responses.JSON(w, http.StatusNotFound, "We couldn't find any balance for this specific user.")
		return
	}

	responses.JSON(w, http.StatusOK, history)
}

// localhost:5000/history/user
func History(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		responses.JSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	parametros := mux.Vars(r)
	nickname := parametros["user"]

	type TimeGiven struct {
		Minutes int `json:"minutes,omitempty"`
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	var minutos TimeGiven

	if erro = json.Unmarshal(corpoRequisicao, &minutos); erro != nil {
		responses.JSON(w, http.StatusNotFound, "Error when reading the JSON data.")
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer db.Close()

	repository := repository.NewFundsRepository(db)
	history, erro, empty := repository.GetHistory(nickname, minutos.Minutes)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	switch empty {
	case 1:
		responses.JSON(w, http.StatusBadRequest, "The parameter minutes needs to be passed and must be greather than 0.")
		return
	case 2:
		responses.JSON(w, http.StatusNotFound, "This user hasn't made any deposits or withdraws in this time given.")
		return
	}

	responses.JSON(w, http.StatusOK, history)
}

// localhost:5000/richestperson
func RichestPerson(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		responses.JSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
	db, erro := banco.Connect()
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer db.Close()

	repository := repository.NewFundsRepository(db)
	users, erro := repository.GetUsers()
	if erro != nil {
		fmt.Println(erro)
		return
	}

	type Richest struct {
		User         string `json:"user,omitempty"`
		TotalDollars string `json:"total_dollars,omitempty"`
		TotalEuros   string `json:"total_euros,omitempty"`
	}

	var richest Richest
	var highestValue float64 = 0

	for _, x := range users {
		user, erro := repository.GetBalance(x)
		if erro != nil {
			fmt.Println(erro)
			return
		}

		if user[len(user)-1].TotalDollars > highestValue {
			highestValue = user[len(user)-1].TotalDollars
			richest.User = x
			richest.TotalDollars = fmt.Sprintf("%f", highestValue)
			richest.TotalEuros = fmt.Sprintf("%f", user[len(user)-1].TotalEuros)
		}
	}
	responses.JSON(w, http.StatusOK, richest)

}

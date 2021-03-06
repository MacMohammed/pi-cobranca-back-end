package controllers

import (
	"encoding/json"
	"fatec/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"fatec/respostas"

	"github.com/gorilla/mux"
)

func AllTransaction(w http.ResponseWriter, r *http.Request) {
	dt_inicial := r.URL.Query().Get("dt-ini")
	dt_final := r.URL.Query().Get("dt-fin")

	if len(dt_inicial) > 0 && len(dt_final) > 0 {
		transactions, err := models.AllTransactionForPeriod(dt_inicial, dt_final)
		if err != nil {
			respostas.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		respostas.JSON(w, http.StatusOK, transactions)
		return
	}

	transactions := models.AllTransaction()
	json.NewEncoder(w).Encode(transactions)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetTransaction(transactions.Id_transaction)

	json.NewEncoder(w).Encode(ret)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		fmt.Print(err)
	}

	ret := models.UpdateTransaction(transactions)
	json.NewEncoder(w).Encode(ret)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transacao
	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &transactions); err != nil {
		fmt.Print(err)
	}

	if err := models.NewTransaction(transactions); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, "Transação cadastrada com sucesso")
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		fmt.Print(err)
	}

	models.DeleteTransaction(transactions)
	json.NewEncoder(w).Encode(0)
}

func TransacaoBaixar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id_transacao, err := strconv.ParseUint(parametros["id_transacao"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if err = models.BaixarTitulo(id_transacao); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("O título %d foi baixado com sucesso.", id_transacao))
}

func TransacaoCancelar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id_transacao, err := strconv.ParseUint(parametros["id_transacao"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if err = models.CancelarTitulo(id_transacao); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("O título %d foi cancelado com sucesso.", id_transacao))
}

func TransacaoExtornar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id_transacao, err := strconv.ParseUint(parametros["id_transacao"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	if err = models.ExtornarTitulo(id_transacao); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("O título %d foi extornado com sucesso.", id_transacao))
}

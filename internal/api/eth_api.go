package api

// Copyright (C) 2020 ConsenSys Software Inc.

// This file takes a REST call on the /eth API, error checks the parameters,
// and does the call to the eth code.

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

func getEthBalance(r *restful.Request, w *restful.Response) {
	queryValues := r.Request.URL.Query()
	account := queryValues.Get("Account")
	log.Printf("Account: %s\n", account)
	if len(account) == 0 {
		w.WriteErrorString(http.StatusBadRequest, "Invalid parameters. Parameter values: Account=account")
		return
	}
	if !isB64OrSimpleASCII(account) {
		w.WriteErrorString(http.StatusBadRequest, "Account not Base64Url or ASCII encoded")
		return
	}

	balance := 7
	// balance := eth.GetBalance(account);
	log.Printf("Balance: %d\n", balance)
	w.WriteAsJson(&balance)
}

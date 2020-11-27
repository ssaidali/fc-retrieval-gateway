package api

// Copyright (C) 2020 ConsenSys Software Inc

import (
	"net/http"
	"strings"
	"syscall"

	restful "github.com/emicklei/go-restful/v3"
)

func getEnv(r *restful.Request, w *restful.Response) {
	env := syscall.Environ()

	queryValues := r.PathParameters()
	numQueryKeyValuePairs := len(queryValues)

	name := queryValues["name"]

	if len(name) != 0 {
		// "name" key-value pair exists.
		if numQueryKeyValuePairs != 1 {
			w.WriteErrorString(http.StatusBadRequest, "Invalid parameter. Parameter values: name or no parameter")
			return
		}

		if !isB64OrSimpleASCII(name) {
			w.WriteErrorString(http.StatusBadRequest, "Name not Base64Uri encoded or ASCII")
			return
		}

		// Get a single environment variable.
		for _, s := range env {
			nameVal := strings.Split(s, "=")
			if nameVal[0] == name {
				value := nameVal[1]
				w.WriteAsJson(&value)
				return
			}

		}

		errMsg := name + " not found"
		w.WriteAsJson(&errMsg)
		return
	}
	// return all environment variables.

	if numQueryKeyValuePairs != 0 {
		w.WriteErrorString(http.StatusBadRequest, "Invalid parameter. Parameter values: name or no parameter")
		return
	}

	w.WriteAsJson(&env)
}

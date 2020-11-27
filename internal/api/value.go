package api

// Copyright (C) 2020 ConsenSys Software Inc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ConsenSys/fc-retrieval-gateway/internal/storage"
	restful "github.com/emicklei/go-restful/v3"
)

// KeyValue holds keys and values that are stored
type KeyValue struct {
	Key   string
	Value string
}

func setValue(r *restful.Request, w *restful.Response) {
	log.Println("setValue: start")
	//TODO this doesn't compile with new version of Go	log.Printf("setValue: request: \n%s\n", r)

	key, value, e := extractKeyValue(r, w)

	if e {
		log.Println("setValue: Extract key value pair error")
		// NOTE: HTTP error already set-up.
		return
	}
	if !isB64OrSimpleASCII(key) {
		log.Println("setValue: key not Base64Url or ASCII encoded")
		w.WriteErrorString(http.StatusBadRequest, "Name not Base64Uri encoded or ASCII")
		return
	}
	if !isB64OrSimpleASCII(value) {
		log.Println("setValue: value not Base64Url or ASCII encoded")
		w.WriteErrorString(http.StatusBadRequest, "Value not Base64Uri encoded or ASCII")
		return
	}

	store := getStorage()
	//store := storage.GetKeyValueStorage()

	store.Put(key, value)

	w.WriteHeader(http.StatusOK)

	log.Println("setValue: done")
}

func getKeyValues(r *restful.Request, w *restful.Response) {
	store := getStorage()

	queryValues := r.Request.URL.Query()
	numQueryKeyValuePairs := len(queryValues)

	key := queryValues.Get("Key")
	log.Printf("Key: %s\n", key)
	if len(key) != 0 {
		if numQueryKeyValuePairs != 1 {
			w.WriteErrorString(http.StatusBadRequest, "Invalid parameter. Parameter values: Key=key, or no parameters")
			return
		}

		if !isB64OrSimpleASCII(key) {
			w.WriteErrorString(http.StatusBadRequest, "Key not Base64Url or ASCII encoded")
			return
		}

		value, exists := store.GetValue(key)
		//log.Printf("value: %s\n", value)
		if exists {
			w.WriteAsJson(&value)
			return
		}
		value = "ERROR: Key value not set"
		w.WriteAsJson(&value)
		return
	}
	// return all keys.

	if numQueryKeyValuePairs != 0 {
		w.WriteErrorString(http.StatusBadRequest, "Invalid parameter. Parameter values: Key=key, or no parameters")
		return
	}

	keys := store.GetKeys()
	w.WriteAsJson(&keys)
}

func getStorage() storage.Storage {
	store := storage.GetSingleInstance(storage.Redis)
	return *store
}

func extractKeyValue(r *restful.Request, w *restful.Response) (k, v string, e bool) {
	log.Println("ExtractKeyValue: start")

	e = true

	keyValue := KeyValue{}

	err := decodeJSONPayload(r, &keyValue)

	if err != nil {
		log.Println("Error: ", err.Error())
		w.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	log.Printf("Key: %s, Value: %s\n", keyValue.Key, keyValue.Value)
	k = keyValue.Key
	v = keyValue.Value

	log.Println("ExtractKeyValue: done")

	e = false
	return
}

func decodeJSONPayload(r *restful.Request, v interface{}) error {
	content, err := ioutil.ReadAll(r.Request.Body)
	r.Request.Body.Close()
	if err != nil {
		return err
	}
	if len(content) == 0 {
		return errors.New("JSON payload is empty")
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		return err
	}
	return nil
}

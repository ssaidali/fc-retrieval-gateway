package clientapi

// Copyright (C) 2020 ConsenSys Software Inc
import (
	"log"
	"net/http"

	"github.com/ConsenSys/fc-retrieval-gateway/pkg/messages"
	"github.com/ant0ine/go-json-rest/rest"
)

// HandleClientStandardCIDDiscover is used to handle client request for cid offer
func (g *ClientAPI) HandleClientStandardCIDDiscover(w rest.ResponseWriter, r *rest.Request) {
	payload := messages.ClientStandardDiscoverRequest{}
	err := r.DecodeJsonPayload(&payload)
	if err != nil {
		log.Println(err.Error())
		rest.Error(w, "Fail to decode payload.", http.StatusBadRequest)
		return
	}
	// TODO: For now just print the payload
	log.Println(payload)

	// Dummy response
	response := messages.ClientStandardDiscoverResponse{MessageType: messages.ClientStandardDiscoverResponseType}
	w.WriteJson(response)
}
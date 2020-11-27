package api

// Copyright (C) 2020 ConsenSys Software Inc

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/ConsenSys/fc-retrieval-gateway/internal/util"
	restful "github.com/emicklei/go-restful/v3"
)

// StartRestAPI starts the REST API as a separate go routine.
// Any start-up errors are returned.
func StartRestAPI(settings util.AppSettings) error {
	// Start the REST API and block until the error code is set.
	errChan := make(chan error, 1)
	go startRestAPI(settings, errChan)
	return <-errChan
}

func startRestAPI(settings util.AppSettings, errChannel chan<- error) {
	service := new(restful.WebService)
	service.
		Path("/").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).
		Route(service.GET("/version").To(checkVersion)). // Get code version
		Route(service.GET("/id").To(showID)).            // Get something to show which service has been connected with.

		Route(service.GET("/env").To(getEnv)). // Get environment variable(s).
		// /env returns all environment variables.
		// /env?name=<env> returns the environment variable env

		Route(service.GET("/time").To(getTime)). // Get system time.

		Route(service.GET("/ip").To(getIP)).         // Get IP address.
		Route(service.GET("/host").To(getHostname)). // Get host name.

		// Route(service.GET("/config").To(getConfig)). // Get the current config.

		Route(service.POST("/value").To(setValue)).   // Set a value.
		Route(service.GET("/value").To(getKeyValues)) // Get a value given a key or a list of all the keys.

	log.Println("Running REST API on: " + settings.BindRestAPI)
	restful.Add(service)
	errChannel <- nil
	log.Fatal(http.ListenAndServe(":"+settings.BindRestAPI, nil))
}

func checkVersion(r *restful.Request, w *restful.Response) {
	v := util.GetVersion()
	w.WriteAsJson(&v)
}

func getTime(r *restful.Request, w *restful.Response) {
	w.WriteAsJson(time.Now())
}

func getHostname(r *restful.Request, w *restful.Response) {
	name, err := os.Hostname()
	if err != nil {
		log.Printf("Get host name1: %v\n", err)
		return
	}

	w.WriteAsJson(name)
}

func getIP(r *restful.Request, w *restful.Response) {
	name, err := os.Hostname()
	if err != nil {
		log.Printf("Get host name2: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Printf("Lookup host: %v\n", err)
		return
	}

	w.WriteAsJson(addrs)
}

func showID(r *restful.Request, w *restful.Response) {
	w.WriteAsJson("GATEWAY")
}

func ping(r *restful.Request, w *restful.Response) {
	// TODO check that the request includes the word "PING"
	w.WriteAsJson("PONG")
}

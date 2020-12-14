package main

import (
	"log"

	"github.com/ConsenSys/fc-retrieval-gateway/internal/api"
	"github.com/ConsenSys/fc-retrieval-gateway/internal/api/gatewayapi"
	"github.com/ConsenSys/fc-retrieval-gateway/internal/api/providerapi"
	"github.com/ConsenSys/fc-retrieval-gateway/internal/gateway"
	"github.com/ConsenSys/fc-retrieval-gateway/internal/util/settings"
	"github.com/ConsenSys/fc-retrieval-gateway/internal/util"
	"github.com/ConsenSys/fc-retrieval-gateway/pkg/logging"

)


func main() {
	settings := settings.LoadSettings()
	logging.Info("Filecoin Gateway Start-up: Started")

	_, err := gateway.Create(settings)
	if err != nil {
		log.Println("Error starting server: Client REST API: " + err.Error())
		return
	}

	// Initialise a dummy gateway instance.
	g1 := api.Gateway{ProtocolVersion: 1, ProtocolSupported: []int{1, 2}}

	err = gatewayapi.StartGatewayAPI(settings, &g1)
	if err != nil {
		log.Println("Error starting gateway tcp server: " + err.Error())
		return
	}

	err = providerapi.StartProviderAPI(settings, &g1)
	if err != nil {
		log.Println("Error starting provider tcp server: " + err.Error())
		return
	}

	// Configure what should be called if Control-C is hit.
	util.SetUpCtrlCExit(gracefulExit)

	logging.Info("Filecoin Gateway Start-up Complete")

	// Wait forever.
	select {}
}

func gracefulExit() {
	logging.Info("Filecoin Gateway Shutdown: Start")

	logging.Error("graceful shutdown code not written yet!")
	// TODO

	logging.Info("Filecoin Gateway Shutdown: Completed")
}

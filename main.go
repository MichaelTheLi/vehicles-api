package main

import (
	"log"
	"net/http"
	"vehicles_api/adapters"
	"vehicles_api/api"
)

func main() {
	apiConfig := api.BuildConfigFromEnv()
	adaptersConfig := adapters.BuildConfigFromEnv()
	app := adapters.BuildApp(adaptersConfig)

	serverState := adapters.BuildServerState(app)

	router := api.BuildServer(serverState, apiConfig)
	log.Fatal(http.ListenAndServe(":"+apiConfig.Port, router))
}

package main

import (
	"log"
	"os"
	"plandex-server/routes"
	"plandex-server/setup"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	log.Println("Error loading .env filehjjhhh")
	
	// Configure the default logger to include milliseconds in timestamps
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	routes.RegisterHandlePlandex(func(router *mux.Router, path string, isStreaming bool, handler routes.PlandexHandler) *mux.Route {
		return router.HandleFunc(path, handler)
	})

	r := mux.NewRouter()
	routes.AddHealthRoutes(r)
	routes.AddApiRoutes(r)
	routes.AddProxyableApiRoutes(r)
	setup.MustLoadIp()
	setup.MustInitDb()
	setup.StartServer(r, nil, nil)
	os.Exit(0)
}

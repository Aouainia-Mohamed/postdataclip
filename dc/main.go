package main

import (
	"dataclips/Helpers"
	"dataclips/Routers"
	_ "dataclips/docs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Dataclip post api
// @version 1.0
// @description This is a sample service for creating dataclips
// @termsOfService http://swagger.io/terms/
// @contact.name Awaniya Mohamed Elmabrouk
// @contact.email mohamed@craftfoundry.tech
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /

func main() {
	Helpers.Migration()
	r := Routers.InitRouter()
	port := GetPort()
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	err := http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"https://frontdataclip.herokuapp.com", "http://localhost:3000"}))(r))
	if err != nil {
		panic(err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}

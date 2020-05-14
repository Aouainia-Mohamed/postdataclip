package Routers

import (
	"dataclips/Controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/Dataclips", Controllers.GetDataclips).Methods("GET")
	r.HandleFunc("/Dataclips", Controllers.CreateDataclip).Methods("POST")

	return r

}

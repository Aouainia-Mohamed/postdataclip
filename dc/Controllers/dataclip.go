package Controllers

import (
	"dataclips/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	h "dataclips/Helpers"

	"github.com/jinzhu/gorm/dialects/postgres"
)

var Dataclips []Models.Dataclip

type Data struct {
	Dataclipkey string
	Groupename  postgres.Jsonb
	Name        postgres.Jsonb
	Description postgres.Jsonb
	//Descriptionarg postgres.Jsonb `json:"argdesc"`
	Sqltext  string
	Nargs    int
	Saastype string
}

// Getdataclips godoc
// @Summary list all dataclips
// @Description list all available dataclips
// @Tags back_end_dataclips
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.Dataclip
// @Router /Dataclips [get]
func GetDataclips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := h.DbConnect()
	err := db.Debug().Model(&Models.Dataclip{}).Find(&Dataclips).Error
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&Dataclips)
}

// Createdataclip Handler
// @Summary Create a new dataclip
// @Description Create a new dataclip with the input payload
// @Tags back_end_dataclips
// @Accept  json
// @Produce  json
// @Param dataclip body Models.Dataclip false "Dataclip"
// @Success 200 {object} Models.Dataclip
// @Router /Dataclips [post]
func CreateDataclip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		fmt.Println("inputs error", err2)
	}
	var data Data
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		fmt.Println("unmarchal error", err1)
	}
	decoder := json.NewDecoder(r.Body)
	var DC Models.Dataclip

	err := decoder.Decode(&DC)
	if err != nil {
		fmt.Println("decode error", err)
	}
	log.Println("new Dataclip :", DC)
	db := h.DbConnect()
	db.NewRecord(DC)
	db.Create(&DC)

	// _ = json.NewDecoder(r.Body).Decode(&DC)
	// //fmt.Println(&DC)
	// Dataclips = append(Dataclips, DC)
	json.NewEncoder(w).Encode(DC)

}

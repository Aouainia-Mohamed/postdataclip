package Controllers

import (
	"bytes"
	"dataclips/Models"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatedataclip(t *testing.T) {
	var jsonStr = []byte(`{"Dataclipkey":"dc1","GroupeName":"{en:drivers,fr:chauffeur}","Name":"{en:drivers,fr:chauffeur}","Description":"{en:drivers,fr:chauffeur}","Sqltext":"select * from public.drivers","Nargs":3,"ArgDesc":"{en:drivers,fr:chauffeur}","SaasType":"string"}`)
	request, err := http.NewRequest("POST", "/Dataclips", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateDataclip)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetDataclips(t *testing.T) {

	// This creates a new request to the /entries endpoint.
	req, err := http.NewRequest("GET", "/Dataclips", nil)

	if err != nil {
		t.Fatal(err)
	}

	//Creates a new recorder to record the response received by the entries endpoint.
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetDataclips)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {

		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.

	var result []Models.Dataclip
	err = json.Unmarshal([]byte(rr.Body.String()), &result)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	if len(result) == 0 {
		t.Errorf("handler returned empty data %v",
			len(result))
	}
}

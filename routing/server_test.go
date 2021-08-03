package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"skeleton.service/cars/database"
	"skeleton.service/cars/handlers"
	"skeleton.service/cars/service"
	db "skeleton.service/database"
	"skeleton.service/env"
	"skeleton.service/validators"
	"strings"
	"testing"
)

const (
	SubaruForesterKey = "31230o3012o3"
)

func TestRouter(t *testing.T) {
	t.Run("handlers.GetCar: happy path", GetCarHappyPath)
	t.Run("handlers.GetCar: id is required", GetCarIdIsRequired)
	t.Run("handlers.GetCar: id is required", GetCarCarNotFound)
	t.Run("handlers.PostCar: happy path", PostCarHappyPath)
	t.Run("handlers.PostCar: brand and model is required", PostCarBrandAndModelIsRequired)
	t.Run("handlers.PostCar: brand and model is too long", PostCarBrandAndModelIsTooLong)
}

func initTest() {
	env.MockEnabled = true
	env.DBCarsCollectionName = "cars"
	db.Init()
	validators.Init()
}

func appendSubaruForesterCar() {
	c := &database.Car{
		ID:      SubaruForesterKey,
		Brand:   "Subaru",
		Model:   "Forester",
		Created: "2021-06-30 12:07:18",
	}

	database.AppendCarToMock(c)
}

func GetCarHappyPath(t *testing.T) {
	initTest()
	appendSubaruForesterCar()

	url := fmt.Sprintf("/api/car?id=%s", SubaruForesterKey)

	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(`{}`))
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetCar)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"data":{"id":"31230o3012o3","brand":"Subaru","model":"Forester","created":"2021-06-30 12:07:18"},"_embedded":{},"_links":{},"_status":"completed"}`
	// Check the response body is what we expect.
	if string(rr.Body.Bytes()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(rr.Body.Bytes()), expected)
	}
}

func GetCarIdIsRequired(t *testing.T) {
	initTest()
	appendSubaruForesterCar()

	url := fmt.Sprintf("/api/car?id=%s", "")

	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(`{}`))
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetCar)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":["Query param 'id' is required"]}`
	// Check the response body is what we expect.
	if string(rr.Body.Bytes()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(rr.Body.Bytes()), expected)
	}
}

func GetCarCarNotFound(t *testing.T) {
	initTest()
	appendSubaruForesterCar()
	docID := driver.NewDocumentID(env.DBCarsCollectionName, "someID")

	url := fmt.Sprintf("/api/car?id=%s", docID)

	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(`{}`))
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("api_version", "v1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetCar)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	expected := `{"message":"Car Not Found to provide ID. cars/someID"}`
	// Check the response body is what we expect.
	if string(rr.Body.Bytes()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(rr.Body.Bytes()), expected)
	}
}

func PostCarHappyPath(t *testing.T) {
	initTest()

	expectedBrand := "Subaru"
	expectedModel := "forester"
	expectedStatus := "completed"

	payload := fmt.Sprintf(`
		{
			"brand": "%s", 
			"model": "%s"
		}
	`, expectedBrand, expectedModel)

	req, err := http.NewRequest(http.MethodPost, "/api/car", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostCar)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var postResponse service.HalResponse
	if err = json.Unmarshal(rr.Body.Bytes(), &postResponse); err != nil {
		t.Fatal(err)
	}

	// Check the response body is what we expect.
	if postResponse.Data.Brand != expectedBrand {
		t.Errorf("handler returned unexpected body: got %v want %v", postResponse.Data.Brand, expectedBrand)
	}
	if postResponse.Data.Model != expectedModel {
		t.Errorf("handler returned unexpected body: got %v want %v", postResponse.Data.Model, expectedModel)
	}
	if postResponse.Status != expectedStatus {
		t.Errorf("handler returned unexpected body: got %v want %v", postResponse.Status, expectedStatus)
	}
}

func PostCarBrandAndModelIsRequired(t *testing.T) {
	initTest()

	req, err := http.NewRequest(http.MethodPost, "/api/car", strings.NewReader(`{}`))
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostCar)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"message":["brand_is_required","model_is_required"]}`
	// Check the response body is what we expect.
	if string(rr.Body.Bytes()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(rr.Body.Bytes()), expected)
	}
}

func PostCarBrandAndModelIsTooLong(t *testing.T) {
	initTest()

	payload := `
		{
			"brand": "e329jd3j24d34 ijfuj34ufj 934jfj39fj9ijer ifjiufj eruhjf lkd flkdlk gjlkjdfgljk dlkjfhg ljkdfljg ljkdfg", 
			"model": "e329jd3j24d34 ijfuj34ufj 934jfj39fj9ijer ifjiufj eruhjf lkd flkdlk gjlkjdfgljk dlkjfhg ljkdfljg ljkdfg"
		}
	`
	req, err := http.NewRequest(http.MethodPost, "/api/car", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostCar)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"message":["brand_too_long","model_too_long"]}`
	// Check the response body is what we expect.
	if string(rr.Body.Bytes()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(rr.Body.Bytes()), expected)
	}
}

package mongo

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTestFunc(t *testing.T) {
	result := TestFunc(1, 3)
	expected := 4

	if result != expected {
		t.Errorf("got %q, wanted %q", result, expected)
	}
 }
 
 func setupHttpTests() (*httptest.ResponseRecorder, *gin.Context) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	return recorder, c
 }
 
 // TODO: godotenv mocking, mongo mocking

 func TestGetAllDates(t *testing.T) {
	recorder, c := setupHttpTests()

	// expectedBody := []primitive.M{"name": "Matt","date": "June 18","type": "Birthday"}

	GetAllDates(c)

	response := recorder.Result()

	if response.StatusCode != 200 {
		t.Error(recorder.Code, "Failed test")
	}
 }

 func TestAddNewDate(t *testing.T) {
	recorder, c := setupHttpTests()

	AddNewDate(c)

	response := recorder.Result()

	if response.StatusCode != 200 {
		t.Error(recorder.Code, "Failed test")
	}
 }

 func TestRemoveDate(t *testing.T) {
	recorder, c := setupHttpTests()

	RemoveDate(c)

	response := recorder.Result()

	if response.StatusCode != 200 {
		t.Error(recorder.Code, "Failed test")
	}
 }

 func TestWipeDatabase (t *testing.T) {
	recorder, c := setupHttpTests()

	WipeDatabase(c)

	response := recorder.Result()

	if response.StatusCode != 200 {
		t.Error(recorder.Code, "Failed test")
	}
 }
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

 func TestGetAllDates(t *testing.T) {
	// TODO: godotenv mocking, mongo mocking
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	GetAllDates(c)

	if recorder.Code != 200 {
		t.Error(recorder.Code, "Failed test")
	}
 }
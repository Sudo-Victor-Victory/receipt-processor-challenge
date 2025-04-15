package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/receipts/process", add_receipt)
    return r
}

func TestAddReceipt_SUCCESS(t *testing.T) {
	// 																GIVEN
    router := setupRouter()
    // Create a sample receipt to add
    receipt := Receipt{
        Retailer:     "Sample Retailer",
        PurchaseDate: "2023-10-01",
        PurchaseTime: "12:00 PM",
        Items: []Item{
            {ShortDescription: "Hot Cheetos", Price: "5.00"},
            {ShortDescription: "Switch", Price: "10.00"},
        },
        Total: "30.00",
    }

    // Convert the receipt to JSON
    jsonData, err := json.Marshal(receipt)
    if err != nil {
        t.Fatalf("Failed to marshal receipt: %v", err)
    }

    // Create a new HTTP request
    req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a response recorder
    resp := httptest.NewRecorder()


	//																WHEN
    // Serve HTTP request
    router.ServeHTTP(resp, req)

	//																THEN
    // Check the response code
	assert.Equal(t, resp.Code, http.StatusCreated)

}



func TestAddReceipt_FAIL_NO_JSON(t *testing.T) {
	// 																GIVEN
    router := setupRouter()

    // Create a new HTTP request
    req, err := http.NewRequest("POST", "/receipts/process", nil)

    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a response recorder
    resp := httptest.NewRecorder()


	//																WHEN
    // Serve HTTP request
    router.ServeHTTP(resp, req)

	//																THEN
	assert.NotEqual(t, resp.Code, http.StatusCreated)

}


func TestAddReceipt_FAIL_BAD_JSON(t *testing.T) {
	// 																GIVEN
	router := setupRouter()
    // Create a sample receipt to add
    receipt := Receipt{
        Retailer:     "None Store",
        PurchaseDate: "2023-10-01",
        PurchaseTime: "12:00 PM",
        Total: "00.00",
    }
    // Convert the receipt to JSON
    jsonData, err := json.Marshal(receipt)
    if err != nil {
        t.Fatalf("Failed to marshal receipt: %v", err)
    }

    // Create a new HTTP request
    req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a response recorder
    resp := httptest.NewRecorder()


	//																WHEN
    // Serve HTTP request
    router.ServeHTTP(resp, req)

	//																THEN
    // Check the response code
	assert.NotEqual(t, resp.Code, http.StatusCreated)

}

// To be perfectly honest my ideal test for this kind of situation would be different.
// Since the given is the exact same between all 3 tests, I'd have made one giant test 
// with the GIVEN being reused, and made subtests for each of these cases since they are idempotent. 
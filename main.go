package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/google/uuid"
)

// Represents the items within a customer's Receipt.
type Item struct {
    ShortDescription string  `json:"shortDescription"`
    Price            string  `json:"price"`
}

// Purchase represents the total purchase transaction
type Receipt struct {
    Retailer     	string `json:"retailer"`
    PurchaseDate 	string `json:"purchaseDate"`
    PurchaseTime 	string `json:"purchaseTime"`
    Items        	[]Item `json:"items"`
    Total        	string `json:"total"`
}

var receiptsMap = make(map[string]Receipt)

func add_receipt(context *gin.Context){
	var newReceipt Receipt
	// Attempts to convert JSON within request to our todo type.
	if err := context.BindJSON(&newReceipt); err != nil{
		// This is where a descriptive error should go.
		return
	}
	// Convert to string otherwise incompatable with Golang's map.
	newID := uuid.New().String()
	receiptsMap[newID] = newReceipt
	context.IndentedJSON(http.StatusCreated, "id:" + newID)

}

func main(){
	router := gin.Default()
	router.POST("/receipts/process", add_receipt)
	router.Run("localhost:1313")
}

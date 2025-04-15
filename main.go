package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Represents the items within a customer's Receipt.
type Item struct {
    ShortDescription string `json:"shortDescription" validate:"required"`
    Price            string `json:"price" validate:"required"`
}

type Receipt struct {
    Retailer      string `json:"retailer" binding:"required"`
    PurchaseDate  string `json:"purchaseDate" binding:"required"`
    PurchaseTime  string `json:"purchaseTime" binding:"required"`
    Items         []Item `json:"items" binding:"required,dive"`
    Total         string `json:"total" binding:"required"`
}

// Create a Map to store our IDs to Receipts.  
var receiptsMap = make(map[string]Receipt)

func add_receipt(context *gin.Context){
	var newReceipt Receipt

	// Attempts to convert JSON within request to our todo type.
	if err := context.BindJSON(&newReceipt); err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Could not transform json into Receipt"})
		return
	}

	// Convert to string otherwise incompatable with Golang's map.
	newID := uuid.New().String()
	// Adds ID : Receipt to Map.
	receiptsMap[newID] = newReceipt
	// Returns {ID: theNewID} to the user
	context.IndentedJSON(http.StatusCreated, "id:" + newID)
}

func main(){
	router := gin.Default()
	router.POST("/receipts/process", add_receipt)
	router.Run("localhost:1313")
}

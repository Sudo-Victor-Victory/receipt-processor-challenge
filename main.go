package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
func init() {
    validate = validator.New()
}

// Represents the items within a customer's Receipt.
type Item struct {
    ShortDescription string `json:"shortDescription" validate:"required"`
    Price            string `json:"price" validate:"required"`
}

type Receipt struct {
    Retailer      string `json:"retailer" validate:"required"`
    PurchaseDate  string `json:"purchaseDate" validate:"required"`
    PurchaseTime  string `json:"purchaseTime" validate:"required"`
    Items         []Item `json:"items" validate:"required,dive"`
    Total         string `json:"total" validate:"required"`
}

var receiptsMap = make(map[string]Receipt)

func add_receipt(context *gin.Context){
	var newReceipt Receipt

	// Attempts to convert JSON within request to our todo type.
	if err := context.BindJSON(&newReceipt); err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Could not transform json into Receipt"})
		return
	}

    // Validate the receipt
	if err := validate.Struct(newReceipt); err != nil {
        context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Validation failed", "details": err.Error()})
        return
    }

	// Convert to string otherwise incompatable with Golang's map.
	newID := uuid.New().String()
	// Adds ID : Receipt to Map.
	receiptsMap[newID] = newReceipt
	context.IndentedJSON(http.StatusCreated, "id:" + newID)
}

func main(){
	router := gin.Default()
	router.POST("/receipts/process", add_receipt)
	router.Run("localhost:1313")
}

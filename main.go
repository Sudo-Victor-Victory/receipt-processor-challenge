package main
import (
	"net/http"
	"unicode"
	"strings"
 	"math"
	"time"
	"strconv"
	"errors"
	"fmt"

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
// Used to convert a string to date.
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


func get_receipt_by_id(id string) (*Receipt, error){
	var retrieved_Receipt Receipt
	retrieved_Receipt = receiptsMap[id]
	// If the receipt doesn't exist, retrieved_Receipt will be zero valued
	if retrieved_Receipt.Retailer == ""{
		return nil, errors.New("Receipt with listed ID not found")
	}
	return &retrieved_Receipt, nil
}

func process_id(receipt *Receipt) (int, error){
	layout := "2006-01-02" 
	var points int = 0
	// One point for every alphanumeric character in Retailer
	for _, r := range receipt.Retailer{
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			points++
		}
	}
	fmt.Println("Points after Retailer:", points)


	// Checks if the total is a multiple of 0.25
	f, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
	   fmt.Println("Error:", err)
	   return 0, err
	}
	if math.Mod(f , 0.25) == 0 {
		fmt.Println("Points after Total being a multiple of 0.25:", points)
		points+=25
	}


	// points every 5 points for evey 2 items
	var  number_of_items int = len(receipt.Items) / 2
	points +=  number_of_items * 5
	fmt.Println("Points after # of items:", points)


	// For every item description, if the len of description % 3 == 0 multiply by 0.2, round, and add to points.
	for _, r := range receipt.Items {
		if len(strings.TrimSpace(r.ShortDescription)) % 3 ==0 {
			price, _ := strconv.ParseFloat(r.Price, 64)
			points += int(math.Ceil(( float64(price) * 0.2)))
		}
	}
	fmt.Println("Points after Item Descriptions :", points)
	// Check if purchase date is odd
	parsedDate, err := time.Parse(layout, receipt.PurchaseDate)
	if err != nil {
		fmt.Println("Error parsing time string:", err)
		return 0, errors.New("Time string is formatted incorrectly")
	}

	if parsedDate.Day() % 2 != 0 {
		points += 6
		fmt.Println("Points after odd day:", points)

	}
	// Check if purchase time is between 2 and 4 PM.
	start:= "14:00"
	end :=  "16:00"
	if start < receipt.PurchaseTime && receipt.PurchaseTime < end {
		fmt.Println("Points if between 2 and 4", points)
		points += 10
	}
	return points, nil
}

func get_receipt_handler(context *gin.Context){
	id := context.Param("id")
	receipt, err := get_receipt_by_id(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Receipt not found"})
		return
	}
	// Process id
	points, err := process_id(receipt)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in processing receipt field in points calculation"})
	}
	return_string := strconv.Itoa(points)
	context.IndentedJSON(http.StatusOK, "points:" + return_string);
}
func main(){
	router := gin.Default()
	router.POST("/receipts/process", add_receipt)
	router.GET("/receipts/:id/points", get_receipt_handler)
	router.Run("localhost:1313")
}

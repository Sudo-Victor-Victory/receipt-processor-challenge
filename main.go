package main
import (
	"net/http"
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
	var points int = 0
	var temp_points int = 0
	// One point for every alphanumeric character in Retailer
	points += points_retailer(receipt.Retailer)
	fmt.Println("Points after Retailer:", points)

	// Checks if the total is a round-dollar amount
	temp_points, err := points_round_dollar_amount(receipt.Total)
	if err != nil {
		fmt.Println("Could not convert Receipt total to float during points_round_dollar_amount check: ", err)
		return 0, err
	}
	points += temp_points
	fmt.Println("Points after round Receipt Total check:", points)


	// Checks if the total is a multiple of 0.25
	temp_points, err = points_multiple_of_25(receipt.Total)
	if err != nil {
		fmt.Println("Could not convert Receipt total to float during points_multiple check: ", err)
		return 0, err
	}
	points += temp_points
	fmt.Println("Points after multiple of 0.25 check:", points)

	// points every 5 points for evey 2 items
	points += points_per_2_items(len(receipt.Items))
	fmt.Println("Points after # of items:", points)

	// For every item description, if the len of description % 3 == 0 multiply by 0.2, round, and add to points.
	points += points_per_item_description(receipt.Items)
	fmt.Println("Points after Item Descriptions :", points)

	// Check if purchase date is odd
	temp_points, err = points_odd_day(receipt.PurchaseDate)
	if err != nil {
		fmt.Println("Could not get points for odd day: ", err)
	}
	points += temp_points
	fmt.Println("Points after odd day:", points)

	// Check if purchase time is between 2 and 4 PM.
	points += points_between_time(receipt.PurchaseTime)
	fmt.Println("Points if between 2 and 4", points)
	return points, nil
}

func get_receipt_points_handler(context *gin.Context){
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
	router.GET("/receipts/:id/points", get_receipt_points_handler)
	router.Run("localhost:1313")
}

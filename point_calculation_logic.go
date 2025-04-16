package main


import (
	"unicode"
	"strings"
 	"math"
	"time"
	"strconv"
	"errors"
	"fmt"
)

// One point for every alphanumeric character in Retailer
func points_retailer(retailer string) int  {
	var points int = 0
	for _, r := range retailer{
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			points++
		}
	}
	return points
}

// Checks if the Receipt total is a round dollar
func points_round_dollar_amount(pointTotal string) (int, error) {
	points_as_float, err := strconv.ParseFloat(pointTotal, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	// Math.mod is required because base operator % only works for Ints.
	if math.Mod(points_as_float , 1.0) == 0 {
		//fmt.Println("Points after Total being a multiple of 0.25:", points)
		return 50, nil
	}
	return 0, nil
}

// Checks if the Receipt total is a multiple of 0.25
func points_multiple_of_25(pointTotal string) (int, error) {
	points_as_float, err := strconv.ParseFloat(pointTotal, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	// Math.mod is required because base operator % only works for Ints.
	if math.Mod(points_as_float , 0.25) == 0 {
		//fmt.Println("Points after Total being a multiple of 0.25:", points)
		return 25, nil
	}
	return 0, nil
}

// For every 2 items the user gets 5 points.
func points_per_2_items(item_total int) int {
	var  number_of_item_duos int = item_total / 2
	number_of_item_duos *= 5
	return number_of_item_duos
}

// For every item description, if the len of description % 3 == 0 multiply by 0.2, round, and add to points.
func points_per_item_description(items []Item) int {
	var temp_points int = 0
	for _, r := range items{
		if len(strings.TrimSpace(r.ShortDescription)) % 3 ==0 {
			price, _ := strconv.ParseFloat(r.Price, 64)
			temp_points += int(math.Ceil(( float64(price) * 0.2)))
		}
	}
	return temp_points
}

// Check if purchase date is odd
func points_odd_day(purchaseDate string) (int, error){
	layout := "2006-01-02" 
	// Check if purchase date is odd
	parsedDate, err := time.Parse(layout, purchaseDate)
	if err != nil {
		fmt.Println("Error parsing time string:", err)
		return 0, errors.New("Time string is formatted incorrectly")
	}

	if parsedDate.Day() % 2 != 0 {
		return 6, nil
	}
	return 0, nil
}

// Check if purchase time is between 2 and 4 PM.
func points_between_time(purchaseTime string) int { 
	// Check if purchase time is between 2 and 4 PM.
	start:= "14:00"
	end :=  "16:00"
	if start < purchaseTime && purchaseTime < end {
		return 10
	}
	return 0 
}

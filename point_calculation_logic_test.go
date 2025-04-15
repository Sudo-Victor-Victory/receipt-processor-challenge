package main

import (
    "testing"

	"github.com/stretchr/testify/assert"
)

func Test_Point_Retailer_Empty_0_Points(t *testing.T) {
	// 																GIVEN
	retailer := ""
	//																WHEN
	temp_points := points_retailer(retailer)
	//																THEN
	assert.Equal(t, temp_points, 0)
}

func Test_Point_Retailer_Non_Empty(t *testing.T) {
	// 																GIVEN
	empty := "Target"
	//																WHEN
	temp_points := points_retailer(empty)
	//																THEN
	assert.Equal(t, temp_points, 6)
}

func Test_Receipt_Total_Multiple_Of_25_Failure(t *testing.T){
	// 																GIVEN
	receipt_total := "2.79"
	//																WHEN
	temp_points,_ := points_multiple_of_25(receipt_total)
	//																THEN
	assert.Equal(t, temp_points, 0)
}

func Test_Receipt_Total_Multiple_Of_25_SUCCESS(t *testing.T){
	// 																GIVEN
	receipt_total := "2.75"
	//																WHEN
	temp_points,_ := points_multiple_of_25(receipt_total)
	//																THEN
	assert.Equal(t, temp_points, 25)
}


func Test_Receipt_Points_Per_2_Items_EVEN(t *testing.T){
	// 																GIVEN
	total_num_of_items := 4 
	//																WHEN
	temp_points := points_per_2_items(total_num_of_items)
	//																THEN
	assert.Equal(t, temp_points, 10)
}

func Test_Receipt_Points_Per_2_Items_ODD(t *testing.T){
	// 																GIVEN
	total_num_of_items := 3 
	//																WHEN
	temp_points := points_per_2_items(total_num_of_items)
	//																THEN
	assert.Equal(t, temp_points, 5)
}

func Test_Receipt_Points_Per_Item_Description(t *testing.T){
	// 																GIVEN
	items := []Item{    
        Item{    
			ShortDescription: "Mountain Dew 12PK",
			Price : "6.49",
        },
        Item{    
			ShortDescription: "Emils Cheese Pizza",
			Price : "12.25",
        },
		Item{    
			ShortDescription: "Knorr Creamy Chicken",
			Price : "1.26",
        },
		Item{    
			ShortDescription: "Doritos Nacho Cheese",
			Price : "3.35",
        },
		Item{    
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price : "12.00",
        },
    }
	//																WHEN
	temp_points := points_per_item_description(items)
	//																THEN
	assert.Equal(t, temp_points, 6)
}

func Test_Receipt_Points_ODD_DAY_ODD(t *testing.T){
	// 																GIVEN
	day := "2006-01-03"  
	//																WHEN
	temp_points, _ := points_odd_day(day)
	//																THEN
	assert.Equal(t, temp_points, 6)
}

func Test_Receipt_Points_ODD_DAY_EVEN(t *testing.T){
	// 																GIVEN
	day := "2006-01-02"  
	//																WHEN
	temp_points, _ := points_odd_day(day)
	//																THEN
	assert.Equal(t, temp_points, 0)
}



func Test_Receipt_Points_BETWEEN_TIME_SUCCESS(t *testing.T){
	// 																GIVEN
	purchaseTime := "14:45"  
	//																WHEN
	temp_points := points_between_time(purchaseTime)
	//																THEN
	assert.Equal(t, temp_points, 10)
}

func Test_Receipt_Points_BETWEEN_TIME_FAILURE(t *testing.T){
	// 																GIVEN
	purchaseTime := "19:45"  
	//																WHEN
	temp_points := points_between_time(purchaseTime)
	//																THEN
	assert.Equal(t, temp_points, 0)
}

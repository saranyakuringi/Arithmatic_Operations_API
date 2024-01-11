package operations

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Operations_Type struct {
	Num []int `json:"num"`
}

func Average(num []int) int {
	//sum/total
	//num := []int{1, 2, 3}
	var data int
	for i := 0; i < len(num); i++ {
		data = num[i] + data
	}
	Avg := data / len(num)
	fmt.Println("The Average is:", Avg)
	return Avg
}

func GET_Average(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request Body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request Body"})
		return
	}
	var data Operations_Type
	err = json.Unmarshal(request, &data)
	if err != nil {
		log.Println("Error in Unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in Unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Average": Average(data.Num)})
}

func Min(num []int) int {
	var min int = num[0]

	for _, value := range num {
		if min > value {
			min = value
		}
	}
	return min
}

func Max(num []int) int {
	var max int = num[0]
	for _, value := range num {
		if max < value {
			max = value
		}
	}
	return max
}
func GET_MinMax(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	var data Operations_Type
	err = json.Unmarshal(request, &data)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Mimimum number": Min(data.Num), "Maximum number": Max(data.Num)})
}

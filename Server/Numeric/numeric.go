package numeric

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Numeric_type struct {
	Num []int `json:"num"`
}

func Even(num []int) []int {
	var data []int
	for _, value := range num {
		if value%2 == 0 {
			data = append(data, value)
		}
	}
	return data
}

func Odd(num []int) []int {
	var data []int
	for _, value := range num {
		if value%2 != 0 {
			data = append(data, value)
		}
	}
	return data
}

func Square(num []int) []int {
	var data []int
	for _, value := range num {
		square := value * value
		data = append(data, square)
	}
	return data
}

func GET_EvenOdd(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	var data Numeric_type
	err = json.Unmarshal(request, &data)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Even numbers": Even(data.Num), "Odd Numbers": Odd(data.Num), "Square of the numbers": Square(data.Num)})
}

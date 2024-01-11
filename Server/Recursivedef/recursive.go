package recursivedef

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recursive struct {
	Num int `json:"num"`
}

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}

func Fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
func GET_Recursive(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}

	var data Recursive
	err = json.Unmarshal(request, &data)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Factorial": Factorial(data.Num), "Fibonacci": Fibonacci(data.Num)})
}

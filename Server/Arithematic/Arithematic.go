package Arithematic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Arithematic_Type struct {
	Num1 []int `json:"num1"`
	Num2 []int `json:"num2"`
}

func Addition(num1 []int, num2 []int) []int {
	// num1 = []int{1, 2, 3}
	// num2 = []int{4, 5, 6}
	num := make([]int, len(num1))
	// fmt.Println(len(num1))
	// fmt.Println(len(num2))

	for i := 0; i < len(num1); i++ {
		num[i] = num1[i] + num2[i]
	}
	fmt.Println("The Sum is:", num)
	return num
}

func GET_Addition(c *gin.Context) {
	//read the data
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(404, gin.H{"message": "Error in request body"})
		return
	}
	//unmarshal the data
	var Add Arithematic_Type
	err = json.Unmarshal(request, &Add)
	if err != nil {
		log.Println("Error in unmarshal data", err)
		c.IndentedJSON(404, gin.H{"message": "Error in unmarshal data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"The Addition is :": Addition(Add.Num1, Add.Num2)})
}

func Subtraction(num1 []int, num2 []int) []int {
	num := make([]int, len(num1))
	for i := 0; i < len(num1); i++ {
		num[i] = num1[i] - num2[i]
	}
	fmt.Println("The Subtraction is:", num)
	return num
}

func GET_Subtraction(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	var data Arithematic_Type
	err = json.Unmarshal(request, &data)
	if err != nil {
		log.Println("Error in Unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in Unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": Subtraction(data.Num1, data.Num2)})
}

func Multiplication(num1 []int, num2 []int) []int {
	// num1 = []int{1, 2, 3}
	// num2 = []int{4, 5, 6}
	num := make([]int, len(num1))
	// fmt.Println(len(num1))
	// fmt.Println(len(num2))

	for i := 0; i < len(num1); i++ {
		num[i] = num1[i] * num2[i]
	}
	fmt.Println("The Multiplication is:", num)
	return num
}

func GET_Multiplication(c *gin.Context) {
	//read the data
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(404, gin.H{"message": "Error in request body"})
		return
	}
	//unmarshal the data
	var Add Arithematic_Type
	err = json.Unmarshal(request, &Add)
	if err != nil {
		log.Println("Error in unmarshal data", err)
		c.IndentedJSON(404, gin.H{"message": "Error in unmarshal data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"The Multiplication is :": Multiplication(Add.Num1, Add.Num2)})
}

func Division(num1 []int, num2 []int) []int {
	// num1 = []int{1, 2, 3}
	// num2 = []int{4, 5, 6}
	num := make([]int, len(num1))
	// fmt.Println(len(num1))
	// fmt.Println(len(num2))

	for i := 0; i < len(num1); i++ {
		num[i] = num1[i] / num2[i]
	}
	fmt.Println("The Division is:", num)
	return num
}

func GET_Division(c *gin.Context) {
	//read the data
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(404, gin.H{"message": "Error in request body"})
		return
	}
	//unmarshal the data
	var Add Arithematic_Type
	err = json.Unmarshal(request, &Add)
	if err != nil {
		log.Println("Error in unmarshal data", err)
		c.IndentedJSON(404, gin.H{"message": "Error in unmarshal data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"The Division is :": Division(Add.Num1, Add.Num2)})
}

func Modulo(num1 []int, num2 []int) []int {
	// num1 = []int{1, 2, 3}
	// num2 = []int{4, 5, 6}
	num := make([]int, len(num1))
	// fmt.Println(len(num1))
	// fmt.Println(len(num2))

	for i := 0; i < len(num1); i++ {
		num[i] = num1[i] % num2[i]
	}
	fmt.Println("The Modulo is:", num)
	return num
}

func GET_Modulo(c *gin.Context) {
	//read the data
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(404, gin.H{"message": "Error in request body"})
		return
	}
	//unmarshal the data
	var Add Arithematic_Type
	err = json.Unmarshal(request, &Add)
	if err != nil {
		log.Println("Error in unmarshal data", err)
		c.IndentedJSON(404, gin.H{"message": "Error in unmarshal data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"The Modulo is :": Modulo(Add.Num1, Add.Num2)})
}

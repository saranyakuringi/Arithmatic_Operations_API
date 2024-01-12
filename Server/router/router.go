package router

import (
	"Basic_ArithematicOperations/Server/Arithematic"
	numeric "Basic_ArithematicOperations/Server/Numeric"
	operations "Basic_ArithematicOperations/Server/Operations"
	recursivedef "Basic_ArithematicOperations/Server/Recursivedef"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/Addition", Arithematic.GET_Addition)
	router.GET("/Subtraction", Arithematic.GET_Subtraction)
	router.GET("/Multiplication", Arithematic.GET_Multiplication)
	router.GET("/Division", Arithematic.GET_Division)
	router.GET("/Modulo", Arithematic.GET_Modulo)
	router.GET("/Average", operations.GET_Average)
	router.GET("/MinMax", operations.GET_MinMax)
	router.GET("/Recursive", recursivedef.GET_Recursive)
	router.GET("/EvenOdd", numeric.GET_EvenOdd)
	router.Run("localhost:8089")
}

package router

import (
	"Basic_ArithematicOperations/Server/Arithematic"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/Addition", Arithematic.GET_Addition)
	router.GET("/Subtraction", Arithematic.GET_Subtraction)
	router.GET("/Multiplication", Arithematic.GET_Multiplication)
	router.GET("/Division", Arithematic.GET_Division)
	router.GET("/Modulo", Arithematic.GET_Modulo)
	router.Run("localhost:8089")
}

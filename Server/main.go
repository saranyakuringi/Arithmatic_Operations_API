package main

import (
	"Basic_ArithematicOperations/Server/router"
	"fmt"
)

func main() {
	fmt.Print("This is an API for basic operations like:\n", "1. Addition\n",
		"2. Subtraction\n", "3. Multiplication\n", "4. Division\n",
		"5. Modulo\n", "6. Average\n", "7. Mimimum and Maximum\n")
	router.Router()

}

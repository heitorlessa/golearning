// big flat file containing language basics 
// in the form of code to speed learning process

	// :: Import libraries

import (

	// no need for ',' 
	"fmt" 
	"math/rand"
	"time"
	"os"
	"github.com/gin-gonic/gin"
)

	// :: 3rd-party libraries
	// ::: go get github.com/gin-gonic/gin
	
	// :: Variables

var beingDeclared string
var integer int
var pointer *int
var float float64

shorter := "myString"

	// :: Arrays

stringArray := [2]string{"First", "Second"}

	// :: Slices

slice := []{"First", "Second", "Third"}
slicing = =slice[2:]
bytes := make([]byte, 5) // []byte == byte type and 5 length

	// :: Map

	// :: Functions

func anonFunction() {
	fmt.Println("Anonymous function ")
}

declaredAnonFunction := func() {
	fmt.Println("Anon function saved into a variable otherwise it becomes useless")
}

func namedFunction() string {
	fmt.Println("Normal function without parameters that has to return a String")
}

func namedFunctionWithParameters(message string) string {
	fmt.Println("Message to be printed -> ", message)
}



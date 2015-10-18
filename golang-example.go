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

// binary Function
type binaryFunction func(int, int) int

// slice of functions
fns := []binaryFunction{
	func (x, y int) int { return x + y },
	func (x, y int) int { return x - y }
}

	// seed
	rand.Seed(time.Now().Unix())

	// pick one of those functions at random
	f := fns[rand.Intn(len(fns))]

// Functions as Fields 
type op struct {
    name string
    fn   func(int, int) int
}

func main() {
    // seed your random number generator
    rand.Seed(time.Now().Unix())

    // create a slice of ops
    ops := []op{
        {"add", func(x, y int) int { return x + y }},
        {"sub", func(x, y int) int { return x - y }},
        {"mul", func(x, y int) int { return x * y }},
        {"div", func(x, y int) int { return x / y }},
        {"mod", func(x, y int) int { return x % y }},
    }

    // pick one of those ops at random
    o := ops[rand.Intn(len(ops))]

    x, y := 12, 5
    fmt.Println(o.name, x, y)
    fmt.Println(o.fn(x, y))
}

	// :: Struct

type Circle struct {
	x, y, z float64
}

// Init
var c Circle
d := new(Circle)
test := Circle(0, 0, 5)

// Accessing "Fields"
test.x
test.y
test.y

func circleArea(c *Circle) float64 {
  return math.Pi * c.r*c.r
}

c := Circle{0, 0, 5}
fmt.Println(circleArea(c))

// Creating a method instead
// func (var *Struct) functionName returnType
func (c *Circle) area() float64{
	return math.Pi * c.r*c.r
}

// other struct example
type Name struct {
    First  string
    Middle string
    Last   string
}

// "method" declaration 2
func (n Name) String() string {
    return fmt.Sprintf("%s %c. %s", n.First, n.Middle[0], strings.ToUpper(n.Last))
}

// Constructing it
n := Name{"William", "Mike", "Smith"}
    fmt.Printf("%s", n.String())

// Finding a type (reflection) with reflect package

fmt.Println(reflect.TypeOf(resp))


// Convert struct that is usually a response from an API call (aws-sdk-go) to []byte so it can be written in a file

svc := ec2.New(&aws.Config{Region: aws.String("eu-west-1")})

resp, err := svc.DescribeInstances(nil)
if err != nil {
	panic(err)
}

if err := ioutil.WriteFile("/tmp/ec2Response", []byte(fmt.Sprintf("%v", *resp)), 0644); err != nil {
	fmt.Errorf("Error happened here: ", err)
}

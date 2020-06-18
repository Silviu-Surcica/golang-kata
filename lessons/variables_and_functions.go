/*
Packages
Every Go program is made up of packages.
Programs start running in package main.
*/
package lessons

/* Grouping the imports into a parenthesized, "factored" import statement.
You can also write multiple import statements, like:
import "fmt"
import "math/rand"
But it is good style to use the factored import statement.
*/
import (
	"fmt"
	"math/rand"
)

/*
Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

/*
A function can take zero or more arguments.
In this example, swap takes two parameters of type string.
Because x and y have the same type this can be shortened to:
func swap(x, y string)

Notice that the type comes after the variable name and also you need to specify the returned types.
*/
func swap(x string, y string) (string, string) {
	return y, x
}
/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function.
A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions, as with the example shown here.
They can harm readability in longer functions.
*/
func add(a int, b int) (sum int) {
    sum = a + b
    return
}

/* The var statement declares a list of variables; as in function argument lists, the type is last.
Variables declared without an explicit initial value are given their zero value.
The zero value is:
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/
var i, j int
// Initialized variables
var inter int = 2
var boolean bool = true
/*
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
k := 3 // the type is inferred
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

//Type conversion. The expression T(v) converts the value v to the type T
var f float64 = float64(inter)

/*
Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.
Numeric constants are high-precision values.
*/
const Truth = true

func print_rand(){
    fmt.Println("My favorite number is", rand.Intn(100))
}

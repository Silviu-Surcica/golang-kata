package lessons
import "fmt"

func sum(a int) int {
    sum := 0
    /*
    For loop is the only looping construct in Go.
    There are no parenthesis, only semicolons which separates each component.
    The init and post statements are optional:
    for i:=0; i<a ; {}
    for ; i<a; {} // which can be simply put as: for i<a {} and the for acts as a while statement in C
    If you omit the loop condition, it will run forever: for {}
    */
    for i := 0; i < a; i++ {
        /*
        The if statement can start with a short statement to execute before the condition.
        Variables declared this way are in the scope of if or else only:
        if c:= math.Rand(19); a > 5 {}
        */
        if i % 2 == 0 {
            sum += i
        } else {
            sum += 0
        }
    }
    return sum
}

/*
Switch statement is like in C but it doesnt run all the cases until it reaches break.
Break is by default at the end of each case.
In essence the switch can be used as a if elif else
Cases are evaluated in order from top to bottom, stopping when a case succeeds.
Case values can be anything but constants.
The switch without a condition can be a clean way to write if then else chains:
switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
*/
func SelectOption(a int) {
    switch a {
    case 0:
        fmt.Println("Selected value is 0")
    case 1:
        fmt.Println("Selected value is 1")
    default:
        fmt.Println("Neither 0 or 1")
    }
}

/*
The defer statement will push the function call into a stack, executing the function
in a LIFO style after the surrounding function has returned.
*/
func DeferCount(a int) {
    fmt.Println("Start")
    for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
	// this will print Start Done 9 8 7 6 5 4 3 2 1 0
}

/* ==========EXERCISE=============
Write a function called Pow that will receive two parameters of type int.
You need to implement the power math function.
The first parameter is the base and the second is the exponential power.
The function should return a float64 as a result.
*/

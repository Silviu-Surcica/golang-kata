package lessons


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

/* ==========EXERCISE=============
Write a function called Pow that will receive two parameters of type int.
You need to implement the power math function.
The first parameter is the base and the second is the exponential power.
The function should return a float64 as a result.
*/
func Pow(a int, b int) float64 {
    p := 1.0
    if a == 1 || b == 0{
        return 1
    }
    if a == 0 {
        return 0
    }
    if b > 0 {
        for i:=1; i<b+1;i++ {
            p *= float64(a)
        }
    } else{
        for i := b; i<0; i ++{
                p /= float64(a)
        }
    }
    return p
}
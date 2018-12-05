package main
import "fmt"
func main() {
var intVar int;
var strVar string;
var intArrayVar [10] int;
var structVar struct{
	f int
}
var intPointerVar *int
var funcVar func(a int) int
intVar = 100
strVar = "asdasd"
intArrayVar = [10] int{1,2,3,4,5,6,7,8,9,10}
intPointerVar = &intVar
structVar.f = 11
funcVar = f;
fmt.Println(intVar ,strVar ,intPointerVar ,intArrayVar ,structVar ,funcVar,funcVar(4),"\n");
}
func f(a int) int{
	return a
}
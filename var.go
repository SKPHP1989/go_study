package main
import "fmt"
func main() {
	var myArray [10] int = [10] int{1,2,3,4,5,6,7,8,9,10}
	var mySlice [] int = myArray[:5]
	fmt.Println("Elements of myArray:")
	for _,v := range myArray{
		fmt.Print(v," ")
	}
	fmt.Println("\nElements of mySlice:")
	for _, v := range mySlice{
		fmt.Print(v," ")	
	}
	fmt.Println("\nElements of  myArray[5:]:")
	for _, v := range myArray[5:] {
		fmt.Print(v," ")	
	}
	fmt.Println();
	str := "Hello ,世界"
	for i,v := range str {
		fmt.Println(i,"=>",v)
	}
	arrLen := len(myArray)
	fmt.Println("arrLen" ,arrLen);

	array := [5] int {1,2,3,4,5}
	modify(array);
	fmt.Println("In modify(),array values",array)
}
func modify(array [5] int){
	array[0]=10;
	fmt.Println("In modify(),array values",array)
}
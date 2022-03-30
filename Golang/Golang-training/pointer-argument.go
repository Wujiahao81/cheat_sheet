package main
import "fmt"
func add(xPtr *int){
	*xPtr=*xPtr+10 //對當初的a做運算
	//fmt.Println("Add function:",*xPtr)
}
func main(){
	var a int=10
	// var aPtr *int=&a
	// add(aPtr) // Pass by pointer
	add(&a) //Pass byt pointer
	fmt.Println("Main function:",a)

	//和使用者要求輸入，運用到指標參數 Pass by pointer
	var msg string
	var msgPtr *string=&msg
	fmt.Scanln(msgPtr)
	// fmt.Scanln(&msg) //另一種方法，兩種皆可，看情況使用
	fmt.Println("result:",msg)
}
/*
func add(x int){
	x=x+10
	fmt.Println("Add function:",x)
}
func main(){
	var a int=10
	add(a) // Pass by value
	fmt.Println("Main function:",a)
}
*/
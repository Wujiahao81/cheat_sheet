package main
import "fmt"
// func 函式名稱(參數列表)回傳值資料型態{
// 	函式內部的程式碼
// 	return 回傳值，須符合定義中的資料型態
// }
//回傳值(return)就是把資料從函式的裡面帶到呼叫函式的地方，也就是函式的外面的一個動作

func show(msg string){
	if msg=="Hello"{
		return
	}
	fmt.Println(msg)
}
func add(n1 int,n2 int) int{
	var result int=n1+n2
	// fmt.Println(result)
	return result
}
func test() (int,string){
	return 6,"Test!!"
}
func main(){
	//基本的 return 運作方式
	/*
	show("Hello")
	show("測試")
	*/
	/*
	//單一回傳值
	var x int=add(333,333) //回傳到函數裡面加在一起
	fmt.Println(x*999) //印出回傳值
	*/
	//多個回傳值
	var x int
	var s string
	x,s=test()
	fmt.Println(x,s)
}
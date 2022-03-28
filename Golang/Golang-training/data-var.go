//  資料與變數
package main
import "fmt"
func main(){
	/*
	fmt.Println(4) //int
	fmt.Println(3.1415) //float64
	fmt.Println("測試中文") //string
	fmt.Println(true) //bool
	fmt.Println('a') //字元 rune ，在程式中表達單一字符，會印出97，必須使用單引號
	*/
	//變數的使用
	var x int //宣告變數
	x=4 //指定資料:把資料 4 放進變數
	fmt.Println(x) //使用變數，用變數的名稱代替資料，做操作
	x=10 //指定新的資料:把變數10放進變數，會覆蓋原來的資料
	fmt.Println(x)
	// x="Hello" 資料型態不正確
	var f float64=3.1415926535 //宣告變數順便放進資料: var 變數名稱 資料型態=適當的資料
	fmt.Println(f)
	var s string="哈囉這是字串"
	fmt.Println(s)
	var test bool=true
	fmt.Println(test)
	var r rune='b' //印出98
	fmt.Println(r)
}
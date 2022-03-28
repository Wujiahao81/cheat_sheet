package main
import "fmt"
// 可以印出任何訊息的函式
// 參數名稱 資料型態
func show(msg string){
	fmt.Println(msg)
}
func add(n1 int,n2 int){
	var result=n1+n2
	fmt.Println(result)
}
func sum(max int){
	var result int=0
	var n int
	for n=1;n<=max;n++{
		result+=n
	}
	fmt.Println(result)
}
// 呼叫 show 函式
func main(){
	//基礎函式語法演練
	/*
	show("Hello 你好!!")
	add(123,456) //加法,印出579
	*/
	//計算 1+2+3+...+max的函式包裝
	var clint int
	for{
		fmt.Println("請輸入數字，將自動幫您加總1+2+3+...+?,按666即可跳出")
		fmt.Scanln(&clint)
		if clint==666{
			break
		}
		sum(clint)
	}
}
//若想改成按esc或其他鍵跳出，可參考這個網址 : https://stackoverflow.com/questions/40159137/golang-reading-from-stdin-how-to-detect-special-keys-enter-backspace-etc
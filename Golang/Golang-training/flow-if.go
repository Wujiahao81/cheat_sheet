// 流程控制-if判斷式
package main
import "fmt"
func main(){
	//基本語法練習
	/*
	if false{
		fmt.Println("Go")
	}else{
		fmt.Println("Not go")
	}
	*/
	//簡易情境:ATM 檢測使用者輸入的金額是否恰當
	var money int
	fmt.Println("請問想領多少錢? ")
	fmt.Scanln(&money)
	if money < 100{
		fmt.Println("Too Few")
	}else if money<=100000{
		fmt.Println("OK")
	}else{
		fmt.Println("Too Much")
	}
	fmt.Println("執行完畢")
}
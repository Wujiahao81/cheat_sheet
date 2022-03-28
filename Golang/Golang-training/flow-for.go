// 流程控制 - for 迴圈 (golang 沒有 while)
package main 
import "fmt"
func main(){
	//基本迴圈使用
	/* 
	//無窮迴圈，Golang有提供更簡單的寫法，在for右方不給定任何參數，就會是一個無窮迴圈
	for true{
		fmt.Println("Hello")
	}
	*/
	/*
	var x int=0
	for x<5{
		fmt.Println(x)
		x++
	}
	*/
	/*
	var x int
	for x=0;x<11;x+=2{
		fmt.Println(x)
	}
	*/
	//實際範例，計算 1+2+3+...+50的結果
	/*
	var result int=0
	var x int=1
	for x<=10 { //測試1+2+3+...+10
		fmt.Println("result",result,",x:",x)
		result=result+x
		x++
	}
	fmt.Println(result)
	*/
	// break 強制跳出迴圈
	// continue 強迫回到開頭
	/* //這裡會印出1,3,5
	var x int
	for x=0;x<6;x++{
		if x%2==0{ // x 是偶數 
			continue
		}
		fmt.Println(x)
	}
	*/
	// 實際範例，持續讓使用者輸入數字，計算總合。直到使用者輸入0為止
	var result int=0
	for true{
		var n int
		fmt.Scanln(&n)
		if n==0{
			break
		}
		result+=n //result=result=n
	}
	fmt.Println("總和:",result)
}
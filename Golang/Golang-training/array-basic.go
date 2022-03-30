// array陣列，是可以讓我們按照順序存放多個相同型態資料的容器
// struct結構，沒有順序性，裡面的資料可以是各種的資料型態
package main 
import "fmt"
// 陣列語法 : [長度]資料型態 
// ex. ver arr [4]int 是一個長度為四的整數陣列
func main(){
	/*
	//整數陣列
	var numbers [3]int
	numbers[0]=0
	numbers[1]=1
	numbers[2]=2
	fmt.Println(numbers) // [0 1 2]
	fmt.Println(numbers[2]*100) // 200
	var arr1 [6]int=[6]int{0,1,2,3,4,5}
	fmt.Println(arr1) //[0 1 2 3 4 5]
	// 字串陣列
	var names [2]string=[2]string{"Daoxue","Justin"}
	fmt.Println(names) //[Daoxue Justin]
	//取得陣列長度
	fmt.Println(len(numbers)) //3
	fmt.Println(len(arr1))	  //6
	fmt.Println(len(names))	  //2
	*/
	//巡迴(逐一取得)陣列中的資料
	// var grades [3]int=[3]int{60,90,75}
	var grades[4]int
	var sum int=0
	var index int //代表陣列編號
	//逐一輸入陣列的資料
	fmt.Println("請逐一輸入資料，會自動幫您計算平均")
	for index=0; index<len(grades);index++{
		fmt.Scanln(&grades[index])
	}
	for index=0; index<len(grades);index++{
		// fmt.Println(grades[index]) // 60 下一行 90 下一行 75
		sum = sum+grades[index]
	}
	fmt.Println(sum/len(grades)) //平均數 
}
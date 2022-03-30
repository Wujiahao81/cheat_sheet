package main
import "fmt"
/*
type 結構名稱 struct{
	欄位名稱 資料型態
	欄位名稱 資料型態
}
*/
// 結構實體化 : 實體化結構就是給定結構的資料欄位資料
type Point struct{
	x int
	y int
}
type Human struct{
	name string
	age int
	gender string
}
func main(){
	var p1 Human=Human{"岳飛",100,"male"} // 結構名稱{欄位資料,欄位資料,...}
	// var p2 Point=Point{y:2,x:1} // 結構名稱{欄位名稱:資料,欄位名稱:資料,...}
	var p3 Human=Human{name:"Daoxue",gender:"male",age:20}
	// fmt.Println(p2.x,p2.y) //使用 . 來存取欄位中的資料
	p1.name="小明" //改名字
	fmt.Println(p1.name,p1.gender,p1.age) 
	fmt.Println(p3.name,p3.age,p3.gender)
}
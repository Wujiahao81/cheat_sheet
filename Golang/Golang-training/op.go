//基本運算
package main
import "fmt"
func main(){
	//算術運算: +,-,*,/
	var x int
	x=3*3+41
	fmt.Println(x)
	//指定運算: =,+=,-=,*=,/=
	x=5
	x+=3 //x=x+3 可以簡寫成前面的形式
	fmt.Println(x)
	//單元運算: ++,--
	x++ //加一
	fmt.Println(x)
	//比較運算: >,<,>=,<=,==
	var result bool=4==3 //問電腦4跟3是否有相等
	fmt.Println(result)
	//邏輯運算: !, &&, ||
	// result=!ㄋfalse //會印出true
	result=false&&false //and, 且, &&: 兩邊如果都是 true，結果就是 true
	fmt.Println(result)
	result=true||false //or , 或, ||: 只要有一個 true，結果就是true
	fmt.Println(result)
	
}

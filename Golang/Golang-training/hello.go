package main //可執行程式必須使用 main 封包
import "fmt" //仔入內鍵的 fmt 封包，用來做基本輸出輸入
func main(){ //建立 main 函式，程式的進入點
	//執行程式時從這邊開始
	//輸出訊息到終端 : fmt.Println(資料,資料,...)
	fmt.Println("Hello golang")
}

//撰寫程式 > 建置(build) >執行程式
//建置: go build 程式檔案的名稱 ex: go build hello.go
//執行程式: 直接輸入執行檔的檔名或是 ./檔名
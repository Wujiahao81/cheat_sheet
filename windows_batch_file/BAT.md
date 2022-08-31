# Windows Batch File

>## 練習的程式
- [刪除在該路徑內所有超過7天的log檔](deloldlogs.bat)
- [刪除在該路徑內超過一天的所有jpg檔](deloldphoto.bat)
- [自動關閉並重新啟動程式](auto_restart_your_program.bat)
- [檢查某程式是否在執行，若沒有回應就刪除並重新啟動他](checkapp.bat)。 
可以用於一些必須一直執行的程式，[invis.vbs](invis.vbs) 可以讓 [該批次檔](checkapp.bat) 在執行時命令提示字元不會跳出來。

>## 補充知識
批次檔配合windows系統的工作排程器可讓電腦自動化定時執行批次作業，可以避免人為遺忘或漏失的風險。

## 如何製作批次檔?
```
1. 在記事本裡打一些文字，
2. 另存成 .bat 檔，就這樣。
```
備註:.bat檔就是batch的縮寫

## Batch file 範例 
以下用個wikipedia上簡單的[例子](https://en.wikipedia.org/wiki/Batch_file)來示意(注意 : DOS命令不區分大小寫!)
```bat
@ECHO OFF
ECHO Hello World!
PAUSE
```
解釋器(interpreter)會依次執行每一行，從第一行開始，@開頭的符號會阻止在執行該命令時顯示該命令。ECHO OFF命令可以永久關閉顯示後續的每一行，直到它再次打開。

@ECHO OFF通常是批次檔的第一行，防止顯示任何命令，包括第一行本身(若改成"ECHO ON"的話則會出現)。然後執行下一行並ECHO Hello World!輸出命令``Hello World!``。執行下一行，PAUSE會暫停腳本的執行並顯示等你按任意鍵繼續
```
Press any key to continue . . .
```
按下一個鍵後，腳本終止，因為沒有更多的命令。

方便查詢指令的小工具 : 
- [CMD命令速查手冊](http://www.cas.idv.tw/Documents/Micorsoft/CMDManual/CMD%E5%91%BD%E4%BB%A4%E9%80%9F%E6%9F%A5%E6%89%8B%E5%86%8A.asp)
- [.bat相關資源網址](batch_resource.md)(幫助過我解決過問題的相關資源我都放在這)

## 批次檔下註解的方法
```
::我是註解
REM 我也是註解
```
## CMD檢視執行中的程序
```
tasklist /v /fo list
```
在本機上，要檢視執行中的程序，只要輸入`tasklist`就可以了 

/fo 參數可以改變結果呈現的方式 

/v 顯示詳細工作資訊
## DOS和CMD有什麼不同?
DOS是英文Disk Operating System的縮寫，意思是「磁碟作業系統」，而(cmd.exe)只是在windows下的DOS模擬器，能運行絕大部分DOS命令，但是不是真正的DOS。

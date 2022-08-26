# AJAX

## 什麼是AJAX呢?

[AJAX](https://zh.wikipedia.org/zh-tw/AJAX)全名為「Asynchronous JavaScript and XML」（非同步的JavaScript與XML技術）簡單來說，是一種可以在 "背後" 和伺服器溝通的請求方式，當伺服器回傳資料後，由js來操作這些資料改怎樣使用，可以替換到某個標籤中，也可以做任何你想做的事情，這也是AJAX最大的優點，能在不重新加載整個頁面的前提下和伺服器交換資料，使得 Web 應用程式能更為迅速的回應使用者動作，並避免重複傳送那些沒有改變的資訊。  

## 同步與非同步

在AJAX這個名字裡，最難理解的部分就是非同步(Asynchronous)了，我覺得 ALPHAcamp 的[這篇文章](https://tw.alphacamp.co/blog/ajax-asynchronous-request)解釋的很好，想更了解的可以去看。

同步請求 (Synchronous request)： 客戶端 (client) 對伺服器端 (server) 送出 request ，並且在收到伺服器端的 response 之後才會繼續下一步的動作，等待的期間無法處理其他事情。這個作法並不理想，因為通常伺服器端的運算速度比本地電腦慢上好幾倍。  

非同步請求 (Asynchronous request)：客戶端 (client) 對伺服器端 (server) 送出 request 之後，不需要等待結果，仍可以持續處理其他事情，甚至繼續送出其他 request。Responese 傳回之後，就被融合進當下頁面或應用中。

而AJAX就是使用非同步的方式來提高使用者體驗~

## reference
- [維基百科 - AJAX](https://zh.wikipedia.org/zh-tw/AJAX)  
- [什麼是 Ajax？ 搞懂非同步請求 (Async request) 概念](https://tw.alphacamp.co/blog/ajax-asynchronous-request)

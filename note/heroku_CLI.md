# Heroku CLI cheat sheet

- 登入 Heroku 帳號。
```
heroku login
```
- 登入 Heroku 帳號(有兩個以上帳號的話選擇這個語法較好)。
```
heroku login -i
```
- 在當前 Heroku 帳號下，創建一個新的 Heroku APP。
```
heroku create 你-APP-的名字
```
- 將指定的 Heroku APP 設定為git推送的遠端資料庫。
- 用官方原本指令推到master可能會發生[錯誤](https://stackoverflow.com/a/66899451)
<!-- - Try Using git push heroku main instead of git push heroku master -->
```
heroku git:remote -a 你-APP-的名字
git add .
git commit –m "這次更新的註解"
git push heroku main    
```
- 查看檔案結構
```
tree /F
```
- 查看 Heroku 的log
```
heroku logs -a 你-APP-的名字
```
- 查看 Heroku 的log(持續打印日誌)
```
heroku logs --tail -a 你-APP-的名字 
```
- 查詢 Heroku APP 目前所使用的擴充功能(add-ons)。
```
heroku addons -a 你-APP-的名字
```
- 查詢 Heroku APP 的狀態資訊，可以看到目前免費時數的用量(根據Heroku的官方說法，如果免費的dayno超過每個月能清醒的550小時，那我們所有採用免費dyno的應用程式都會進入睡眠狀態，直到下個月來臨。)
```
heroku ps -a 你-APP-的名字
```
- 新增一個 hobby-dev 方案的 Heroku Postgres 至當前的 Heroku APP。
```
heroku addons:create heroku-postgresql:hobby-dev
```
- 查詢 Heroku 當中指定的設定變數(Config Vars)。
```shell
heroku config

# 以此例為 DATABASE_URL，得到結果是 Heroku Postgres 資料庫的連線位址。  
heroku config:get DATABASE_URL -a 你-APP-的名字

#在當前的 Heroku APP中新增一個鍵值配對 (Key:Value) 為TZ:"Asia/Taipei"的設定變數。
heroku config:add TZ="Asia/Taipei"
```
- heroku cli 更新 (如果在上傳時跳出類似 warning: heroku update available from 7.59.4 to 7.60.1. 這種訊息可以用指令進行更新)
```
heroku update
```
- 如果您的內存受限或應用程序啟動時間很慢，您可能需要考慮啟用該[preload選項](https://devcenter.heroku.com/articles/python-gunicorn#advanced-configuration)。這會在工作進程被派生之前加載應用程序代碼。
```
web: gunicorn hello:app --preload
```
Reference :  
[Heroku CLI Commands官方文檔](https://devcenter.heroku.com/articles/heroku-cli-commands)  
[從LINE BOT到資料視覺化：賴田捕手](https://ithelp.ithome.com.tw/users/20120178/ironman/2654?sc=hot)

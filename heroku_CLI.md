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
<!-- - Try Using git push heroku main instead of git push heroku master -->
```
heroku git:remote -a 你-APP-的名字
git add .
git commit –m "這次更新的註解"
git push heroku main    [Reason](https://stackoverflow.com/a/66899451)
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
- 查詢 Heroku APP 目前所使用的擴充功能(addons)。
```
heroku addons -a 你-APP-的名字

```
- 查詢 Heroku APP 的狀態資訊，可以看到目前免費時數的用量
```
heroku ps -a 你-APP-的名字
```
- 新增一個 hobby-dev 方案的 Heroku Postgres 至當前的 Heroku APP。
```
heroku addons:create heroku-postgresql:hobby-dev
```
- 查詢 Heroku 當中指定的設定變數，以此例為 DATABASE_URL，得到結果是 Heroku Postgres 資料庫的連線位址。
```
heroku config:get DATABASE_URL -a 你-APP-的名字
```
參考網站:[從LINE BOT到資料視覺化：賴田捕手](https://ithelp.ithome.com.tw/users/20120178/ironman/2654?sc=hot)

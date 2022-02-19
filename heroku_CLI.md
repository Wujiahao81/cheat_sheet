# Heroku CLI cheat sheet

### Step by step.
- 登入 Heroku 帳號。
```
heroku login -i
```
- 在當前 Heroku 帳號下，創建一個新的 Heroku APP。
```
heroku create 你-APP-的名字
```
- 將指定的 Heroku APP 設定為git推送的遠端資料庫。
```
heroku git:remote -a 你-APP-的名字
```

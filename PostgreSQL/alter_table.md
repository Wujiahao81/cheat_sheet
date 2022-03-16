ALTER TABLE 是用來對已存在的資料表結構作更改。語法型式如下：
```PostgreSQL
ALTER TABLE table_name [改變方式];
```
假設現在我們已經建立好一個 "userdata" 資料表：

| Id | name | userid | datetime | status  |
| :----: | :----: | :---- | :---- | :---- |

## 增加欄位 (ADD COLUMN)
語法
```PostgreSQL
ALTER TABLE table_name ADD column_name datatype;
```
例如，如果我們想增加一個 gender 欄位：
```PostgreSQL
ALTER TABLE userdata ADD gender VARCHAR(10);
```

## 更改欄位資料型別 (ALTER COLUMN TYPE)
語法
```PostgreSQL
ALTER TABLE table_name ALTER COLUMN column_name datatype;
```
例如，更改 gender 欄位的資料型別：
```PostgreSQL
ALTER TABLE userdata ALTER COLUMN gender DECIMAL(18, 2);
```

## 刪除欄位 (DROP COLUMN)
語法
```PostgreSQL
ALTER TABLE table_name DROP COLUMN column_name;
```
例如，刪除 gender 欄位：
```PostgreSQL
ALTER TABLE userdata DROP COLUMN gender;
```
## 修改表的名稱 (RENAME TABLE)
語法
```PostgreSQL
RENAME TABLE userdata to 新表名稱;
```
參考網站:[SQL ALTER TABLE](https://www.1keydata.com/tw/sql/sql-alter-table.html)

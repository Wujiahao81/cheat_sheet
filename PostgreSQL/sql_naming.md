# PostgreSQL 命名

## 大小寫

PostgreSQL對表格(table)名稱、欄位(field，又稱字段)名稱都是區分大小寫的,但是，PostgreSQL在SQL語句中對大小寫是不敏感的，在不加雙引號的情況下，創建時無論使用大寫還是小寫，表格中都會統一轉為小寫顯示的，因此查詢時也會將語句中的欄位名稱統一改成小寫，因此，此時使用[大小寫查詢均可](https://www.modb.pro/issue/10320)。

## 例子

```PostgreSQL
query = sql.SQL("INSERT INTO {}(uid, \"displayName\", language, \"pictureUrl\") VALUES(%s, %s, %s, %s)").format(sql.Identifier('customers'))
```
為什麼UID不用加雙引號「"」，displayName要加，language卻又不用加，下一個pictureUrl又要加雙引號，答案是`大小寫`，如果你在postgresql中建立了帶有大小寫的欄位名稱，請務必要加上雙引號!  

源自:[Postgresql採到的坑](https://ithelp.ithome.com.tw/articles/10274769)  
補充:[識別項（Identifier）和關鍵字 （Keyword）](https://docs.postgresql.tw/the-sql-language/sql-syntax/lexical-structure#4.1.1.-shi-bie-xiang-identifier-he-guan-jian-zi-keyword)

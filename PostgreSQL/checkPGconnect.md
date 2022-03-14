只需要修改以下""內資料成你自己的就可以使用
```python
    # Connect to an existing database
    connection = psycopg2.connect(user="使用者帳號",
                                  password="密碼",
                                  host="主機位址",
                                  port="5432",
                                  database="資料庫名稱")
```
[程式連結](checkPGconnect.py)
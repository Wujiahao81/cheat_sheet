# PostgreSQL設置時區
- 查看時區
```PostgreSQL
show time zone;
```
- 查看支持的時區列表
```PostgreSQL
SELECT * FROM pg_timezone_names;
```
- 時區設置成 Asia/Taipei
```PostgreSQL
set time zone 'Asia/Taipei';
```


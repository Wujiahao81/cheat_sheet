## SQL查看table內容
```PostgreSQL
SELECT * FROM public.table_name ORDER BY id ASC
```
## SQL查看table名稱
```PostgreSQL
SELECT   tablename   FROM   pg_tables  
WHERE   tablename   NOT   LIKE   'pg%'  
AND tablename NOT LIKE 'sql_%'
ORDER   BY   tablename;
```

## SQL查看表結構
```PostgreSQL
SELECT A
	.attnum,
	A.attname AS field,
	T.typname AS TYPE,
	A.attlen AS LENGTH,
	A.atttypmod AS lengthvar,
	A.attnotnull AS NOTNULL,
	b.description AS COMMENT 
FROM
	pg_class C,
	pg_attribute
	A LEFT OUTER JOIN pg_description b ON A.attrelid = b.objoid 
	AND A.attnum = b.objsubid,
	pg_type T 
WHERE
	C.relname = 'table_name' 
	AND A.attnum > 0 
	AND A.attrelid = C.oid 
	AND A.atttypid = T.oid 
ORDER BY
	A.attnum;
```

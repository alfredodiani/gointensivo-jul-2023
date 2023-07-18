
To create the table install sqlite3 then run the following commands:

```
sqlite3 db.sqlite3
```
then:
```sql
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL,PRIMARY KEY(id));
```

Indented:
```sql
CREATE TABLE orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,

    PRIMARY KEY(id)
);
```




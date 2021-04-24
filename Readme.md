# Ports and Adapter Pattern  

## Execute application  

```bash
docker-compose up --build
docker-compose exec app bash
```

## Run tests  
```bash
go test ./...
```

## Generate mocks
```bash
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

## enter SqLiteDb
```bash
sqlite3 sqlite.db
```

```sqlite
create table products(id string, name string, price float, status string);
.tables
select * from products;
```

```bash
go run main.go
```
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
mockgen -destination=application/mocks/application.go -source=applrun main.goication/product.go application
```

## Initialize cobra configs  
```bash
cobra init --pkg-name=github.com/rfaguiar/ports-adapter-pattern
cobra add cli
```
## Run main cobra cmds  
```bash
go run main.go cli
```

## refresh mod libs
```bash
go mod tidy
```
## connect sqlite db
```bash
sqlite3 db.sqlite
```

## create table products
```sqlite
create table products(id string, name string, price float, status string);
select * from products;
```

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

```bash
cobra init --pkg-name=github.com/rfaguiar/ports-adapter-pattern/application
```



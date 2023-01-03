### Генерация доки

дока по либе
```
swag --help
```
запуск генерации доки
```
swag init
```

### Создание миграций

для работы нужен [migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

**Создание миграции**
```
migrate create -ext mysql -dir migrations create_user_tables
migrate -database "mysql://user:pass@tcp(host:3306)/database" -path "../migrations/" up
```

**Применение миграции** 
```
migrate up|down
```
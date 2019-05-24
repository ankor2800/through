# Test through ID

### Разработка

Создать каталог проекта обязательно внутри `$GOPATH` (по умолчанию `~/go`):
```
mkdir -p ~/go/src/through/
```

Склонировать код проекта в созданный каталог:
```
cd ~/go/src/through/
git clone https://github.com/ankor2800/through.git .
```

Установить менеджер зависимостей `glide` и с помощью его установить все нужные пакеты
```
go get -u github.com/Masterminds/glide
glide install
```

```
golangci-lint run
```

### ENV
```
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USERNAME=default
DB_PASSWORD=
DB_DATABASE=default
INSERT_COUNT=1000 // кол-во заказов
```

### Controllers

#### /add/through
##### POST
##### Desription
Запрос на заполнение трёх таблиц заказов `orders_1`, `orders_2`, `orders_3`, со сквозными ID
через общую таблицу `through` с `autoincrement id` заполняемую `INSERT INTO`.

#### /add/auto
##### POST
##### Desription
Запрос на заполнение одной таблицы заказов `auto_orders` с `autoincrement id`.

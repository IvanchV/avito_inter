# User-segments

Cервис, хранящий пользователя и сегменты, в которых он состоит (реализовано создание, удаление сегментов, а также добавление и удаление пользователей в сегмент)

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- Gorilla Mux (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- pg (драйвер для работы с PostgreSQL)


# Начало работы
Настроить переменные окружения можно в файле `.env`

Пример конфигурации:  

`HTTP_ADDR=0.0.0.0:8080` 

`POSTGRES_PASSWORD=5525`

`POSTGRES_USER=Vasily`

`POSTGRES_DB=avito`

`PG_URL="postgres://Vasily:5525@postgres:5432/avito?sslmode=disable`

Запустить сервис можно с помощью команды `make compose-up` 

Документацию после запуска сервиса можно посмотреть по адресу `http://.../documentation/`

В примере конфигурации документация находится по адресу `http://localhost:8080/documentation/`


Выключает сервис команда `make compose-down`


## Примеры

Некоторые примеры запросов
- [Добавление нового сегмента](#create-segment)
- [Удаление сегмента](#delete-segment)
- [Изменение сегментов у пользователя](#change-segment)
- [Список всех сегментов пользователя](#get-segment)


### Добавление сегмента <a name="create-segment"></a>

Запрос:
```curl
curl -X 'POST' \
  'http://localhost:8080/create_segment' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "30off"
}'
```
Пример ответа:
```json
{
   "Ans": "Segment created"
}
```

### Удаление сегмента <a name="delete-segment"></a>

Запрос на удаение сегмента "1":
```curl
curl -X 'DELETE' \
  'http://localhost:8080/delete_segment' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "1"
}'
```
Пример ответа:
```json
{
   "Ans": "Segment deleted"
}
```

### Изменение сегментов у пользователя <a name="change-segment"></a>

Запрос на изменение сегментов у пользователя с id=1:
```curl
curl -X 'PUT' \
  'http://localhost:8080/change_segment/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "add": [
    "30off"
  ],
  "delete": [
    "1"
  ]
}'
```
Пример ответа:
```json
{
  "message": "success"
}
```

### Изменение сегментов у пользователя <a name="get-segment"></a>

Запрос на получение сегментов пользователя с id=1:
```curl
curl -X 'GET' \
  'http://localhost:8080/user_segment/1' \
  -H 'accept: application/json'
```
Пример ответа:
```json
{
   "segments": [
      "1"
   ]
}
```
# Обзор

Веб-сервис для тестирования http запросов. 


# Быстрый старт

**Для работы программы требуется Golang** 1.22.5

### Установка

1. Клонирование репозитория

    ```shell
    git clone https://github.com/Fego02/jabka-stubs.git
    ```

2. Сборка исполняемого файла

    ```shell
    cd jabka-stubs/cmd/stub-server
    go build -o stub-server.exe
    ```

3. Запуск сервера

    ```shell
    ./stub-server.exe
    ```

Параметры запуска: 
-address="127.0.0.1" 
-port="8080" 
-log="" (по стандарту записывает в stdout, если указать "none", то отключается, иначе воспринимается как путь к файлу) 
-log_for_matched_only=true (отключает журналирование для запросов, для которых не нашлась ни одна заглушка)


# Использование

#### Создание простой заглушки 

**POST** /stubs/http-stubs
Схема тела запроса: application/json

|                          |                                                 |                                                                                                                                                                                                         |
| ------------------------ | ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name<br>(опционально)    | *string* <br><br>Уникальное имя заглушки        |                                                                                                                                                                                                         |
| request<br>(опционально) | *object*<br>                                    |                                                                                                                                                                                                         |
|                          | Методы<br>(возможен только один вариант)        |                                                                                                                                                                                                         |
|                          | method<br>                                      | *string <br>(GET, POST, DELETE, ...)*<br><br>Метод запроса http                                                                                                                                         |
|                          | methods_list                                    | *array*<br><br>Список методов запросов http                                                                                                                                                             |
|                          |                                                 |                                                                                                                                                                                                         |
|                          | Путь<br>(возможен только один вариант)          |                                                                                                                                                                                                         |
|                          | url                                             | *string*<br><br>Путь запроса                                                                                                                                                                            |
|                          | url_matches                                     | *string (golang regex)*<br><br>Регулярное выражение для нахождения совпадений url                                                                                                                       |
|                          | url_not_matches                                 | *string (golang regex)*<br><br>Регулярное выражение для нахождения несовпадений url                                                                                                                     |
|                          |                                                 |                                                                                                                                                                                                         |
|                          | Заголовки<br>(возможен только один вариант)<br> | Каждое имя заголовка должно быть представлено лишь единожды в одном из полей                                                                                                                            |
|                          | headers                                         | *object* <br><br>Шаблоны заголовков для сопоставления в форме : { "": "" }                                                                                                                              |
|                          | headers_matches                                 | *object*<br><br>Шаблоны заголовков для сопоставления в форме : { "": "" }<br><br>Значение заголовков - регулярное выражение для нахождения совпадений значений заголовков <br>*string (golang regex)*   |
|                          | headers_not_matches                             | *object*<br><br>Шаблоны заголовков для сопоставления в форме : { "": "" }<br><br>Значение заголовков - регулярное выражение для нахождения несовпадений значений заголовков <br>*string (golang regex)* |
|                          |                                                 |                                                                                                                                                                                                         |
|                          | Тело<br>(возможен только один вариант)          |                                                                                                                                                                                                         |
|                          | body                                            | *string*<br><br>Строковое содержимое тела для сопоставения                                                                                                                                              |
|                          | body_matches                                    | *string (golang regex)*<br><br>Регулярное выражение для нахождения совпадений строкового содержимого тела<br>                                                                                           |
|                          | body_not_matches                                | *string (golang regex)*<br><br>Регулярное выражение для нахождения несовпадений строкового содержимого тела                                                                                             |
|                          | body_hex_matches                                | *string (golang regex)*<br><br>Регулярное выражение для нахождения совпадения hex-закодированного содержимого тела                                                                                      |
|                          | body_hex_not_matches                            | *string (golang regex)*<br><br>Регулярное выражение для нахождения несовпадения hex-закодированного содержимого тела                                                                                    |
| response                 | *object*                                        |                                                                                                                                                                                                         |
|                          | status                                          | *integer (>= 100 and <= 999)*<br>По умолчанию: 200<br>http статус код возврата                                                                                                                          |
|                          | headers                                         | *object* <br><br>Шаблоны заголовков ответа в форме : { "": "" }<br>                                                                                                                                     |
|                          | body                                            | *string*<br><br>Строковое содержимое тела ответа                                                                                                                                                        |
| properties               | *object*                                        |                                                                                                                                                                                                         |
|                          | is_logging_enabled                              | *boolean*<br>По умолчанию: True<br><br>Активация / деактивация журналирования                                                                                                                           |
|                          | delay                                           | *integer*<br>По умолчанию: 0<br><br>Задержка ответа в мс<br>                                                                                                                                            |
|                          | priority                                        | integer<br>По умолчанию: 0<br>                                                                                                                                                                          |




##### Пример запроса создания заглушки

```
{
    "name": "Сorrect form for creating a stub",
    "request": {
        "method": "GET",
        "url": "/messages/greeting",
        "headers": {
            "Authorization": "Basic UGFya2VyOm5vd2F5aG9tZQ==",
            "Content-Language": "en-US"
        }
    },
    "response": {
        "status": 201,
        "body": "Hello, World!",
        "headers": {
            "Content-Type": "text/plain"
        }
    }
}
```

**Ответ на запрос**
201 Created
Stub created successfully for Сorrect form for creating a stub on /messages/greeting
Location: /stubs/http-stubs/1


#### Создание заглушки для работы с файлами

**POST** /stubs/http-stubs
Схема тела запроса: multipart/form-data

| Key                            | Value                                                                               |
| ------------------------------ | ----------------------------------------------------------------------------------- |
| stub-data<br>(обязательно)     | JSON (формат создания простой заглушки)<br><br>Файл, описывающий поведение заглушки |
| request-body<br>(опционально)  | file<br><br>Файл тела запроса                                                       |
| response-body<br>(опционально) | file<br><br>Файл тела ответа                                                        |

#### Изменение заглушки 

Для изменения уже поставленной заглушки достаточно отправить повторный запрос создания с указанием имени изменяемой заглушки.

##### Пример изменения заглушки

На сервере установлена заглушка с именем "Just a stub" по пути /examples, возвращая при запросе GET сообщение "Hello, World!" с кодом 200.

Изменить заглушку можно с помощью запроса создания с указанием того же имени - "Just a stub"

Отправка запроса POST по пути /stubs/http-stubs с телом JSON 

```
{
    "name": "Just a stub",
    "request": {
        "method": "GET",
        "url": "/examples"
    },
    "response": {
        "status": 201,
        "body": "Hello, World!",
        "headers": {
            "Content-Type": "text/plain"
        }
    }
}
```

**Ответ на запрос**
201 Created
Stub created successfully for Just a stub on /examples


#### Удаление заглушки по id
Удаление заглушки осуществляется **DELETE** запросом на адрес ресурса заглушки, который выдается при создании в заголовке **Location**

Пример запроса удаления
```
	curl -X DELETE {address}:{port}/stubs/http-stubs/{id}
```


#### Обновить заглушку по id
Обновить заглушку можно **PUT** запросом на {address}:{port}/stubs/http-stubs/{id}


#### Удаление всех заглушек

Удаление заглушки осуществляется **DELETE** запросом на {address}:{port}/stubs/http-stubs

Пример запроса удаления
```
curl -X DELETE {address}:{port}/stubs/http-stubs
```


#### Получить данные о заглушке
Получение заглушки осуществляется **GET** запросом на {address}:{port}/stubs/http-stubs/{id}

Пример запроса
```
curl -X GET {address}:{port}/stubs/http-stubs/{id}
```


#### Получить данные о всех заглушках
Получение всех заглушек осуществляется **GET** запросом на {address}:{port}/stubs/http-stubs

Пример запроса
```
curl -X GET {address}:{port}/stubs/http-stubs
```



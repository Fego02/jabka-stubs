# Обзор

Веб-сервис для тестирования http запросов. 


# Быстрый старт

**Для работы программы требуется Golang** 

### Установка

1. Клонирование репозитория 
   ```shell 
   git clone https://github.com/Fego02/jabka-stubs.git
   cd jabka-stubs/server
```
   
2. Сборка исполняемого файла
   ```shell
   go build -o stub-server.exe src/main.go
```

3. Запуск сервера
   ```shell
   ./stub-server.exe
```

Параметры запуска:
-address="127.0.0.1"
-port="8080"


# Использование

#### Создание простой заглушки 

==POST== /generate
Схема тела запроса: application/json

- **name** 
  *string* 
  Уникальное имя заглушки
  (обязательно)  
- **request**
  *object*
  (обязательно)
	*  **method**
	  *string* (GET, POST, DELETE, ...)
	  Метод запроса http
	  (обязательно)
	- **url**
	  *string*
	  Путь запроса
	  (обязательно)
	- **headers**
	  *object* 
	   Шаблоны заголовков для сопоставления в форме : { "": "" }
	   (опционально)
- **response
  *object*
  (обязательно)
	*  **status**
	  *integer* (>= 100 and <= 999)
	  http статус код возврата
	  (обязательно)
	- **body**
	  *string*
	  Строковое содержимое тела ответа
	  (обязательно)
	- **headers**
	  *object* 
	   Шаблоны заголовков для сопоставления в форме : { "": "" }  
	   (опционально)
- properties (ПОКА НЕ РЕАЛИЗОВАНО)
  *object*
  (опционально)
	- **delay**
	  *integer*
	  По умолчанию: 0
	  Задержка ответа в мс
	- **is_logging_enabled**
	  *boolean*
	  По умолчанию: True
	  Активация / деактивация журналирования

##### Пример запроса создания заглушки

```json
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


#### Создание заглушки для работы с файлами

==POST== /generate
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

Отправка запроса POST по пути /generate с телом JSON 

```json
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






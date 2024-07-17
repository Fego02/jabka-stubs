
# Обзор

Система для проведения функционального тестирования веб-сервера 

# Быстрый старт

**Для запуска тестирования требуются следующие программы:**
- Postman
- Postman CLI
- Node.js >= v16
- newman
- newman-reporter-htmlextra


Установить данные программы для linux можно с помощью скрипта 

1. Перейти в директорию testing
2. Запустить скрипт install-linux.sh
   ```shell
   bash install-linux.sh
```


# Тестирование

Тестирование по умолчанию проходит по локальному адресу с портом 8080
http::/localhost:8080

Для изменения данных параметров требуется поменять значения переменных "address" и "port" в файле test.postman_collection.json
Переменные "address" и "port" располагаются в поле "variable" в самом низу JSON файла

```json
"variable": [
        {
            "key": "address",
            "value": "localhost",
            "type": "string"
        },
        {
            "key": "port",
            "value": "8080",
            "type": "string"
        },
        ...
```
\
s### Запуск тестирования

1. Запустить скрипт create-test-report.sh
   $1 - путь к файлу коллекции (JSON)
   $2 - путь для создания отчета (html)
   
   ```shell
   bash create-test-report.sh ./test.postman_collection.json test-report.html
```

2.  Открыть получившийся отчет
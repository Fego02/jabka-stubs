
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

Для изменения данных параметров требуется поменять значения переменных "address" и "port" в файле коллекции тестов (содержит postman_collection в названии).
Переменные "address" и "port" располагаются в поле "variable" в самом низу JSON файла.

При работе с коллекцией "test-files-relative-path.postman_collection.json" необходимо указать абсолютный путь до папки jabka-stubs/test/test-files в переменной "test-files_path" в файле коллекции. Переменная "test-files_path" располагается в поле "variable" в самом низу JSON файла.

```
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
	{
		"key": "generate_path",
		"value": "/stubs/http-stubs",
		"type": "string"
	},
	{
		"key": "test-files_path",
		"value": "/home/user/jabka/jabka-stubs/test/test-files",
		"type": "string"
	},
        ...
```
### Запуск тестирования

1. Запустить скрипт create-test-report.sh
   $1 - путь к файлу коллекции (JSON)
   $2 - путь для создания отчета (html)
   
   ```shell
   bash create-test-report.sh ./test.postman_collection.json test-report.html
   ```

2.  Открыть получившийся отчет

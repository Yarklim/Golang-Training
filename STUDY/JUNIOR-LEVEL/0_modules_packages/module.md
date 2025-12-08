Study/
|
|----- main.go
|----- greeting.go
|
|----- Users/
|			|----- user1.go
|			|----- user2.go
|			|----- user3.go
|
|----- Database/
|			|----- database1.go
|			|----- database2.go
|
|----- Tools/
|			|----- Logger/
|			|			|----- logger1.go
|			|			|----- logger2.go
|			|
|			|----- Metrics/
|			|			|----- metric1.go

# Users, Database, Logger, Metrics - это Пакеты (папка с .go файлами внутри на первом уровне)
# В рамках пакета все функции, переменные, структуры доступны в любом файле пакета.
# Названия с большой буквы доступны в других пакетах, с маленькой - не доступны.

## Загрузка библиотек
- go get github.com/k0kubun/pp

## Проверка зависимостей go mod tidy
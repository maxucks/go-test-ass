# Общее

Я сделал 2 сервиса: app и collector. Collector - читает очередь и пишет ивенты в Clickhouse. App - основной сервис который предоставляет CRUD к товарам.

# Запуск

1. Прогнать миграции с гусем (ps. Я попытался сделать скрипт автомигратор и засунуть его в докер, но гусь в баше почему-то отказывается запускать команду с env переменными)

```bash
goose -dir migrations/postgres postgres "postgres://root:1234@localhost:5432/postgres?sslmode=disable" up
```

```bash
goose -dir migrations/clickhouse clickhouse "clickhouse://root:1234@localhost:9000/main" up
```

2. make run - Запустит все сервисы

# Примечания

Формат эндпоинтов /good/remove, /good/create и так далее - не соответствует REST. На самом деле мне не принципиально, но здесь я позволил себе отойти от тз

- GET /projects/goods - Получение списка товаров
- POST /projects/{projectID}/goods - Создание товара
- PATCH /projects/{projectID}/goods/{goodsID} - Изменение товара
- DELETE /projects/{projectID}/goods/{goodsID} - Удаление товара
- PATCH /projects/{projectID}/goods/{goodsID}/priority - Реприоритизация

В .env можно поменять практически все настройки

# Прочее

- Основной сервис работает на localhost:8000
- Графана доступна на localhost:3000 - можно потыкать кликхаус

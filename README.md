# Game
Задача: 
Реализовать простой HTTP сервер для управления пользователями на языке Go

Основная логика приложения - создание пользователя, который выполняет какие-либо целевые действия, например, вводит реферальный код, подписывается на телеграм канал или твиттер и получает за это награду в виде поинтов. Награду за каждое задание вы можете определить самостоятельно, также вы можете добавить другие задачи, дайте волю фантазии

Нужно реализовать следующий функционал:
Middleware авторизация по Access token ко всем эндпоинтам (например JWT)
Реализация HTTP API:
GET /users/{id}/status - вся доступная информация о пользователе
GET /users/leaderboard - топ пользователей с самым большим балансом
POST /users/{id}/task/complete - выполнение задания 
POST /users/{id}/referrer - ввод реферального кода (может быть id другого пользователя)
Создание хранилища для всех этих данных по пользователю (postgres). Обязательно использование инструментов для миграций (golang-migrate)
Сборка всего проекта в docker-compose

Дополнительные требования:
Протестировать все указанные маршруты с помощью Postman или аналогичного инструмента
Обеспечить обработку ошибок (например, неверные данные, несуществующие пользователи и т.д.)
Не забывайте про принципы SOLID и чистоту кода
Управление пользователями:

1. POST /api/users/ - Добавить нового пользователя
2. DELETE /api/users/{id} - Удалить пользователя по ID
3. GET /api/users/{id}/link - Сгенерировать VLESS-ссылку для пользователя
4. GET /api/users/all - Получить список всех пользователей

Управление Xray:

1. POST /api/management/restart - Перезапустить Xray
2. GET /api/management/status - Получить текущий статус Xray
3. POST /api/management/start - Запустить Xray
4. POST /api/management/stop - Остановить Xray

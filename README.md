Управление пользователями (АДМИНКА):
1. GET /admin/users — список всех пользователей.
2. GET /admin/users/{id} — информация о конкретном пользователе.
3. POST /admin/users — добавить нового пользователя.
4. (DELETE) / (POST) /admin/users/{id} — удалить пользователя.

Управление серверами (АДМИНКА):
1. GET /admin/servers — список всех серверов.
2. POST /admin/servers — добавить новый сервер.
3. (DELETE) / (POST) /admin/servers/{(id) / (ip)} — Удалить сервер.
4. GET /admin/xray/status — получить статус xray.
5. POST /admin/xray/restart — перезапустить xray.
6. GET /admin/servers/{(id) / (ip)}/stats — получить статистику по конкретному серверу.

Статистика (если она когда-то появится, офк...) (АДМИНКА):
1. GET /admin/stats/* — получить статистику по использованию VPN.
/admin/stats/{server_id}/*
2. GET /admin/logs — получить логи сервера. (если нужно, то надо будет реализовать в xray_manager)



Управление подключением (КЛИЕНТ):
1. GET /client/config — получить конфигурацию для подключения (или готовой ссылкой, или данными по отдельности).
2. GET /client/servers — список всех доступных серверов
это надо нам?...

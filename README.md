# Сервис динамического сегментирования пользователей

### Инструкция по запуску:
- Склонируйте репозиторий
- Войдите в директорию проекта
- ```docker-compose up --build```

## API

### POST /api/create

Body:
```
{
    "name": "TEST8",
    "percentage": 50
}
```
Code: 200\
Response:
```
{
    "affected_users": [
        1,
        2,
        6,
        7
    ]
}
```

### POST /api/set

Body:
```
{
    "user_id": 1,
    "add": ["SEGMENT1", "SEGMENT_NOT_EXISTS", "SEGMENT2", "SEGMENT_ALREADY_PRESENT"],
    "remove": ["SEGMENT_NOT_EXISTS", "SEGMENT_NOT_PRESENT", "TEST8"]
}
```
Code: 200\
Response:
```
{
    "add_error": [
        "SEGMENT_NOT_EXISTS",
        "SEGMENT_ALREADY_PRESENT"
    ],
    "remove_error": [
        "SEGMENT_NOT_EXISTS",
        "SEGMENT_NOT_PRESENT"
    ]
}
```

### GET /api/user/1/

Code: 200\
Response:
```
{
    "segments": [
        "TEST1",
        "TEST2",
        "TEST8"
    ]
}
```

### DELETE /api/delete

Body:
```
{
    "name": "TEST2"
}
```
Code: 200
# Tests

## Unit tests: (task_service_test)

### CreateTask:

1. [x] корректная задача;
2. [x] текст пустой → `ErrTaskTextEmpty`;
3. [x] текст пробелы → `ErrTaskTextEmpty`;
4. [x] текст длиной > 1000 → `ErrTaskTextTooLong`;
5. [x] ошибка Repository → проброс наверх

### DeleteTask:

1. [x] существующая задача → передаётся в Delete;
2. [x] несуществующая → `ErrTaskNotFound`;
3. [x] ошибка Repository → проброс наверх

### CompleteTask:

1. [x] существующая невыполненная задача → Completed = true, возвращается задача;
2. [x] несуществующая → `ErrTaskNotFound`;
3. [x] ошибка Repository → проброс наверх

## Api tests: (api_test) 

### Данные в формате Json

### POST `/tasks`

1. [x] Корректный запрос → `201 Created`. Возвращается созданная задача.
2. [x] Некорректный JSON → `400 Bad Request`.
3. [x] Пустой текст → `400 Bad Request`.
4. [x] Текст из пробелов → `400 Bad Request`.
5. [x] Текст длиной > 1000 символов → `400 Bad Request`.

### PATCH `/tasks/{id}` (сделать задачу выполненной)

1. [x] Корректный запрос → `204 No Content` и `completed` становится `true`. Возвращается обновленная задача.
2. [x] Несуществующая задача → `404 Not Found`.
3. [x] Некорректный идентификатор → `400 Bad Request`.

### DELETE `/tasks/{id}`

1. [x] Корректный запрос → `204 No Content`. Задача удаляется из БД
2. [x] Некорректный идентификатор → `400 Bad Request`.
3. [x] Несуществующая задача → `404 Not Found`.

### GET `/tasks`

1. [ ] Запрос → `200 OK`. Возвращаются НЕвыполненные задачи.

### GET `/tasks/archive`

1. [ ] Запрос → `200 OK`. Возвращаются выполненные задачи.

### Общие API tests

1. [ ] Неподдерживаемый HTTP-метод → `405 Method Not Allowed`.
2. [ ] Техническая ошибка сервера → `500 Internal Server Error`.
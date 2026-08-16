task_service_test:

### CreateTask:

1. [x] корректная задача;
2. [x] текст пустой/пробелы → ErrTaskTextEmpty;
3. [x] текст длиной > 1000 → ErrTaskTextTooLong;
4. [x] ошибка Repository → проброс наверх

### DeleteTask:

1. [x] существующая задача → передаётся в Delete;
2. [x] несуществующая → ErrTaskNotFound;
3. [x] ошибка Repository → проброс наверх

### CompleteTask:

1. [x] существующая невыполненная задача → Completed = true, возвращается задача;
2. [x] несуществующая → ErrTaskNotFound;
3. [x] ошибка Repository → проброс наверх
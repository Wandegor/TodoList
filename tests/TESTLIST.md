### CreateTask:

1. [x] корректная задача;
2. [ ] текст пустой/пробелы → ErrTaskTextEmpty;
3. [ ] текст длиной >= 1001 → ErrTaskTextTooLong;

### CompleteTask:

1. [x] существующая невыполненная задача → Completed = true, возвращается задача;
2. [ ] несуществующая → ErrTaskNotFound;
3. [ ] при ошибке поиска Update не вызывается.

### DeleteTask:

1. [x] существующая задача → передаётся в Delete;
2. [ ] несуществующая → ErrTaskNotFound;
3. [ ] при ошибке поиска Delete не вызывается.

### GetActiveTasks:

1. [ ] возвращаются только задачи с Completed=false;

### GetArchivedTasks:

1. [ ] возвращаются только задачи с Completed=true;
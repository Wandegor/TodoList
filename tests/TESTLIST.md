### CreateTask:

1. [x] корректная задача;
2. [x] текст пустой/пробелы → ErrTaskTextEmpty;
3. [x] текст длиной >= 1001 → ErrTaskTextTooLong;

### DeleteTask:

1. [x] существующая задача → передаётся в Delete;
2. [x] несуществующая → ErrTaskNotFound;

### CompleteTask:

1. [x] существующая невыполненная задача → Completed = true, возвращается задача;
2. [x] несуществующая → ErrTaskNotFound;

### GetActiveTasks:

1. [ ] возвращаются только задачи с Completed=false;

### GetArchivedTasks:

1. [ ] возвращаются только задачи с Completed=true;
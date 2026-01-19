# lvl1/task1 — Встраивание (Embedding) Human в Action на Go

Этот репозиторий — учебный пример на Go, который показывает **встраивание (embedding)** структуры `Human` в структуру `Action`.
Когда `Human` встроен в `Action`, методы `Human` можно вызывать у объекта `Action` напрямую (как будто они принадлежат `Action`).

## Версия и модуль

Проект использует:
- Go: **1.25.4**
- Module: **lvl1/task1**

Импорты в `main.go` соответствуют модулю:
- `lvl1/task1/human`
- `lvl1/task1/action`

## Запуск

В корне проекта:
```bash
go run main.go

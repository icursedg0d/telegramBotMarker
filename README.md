# Telegram Bookmark Bot

Простой Telegram-бот, реализованный без использования сторонних библиотек, для добавления и случайного получения закладок.

## 📋 Команды

1. **Добавление закладки**  
   Отправьте ссылку в чат, чтобы добавить её в список закладок.

2. **Получение случайной закладки**  
   `/rnd` — возвращает случайную закладку из списка.

3. **Помощь**  
   `/help` — отображает список доступных команд и краткую информацию.

4. **Приветственное сообщение**  
   `/start` — знакомит пользователя с ботом при первом взаимодействии.

---

## ⚙️ Запуск

1. Убедитесь, что у вас установлен Go.
2. Склонируйте репозиторий и перейдите в папку проекта:

```bash
 Скомпилируйте проект: go build
 Запустите бота, указав токен Telegram: ./tg -tg-bot-token "ВАШ_TELEGRAM_ТОКЕН"
```

🔧 Расширение функционала

Бот построен с возможностью лёгкого расширения:

    • Добавление новых сервисов (например, управление задачами или напоминания).
    • Реализация интеграции с внешними API.
    • Поддержка хранения данных в базе (например, SQLite).

🛠 Технические детали:

    • Язык разработки: Go
    • Особенности:
    • Полностью без сторонних библиотек.
    • Работа с Telegram Bot API через HTTP-запросы.
    • Хранение данных в памяти.

Пример использования

    Пользователь: https://example.com
    Бот: Закладка добавлена!

    Пользователь: /rnd
    Бот: https://example.com

© 2024. Сделано с ❤️ и минимализмом.

# Тестовые задания.

### 1. Backend тестовое задание (GO)

Сделать клиента для получения курсов.

https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1

Добавить возможность получать курс для определенной криптовалюты.

Курсы обновляем не чаще чем раз в 10 минут.

### 2. Backend тестовое задание (GO)

Сделать парсер.

https://hypeauditor.com/top-instagram-all-russia/

Берем первую страницу результата выборки.
Собираем данные по всем колонками (рейтинг, имя, ник и т.д.)

На выходе получаем csv файл с 50 строчками.

# Описание

Программа работает из командной строки.
Для запуска программы служит исполняемый файл `app.exe`, который находится в корне проекта.

+ Запуск программы с аргументом `.\app.exe polling`, заставит программу опрашивать API раз в 10 минут и сохранять в csv файл `coins.csv`. CSV файл появится рядом с исполняемым exe файлом.
+ Если использовать команду `.\app.exe get bitcoin` или `.\app.exe get btc`, то ты получишь конкретную монету прямо в консоль.
Если вдруг захочется выйти из программы во время опроса API, используй `Ctrl + C`.

Так же программа содержит в себе небольшой скрапер. Чтобы его запустить, нужно выполнить команду `.\app.exe scrape`. На рабочем столе появится csv файл с извлеченными данными.
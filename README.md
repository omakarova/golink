Сокращатель ссылок - сервис, который позволяет пользователю создавать более
короткие адреса, которые лучше передавать другим пользователям и собирает
статистику по совершенным переходам. 

Так как программа является учебно-тестовой, а не промышленной, сделан ряд допущений:
- используется HTTP Basic Authorization (не требуется для регистрации нового пользователя и перехода по короткой ссылке)
- пароль передается в теле POST запроса в открытом виде. И в открытом виде хранится в базе.
- я не проверяю имя пользователя и пароль на предмет "плохих" символов. Надеюсь, что сервис-клиент или фронтенд позаботились о том, что имя пользователя и пароль набраны латиницей.
- нет защиты от DDoS
- нет защиты от обрыва соединения с базой
- не используется защищенное соединение

## Собрать и запустить

### База и таблицы

- создаем MySQL базу ***ИМЯ_БАЗЫ*** и назначаем себе привилегии на чтение-запись-создание таблиц.
Например, так

    ```create database golink;
    grant all privileges on golink.* to golink@localhost identified by '123';
    use golink;
    ```
- создаем три таблицы: для хранения пользователей, для ссылок и для статистики
```
    CREATE TABLE
	    `users` (
	        `id` INT(11) NOT NULL AUTO_INCREMENT,
	        `username` CHAR(32) NOT NULL,
            `password` CHAR(128) NOT NULL,
	        `auth` CHAR(255) NOT NULL,
	        PRIMARY KEY(`id`),
     		UNIQUE(`username`, `password`)
	    );

    CREATE TABLE
	    `links` (
	        `id` INT(11) NOT NULL AUTO_INCREMENT,
	        `longurl` CHAR(255) NOT NULL,
	        `shorturl` CHAR(255) NOT NULL UNIQUE,
            `userid` INT(11) NOT NULL,
 	        PRIMARY KEY(`id`),
    		FOREIGN KEY `linkToUserFK` (`userid`)
            REFERENCES users(id)
            ON DELETE CASCADE
	    );

    CREATE TABLE
	    `statistics` (
	        `id` INT(11) NOT NULL AUTO_INCREMENT,
	        `linkid` INT(11) NOT NULL,
	        `referer` CHAR(255),
            `f_date_time` DATETIME NOT NULL,
	        PRIMARY KEY(`id`),
    		FOREIGN KEY `statToLinkFK` (`linkid`)
            REFERENCES links(id)
            ON DELETE CASCADE
	    );
```

### Конфиг

 В конфиг-файле **config.toml** прописывает строку коннекта базы.

 Например,
 ***golink:123@/golink?charset=utf8***

 В общем случае:

  ```php
  <username>:<pw>@tcp(<HOST>:<port>)/<dbname>
  ```

### Сборка и запуск
 Берем ветку master и собираем исполняемый файл в IDE или из командной строки.
 Например, так:
```
      go build ./project/myLinkMain.go
```

 Потом запустить исполняемый файл.

 По умолчанию сервер стартует на порте 3000.

 При запуске под windows могут понадобиться разрешения на брандмауэре.

## API

### Работа с пользователем

● Регистрация пользователя (авторизация не требуется).
**http://HOST_NAME/api/users**
**POST**
В теле запроса должен содержаться JSON
```
{
    "username" : "Имя пользователя латиницей без пробелов и "плохих" символов",
    "password" : "Пароль латиницей без пробелов и "плохих" символов"
}
```
Варианты ответов: StatusNotAcceptable (например, такой уже есть), StatusOK, StatusBadRequest

● Получение информации о текущем авторизованном пользователе.
**http://HOST_NAME/api/user**
**POST**
Требует авторизации

Ответ должен содержать JSON из структуры
```
type UserInfo struct {
	Username        string `json:"username"`
	LinksCount       int `json:"linkscount"`
}
```
Еще возможны ответы: 401 Unauthorized

### Короткие ссылки пользователя

● создание новой короткой ссылки
**http://HOST_NAME/api/links**
**POST**
Требует авторизации
В теле запроса должен содержаться JSON
```
{
    "url" : "URL"
}
```
Ответ должен содержать JSON вида
```
{"shorturl":"RAND_STRING"}
```

**Внимание, эта короткая строка (RAND_STRING) и является одновременно короткой ссылкой (можно переходить на http://HOST_NAME/RAND_STRING) и идентификатором ссылки, который будет использоваться в API.**

● получение всех созданных коротких ссылок пользователя
http://HOST_NAME
● получение информации о конкретной короткой ссылке пользователя (также
включить количество переходов)
http://HOST_NAME
● удаление короткой ссылки пользователя
http://HOST_NAME
Статистика по ссылкам
● получение временного графика количества переходов с группировкой по дням,
часам, минутам.
http://HOST_NAME
● получение топа из 20 сайтов иcточников переходов
http://HOST_NAME
Ссылки
●
http://HOST_NAME

Примечание: теоретически, любой метод может вернуть StatusInternalServerError

 //TODO
 - Сделать нормальную авторизацию,
 - типизировать ошибки
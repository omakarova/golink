Сокращатель ссылок - сервис, который позволяет пользователю создавать более
короткие адреса. Кроме того, он собирает статистику по совершенным переходам.

Так как программа является учебно-тестовой, а не промышленной, сделан ряд допущений:

- используется HTTP Basic Authorization (не требуется для регистрации нового пользователя и перехода по короткой ссылке)
- пароль передается в теле POST запроса в открытом виде. И в открытом виде хранится в базе.
- имя пользователя и пароль не проверяются на предмет "плохих" символов. Предполагается, что сервис-клиент или фронтенд позаботились о том, что имя пользователя и пароль набраны латиницей.
- нет защиты от DDoS
- нет защиты от обрыва соединения с базой
- не используется защищенное соединение

## Собрать и запустить

### База и таблицы

- создаем MySQL базу и назначаем себе привилегии на чтение-запись-создание таблиц.

Например, так

    ```create database golink;
    grant all privileges on golink.* to golink@localhost identified by '123';
    use golink;
    ```
- создаем три таблицы: для хранения пользователей, ссылок и статистики
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

 В конфиг-файле **config.toml** прописываем строку коннекта базы.

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

 Потом запускаем исполняемый файл.

 По умолчанию сервер стартует на порте 3000.
 Порт дожен быть свободным от других приложений.
 При запуске под windows могут понадобиться разрешения на брандмауэре.

# API

## Работа с пользователем

### Регистрация пользователя (авторизация не требуется).

**http://HOST_NAME/api/users**

Метод: **POST**

В теле запроса должен содержаться JSON
```
{
    "username" : "Имя пользователя латиницей без пробелов и "плохих" символов",
    "password" : "Пароль латиницей без пробелов и "плохих" символов"
}
```
Варианты ответов: StatusNotAcceptable (например, такой пользователь уже есть), StatusOK, StatusBadRequest


### Получение информации о текущем авторизованном пользователе.

**http://HOST_NAME/api/user**

Метод: **GET**

Требует авторизации

Ответ должен содержать JSON из структуры
```
type UserInfo struct {
	Username        string `json:"username"`
	LinksCount       int `json:"linkscount"`
}
```
Еще возможны ответы: 401 Unauthorized



## Короткие ссылки пользователя

### создание новой короткой ссылки

**http://HOST_NAME/api/links**

Метод: **POST**

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

Еще возможны ответы: 401 Unauthorized, StatusNotAcceptable, StatusBadRequest


### получение всех созданных коротких ссылок пользователя

**http://HOST_NAME/api/links**

Метод: **GET**

Требует авторизации

Ответ должен содержать json-массив всех ссылок данного пользователя (возможно, пустой)

Еще возможны ответы: 401 Unauthorized


### получение информации о конкретной короткой ссылке пользователя (также
включить количество переходов)

**http://HOST_NAME/api/links/:id**

Метод: **GET**

Требует авторизации.

**:id** - та самая короткая RAND_STRING, которая и является короткой ссылкой.

Ответ должен содержать JSON из структуры
```
type LinkInfo struct {
	ShortURL          string `json:"shorturl"`
	LongURL          string `json:"longurl"`
	NumberOfClicks	int `json:"numberofclicks"`
}
```

Еще возможны ответы: 401 Unauthorized, StatusNotFound


### удаление короткой ссылки пользователя

**http://HOST_NAME/api/links/:id**

Метод: **DELETE**

Требует авторизации.

**:id** - та самая короткая RAND_STRING, которая и является короткой ссылкой.

Если удаление прошло успешно, мы получим StatusOK.

Еще возможны ответы: 401 Unauthorized, StatusNotFound



## Статистика по ссылкам


### получение временного графика количества переходов с группировкой по дням,
часам, минутам. По всем ссылкам данного пользователя.

**http://HOST_NAME/api/stat/interval/:intervalLetter**

Метод: **GET**

Требует авторизации.

intervalLetter:
    d - группировка по дням,
    h - по часам
    i - по минутам

В ответе будет json (от **Map**) со множеством пар Интервал-Количество переходов.

Еще возможны ответы: 401 Unauthorized

//TODO сделать еще статистику для каждой конкретной ссылки


### получение топа из 20 сайтов иcточников переходов

**http://HOST_NAME/api/stat/topref**

Метод: **GET**

Требует авторизации.

В ответе будет json-массив с адресами топа из 20 сайтов иcточников переходов (возможно, пустой).

**Внимание! Клиентское приложение должно присылать заголовок Referer, иначе переходы не будут учтены в "топ20 сайтов иcточников переходов"**

Еще возможны ответы: 401 Unauthorized




## Ссылки


### Переход по короткой ссылке

**http://HOST_NAME/RAND_STRING**

Метод: **GET**

Например:
```
    http://localhost:3000/mjqwrrakt
```

Не требует авторизации.

Возможны ответы (кроме нормального тут StatusPermanentRedirect): StatusNotFound






Примечание: теоретически, любой метод может вернуть StatusInternalServerError




 //TODO
 - Сделать нормальную авторизацию,
 - типизировать ошибки

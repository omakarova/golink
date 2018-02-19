Сокращатель ссылок - сервис, который позволяет пользователю создавать более
короткие адреса, которые лучше передавать другим пользователям и собирает
статистику по совершенным переходам. 

Так как программа является учебно-тестовой, а не промышленной, сделан ряд допущений:
- используется HTTP Basic Authorization (не требуется для регистрации нового пользователя и перехода по короткой ссылке)
- пароль передается в теле POST запроса в открытом виде. И в открытом виде хранится в базе.
- я не проверяю имя пользователя и пароль на предмет "плохих" символов. Надеюсь, что сервис-клиент или фронтенд позаботились о том, что имя пользователя и пароль набраны латиницей.

## Собрать и запустить

###База и таблицы

- создаем MySQL базу ***ИМЯ_БАЗЫ*** и назначаем себе привилегии на чтение-запись-создание таблиц.
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
  ```<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
  ```

### Сборка и запуск
 Берем ветку master и собираем исполняемый файл в IDE или из командной строки.
 Например, так:
      ```go build ./project/myLinkMain.go
      ```
 Потом запустить исполняемый файл.

 По умолчанию сервер стартует на порте 3000.
 
 При запуске под windows могут понадобиться разрешения на брандмауэре.
# DOCKER lean

## 1. Команда для запуска контейнера БД

    $ docker run --name db -dp 5432:5432 -e POSTGRES_PASSWORD='qwerty' --rm -v roach:/var/lib/postgresql/data --network mynet postgres

--name db установит нам хост 'db'

Создать таблицу в postgres: 

    $ docker exec -it db /bin/bash
    $ psql -U postgres
    $ create table cars (model varchar(255), price int);


## 2. Dockerfile.multi - образ мгногоэтапной сборки.
Команда для создания образа: 

    $ docker build -f Dockerfile.multi -t <name_image> .  

Команда запуска контейнера: 

    $ docker run -it --rm -dp 8080:8080 --network mynet --name apidb -e PORT=':8080' -e DBHOST=db -e DBPORT=5432 -e DBUSER=postgres -e DBNAME=postgres -e DBPASSWORD=qwerty  <name_image>

gcr.io/distroless/base-debian11 - образ докера без дистрибутива


## 3. Docker-compose
docker-compose.yml 
В качестве сборки берется Dockerfile.multi, а не Dockerfile
    $ build:
        context: .
            dockerfile: Dockerfile.multi

Запуск зависит от контейнера db

    $ depends_on:
              - db

Значения среды окружения берутся из файла .env

       $ env_file: .env
     

Команда:

    $ docker compose up

Команды docker compose: https://docs.docker.com/compose/compose-file/

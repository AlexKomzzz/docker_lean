# DOCKER lean

## 1. Команда для запуска контейнера БД
docker run --name db -dp 5432:5432 -e POSTGRES_PASSWORD='qwerty' --rm -v roach:/var/lib/postgresql/data --network mynet postgres

--name db установит нам хост 'db'

    войти в оболочку postgres: 
docker exec -it db /bin/bash
psql -U postgres
create table cars (model varchar(255), price int);


## 2. Dockerfile.multi - образ мгногоэтапной сборки.
команда для создания образа: docker build -f Dockerfile.multi -t <name_image> .  

команда запуска контейнера: docker run -it --rm -dp 8080:8080 --network mynet --name apidb -e PORT=':8080' -e DBHOST=db -e DBPORT=5432 -e DBUSER=postgres -e DBNAME=postgres -e DBPASSWORD=qwerty  <name_image>

gcr.io/distroless/base-debian11 - образ докера без дистрибутива


## 3. Docker-compose

команды docker compose: https://docs.docker.com/compose/compose-file/

использовать environment или env_file:
  - ./a.env
  - ./b.env

  build использовать как
    - build: ./                (тогда для контейнера будет использоваться Dockerfile)
    - build:                (а здесь для контейнера будет использоваться Dockerfile.multi)
        context: .
            dockerfile: Dockerfile.multi
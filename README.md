# DOCKER lean


## Dockerfile.multi - образ мгногоэтапной сборки.
команда для создания образа: docker build -f Dockerfile.multi -t <name_image> .  

команда запуска контейнера: docker run -it --rm -dp 8080:8080 --network mynet --name apidb -e PORT=':8080' -e DBHOST=db -e DBPORT=5432 -e DBUSER=postgres -e DBNAME=postgres -e DBPASSWORD=qwerty  <name_image>

gcr.io/distroless/base-debian11 - образ докера без дистрибутива

## Команда для запуска контейнера БД
docker run --name db -dp 5432:5432 -e POSTGRES_PASSWORD='qwerty' --rm -v roach:/var/lib/postgresql/data --network mynet postgres

--name db установит нам хост 'db'
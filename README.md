# DOCKER lean


## Dockerfile.multi - образ мгногоэтапной сборки.
команда для создания образа: docker build -f Dockerfile.multi -t <name_image> .    
 
gcr.io/distroless/base-debian11 - образ докера без дистрибутива

## Команда для запуска контейнера БД
docker run --name db -dp 5432:5432 -e POSTGRES_PASSWORD='qwerty' --rm -v roach:/var/lib/postgresql/data --hostname db --network mynet postgres

# To migrate db
    migrate -database ${POSTGRESQL_URL} -path db/migrations up
# To generate graph-folder
    go run github.com/99designs/gqlgen@v0.17.39 generate

# To run postgres container
    docker run --rm -P -p 127.0.0.1:5432:5432 -e POSTGRES_PASSWORD="password" --name pg postgres:alpine
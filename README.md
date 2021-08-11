Forked from https://github.com/juancurti/go_tdd_tutorial/

# Running Postgres on your local machine
Using Docker is both simple and convenient. 

    docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=12345 postgres

Optionally if you need to connect to your Docker Postgre instance use the following line: 

    sudo docker exec -it my-postgres bash
    psql -U postgres

# Before you start
In order to run this application you need to define 3 env variables such as

    export APP_DB_USERNAME=postgres
    export APP_DB_PASSWORD=12345
    export APP_DB_NAME=postgres


# Running the Tests
In order to run the tests just use:

    go test -v

.
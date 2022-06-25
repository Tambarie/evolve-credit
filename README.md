

# Evolve-Credit API

A Golang `RESTful-API` that returns a dataset of your choosing (e.g `users`).
You are required to set up a Postgres database on any free platform you are familiar with.


## Using the API
To use API, clone this repository and follow one of two methods shown below:


### Using Your Local Machine

1. [Install Go](https://golang.org/doc/install)
2. [Install Postgres DB](https://www.postgresql.org/download/)
3. Create a database named `evolve-credit`.
4. Install all dependencies using:

```
go mod tidy
```
5. Run the application using:

```
1. Using `modd`
2. `go run main.go`
3. `make run`
```


## Setting Environmental Variables
An environment variable is a text file containing ``KEY=value`` pairs of your secret keys and other private information. For security purposes, it is ignored using ``.gitignore`` and not committed with the rest of your codebase.

To create, ensure you are in the root directory of the project then on your terminal type:
```
touch .env
```
All the variables used within the project can now be added within the ``.env`` file in the following format:
```
SERVICE_MODE="dev"
DB_TYPE = "mongodb"
MONGO_DB_HOST="localhost"
MONGO_DB_NAME="payment-gateway"
MONGO_DB_PORT="27017"
SERVICE_PORT=8080
DB_PASS=<your db password>
```


## Author
* Emmanuel Gbaragbo ([Tambarie](https://github.com/Tambarie)) 🐛





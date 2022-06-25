# Evolve-Credit API
A Golang RESTful API that returns a dataset of your choosing (e.g users). You are required to set up a Postgres database on any free platform you are familiar with

## Using the API
To use API, clone this repository and follow one of two methods shown below:

## Using Your Local Machine
Install Go in your machine
Install Postgres Workbench
Create a database named `evolve-credit`.
Install all dependencies using:
go mod tidy
Run the application using:
make run
Setting Environmental Variables
An environment variable is a text file containing KEY=value pairs of your secret keys and other private information. For security purposes, it is ignored using .gitignore and not committed with the rest of your codebase.

To create, ensure you are in the root directory of the project then on your terminal type:

touch `example.env`
All the variables used within the project can now be added within the .env file in the following format:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_NAME=evolve-credit
DB_PASS=******
DB_TIMEZONE=Africa/lagos
DB_MODE=disable
PORT=8080

## Starting the App
1. Using `modd`
2. `go run main.go`
3. `make run`

Author
Emmanuel Gbaragbo (Tambarie) 🐛
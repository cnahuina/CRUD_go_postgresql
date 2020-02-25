# CRUD_go_postgresql
crud with go and postgres + library glide

brew install glide

glide init
glide get "github.com/lib/pq"

======================================
Send params connexion & executed
======================================
POSTGRES_HOST=localhost POSTGRES_USER=postgres POSTGRES_PASSWORD=123456 POSTGRES_DB_NAME=dbApi go run main.go

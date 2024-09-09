# 04-Go
A small rest application for an internet store written in Go in echo framework, using gorm to handle the SQLite database.

To run the app, after cloning the repository please create a `store.db` file with `create.sql` script:

`sqlite3 store.db < create.sql`

After that run the Go compilation:

`go run ./main.go`



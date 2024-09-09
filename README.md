# 04-Go
A small rest application for an internet store written in Go in echo framework, using gorm to handle the SQLite database.

To run the app, after cloning the repository please create a `store.db` file with `create.sql` script:

`sqlite3 store.db < create.sql`

After that run the Go compilation:

`go run ./main.go`

The application should be accessible on `localhost:8080`

Gzipped MP4 demo is here:: https://github.com/AndPust/04-Go/blob/main/demos/04-Go.mp4.gz
 
✅ 4.0 requirement: https://github.com/AndPust/04-Go/commit/4c9f951f37ed293a10c627c0eb797bd9a2ed5fdc 
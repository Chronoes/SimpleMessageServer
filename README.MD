# SimpleMessageServer
## Implementation for TUT Automated Testing course.

### Code standards and tools
- Language: Go
- Editor: Sublime (GoSublime)
- Git repo: GitHub
- Database: SQLite (github.com/mattn/go-sqlite3)

Coding standards will be used as defined in https://golang.org/doc/code.html and https://golang.org/doc/effective_go.html
Code formatting will mostly be done by Gofmt packaged with Go binaries, with the help of GoSublime.

In addition, the following will be used for naming:
- Any major function name (such as a handler) is done with MixedCaps.
- Any minor non-struct function name is done with mixedCaps.
- All user-defined structs and their methods must be in a separate file.

### Architecture
The architecture is defined simply as a server using a database to store information.
The server will accept requests, relay the data to ORM which in turn will communicate with the database.
Requested data will then be sent back to the client.

All eventually possible requests are currently defined in "requests" file.

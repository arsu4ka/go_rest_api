# TODO:
- [] Implement middleware components (will be used to display pages that require authorization)
- [] Split project into two branches, the main one will have *SQLite* db implementation, the secondary one will be based on *PostgreSQL* (as it is right now)
- [] Try to get rid of gorilla packages as they are not maintained anymore. Replace [mux](https://github.com/gorilla/mux) with a [GIN](https://github.com/gin-gonic/gin) or create a custom router. It'd also be better to find alternatives for [gorilla/sessions](https://github.com/gorilla/sessions)
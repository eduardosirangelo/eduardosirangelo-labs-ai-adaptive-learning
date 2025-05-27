# Content Engine

This microservice in Go is responsible for:
- Generating and delivering adaptive content to students.
- Exposing REST endpoints (e.g., `GET /content/{studentId}`) that return personalized lessons.
- Integrating with the RAG/LLM mechanism to search and compose material.

## Structure

```
services/content-engine/
├── cmd/
│   └── content-engine/      # main.go and server initialization
│
├── internal/
│   ├── domain/              # entity definitions (structs)
│   ├── service/             # use cases / business rules
│   └── repository/          # data access interface and implementation
│
├── pkg/                     # helpers or libs that can be exported
│
└── go.mod
```

## How to run

```bash
  cd services/content-engine
```
```bash
  docker build -t content-engine .
```
```bash
  docker run --network aal_engine_network content-engine
```

## Next steps
- Add `PostgreSQL` schema migrations.
- Implement request caching using `Redis`.


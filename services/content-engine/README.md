# Content Engine

This microservice in Go is responsible for:
- Generating and delivering adaptive content to students.
- Exposing REST endpoints (e.g., `GET /content/{studentId}`) that return personalized lessons.
- Integrating with the RAG/LLM mechanism to search and compose material.

## Structure
- `cmd/` – application entry point

- `internal/` – business logic and repositories
- `pkg/` – reusable utilities and helpers

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


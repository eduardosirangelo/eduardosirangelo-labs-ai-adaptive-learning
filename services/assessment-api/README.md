# Assessment API

Microservice in Go that:
- Receives assessment request events via Kafka (`assessment.requested`).
- Processes scoring logic and sends completion events (`assessment.completed`).
- Stores results in PostgreSQL.

## Structure
- `api/` – HTTP handlers and request/response bindings
- `domain/` – entities and business rules
- `infrastructure/` – Kafka clients and database

## How to run

```bash
    cd services/assessment-api
```
```bash
    docker build -t assessment-api .
```
```bash
  docker run --network aal_engine_network assessment-api
```

## Attention points
- Validate idempotency when processing `Kafka events`.
- Write tests that simulate producer and consumer.

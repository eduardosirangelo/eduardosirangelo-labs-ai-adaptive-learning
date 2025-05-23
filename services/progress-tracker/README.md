# Progress Tracker

Responsible for:
- Listening to progress events (via Kafka or Redis Streams).
- Providing a WebSocket channel for the front-end to receive real-time updates.
- Persisting progress checkpoints in Redis and/or PostgreSQL.

## Structure
- `ws/` – WebSocket server in Go
- `stream/` – Kafka/Redis Stream consumers
- `models/` – representation of progress events

## How to run

```bash
  cd services/progress-tracker
```
```bash
  docker build -t progress-tracker .
```
```bash
  docker run --network aal_engine_network progress-tracker
```

## Next steps
- Adjust automatic reconnection on the `WebSocket` client.
- Monitor message latency via `Prometheus`.

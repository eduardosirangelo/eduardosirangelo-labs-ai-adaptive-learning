# Arquitetura do Sistema / System Architecture

### 📦 Estrutura do Repositório / Repository Structure

```plaintext
ai-adaptive-learning/
├── go.work
├── pkg/
│   ├── config/
│   ├── http/
│   ├── auth/
│   └── logging/
├── services/
│   ├── auth-service/       ← Authorization Server (OAuth2/OIDC)
│   ├── content-engine/     ← Conteúdo adaptativo (LLM/RAG)
│   ├── assessment-api/     ← Avaliação via Kafka
│   └── progress-tracker/   ← Progresso em tempo real (WebSocket)
├── ui/
│   ├── web-app/            ← Shell MFE Angular
│   └── mfe-dashboard/      ← Dashboard micro-frontend
├── infrastructure/
│   ├── terraform/
│   └── k8s-manifests/
├── docs/
│   ├── ARCHITECTURE.md
│   └── STUDY-PLAN.md
└── .github/ /workflows/
```

---

## 🎓 AI Adaptive Learning Engine

> Plataforma de microserviços e micro-frontends para educação adaptativa, altamente escalável, em nuvem, com personalização em tempo real e integração LMS.

---

### 🏗️ Arquitetura / Architecture

Nossa plataforma segue os princípios de **Clean Architecture**, dividindo o software em camadas concêntricas e separando regras de negócio de detalhes de infraestrutura:

* **Domain (Entidades)**: Modelos puros de negócio (`Student`, `Content`, `AssessmentResult`) sem dependências externas.
* **Use Cases (Casos de Uso)**: Lógica de aplicação como `GenerateAdaptivePath`, `ScoreAssessment`, `RecordProgress`, orquestrando entidades e serviços.
* **Interface Adapters (Adaptadores)**: Controllers REST/gRPC, handlers WebSocket, repositórios Kafka/DB que convertem dados entre casos de uso e infraestrutura.
* **Frameworks & Drivers (Infra)**: Servidor HTTP (Chi), Kafka, PostgreSQL, Redis, Angular MFE, Kubernetes/Terraform.

As dependências fluem **sempre para dentro**, garantindo que mudanças na infraestrutura ou UI não afetem as regras de negócio centrais.

---

### 🎯 Pré-requisitos / Prerequisites

* **Go 1.21+**
* **Node.js 20+ & Angular CLI**
* **Docker & Docker Compose**
* **Terraform 1.8+**
* **Kafka 3.6+**
* **PostgreSQL**

Consulte as versões e instruções de instalação em `docs/` e nos README.md de cada serviço.

---

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

### 🔗 Links Rápidos / Quick Links

* **Services**

  * [Auth Service](services/auth/README.md)
  * [Content Engine](services/content-engine/README.md)
  * [Assessment API](services/assessment-api/README.md)
  * [Progress Tracker](services/progress-tracker/README.md)
* **UI**

  * [Web-App Shell](ui/web-app/README.md)
  * [MFE Dashboard](ui/mfe-dashboard/README.md)
* **Infrastructure**

  * [Terraform](infrastructure/terraform/README.md)
  * [K8s Manifests](infrastructure/k8s-manifests/README.md)
* **Docs**

  * [Arquitetura Detalhada](docs/ARCHITECTURE.md)
  * [Plano de Estudos](docs/STUDY-PLAN.md)

---

### 📚 Material de Estudo / Study Material

Acesse nosso quadro público no Trello com recursos complementares:
[https://trello.com/b/SEU\_BOARD\_ID](https://trello.com/b/SEU_BOARD_ID)

---

### 🤝 Contribuindo / Contributing

Veja [CONTRIBUTING.md](/.github/CONTRIBUTING.md) para padrões de pull request, estilo de código e pipeline.

---

### 📄 Licença / License

MIT License — consulte [LICENSE](LICENSE).

---

## English Version

### 🎓 AI Adaptive Learning Engine

> A microservices and micro-frontends platform for adaptive education, scalable, cloud-native, real-time personalization, and LMS integration.

---

#### 🏗️ Architecture

Our platform adheres to **Clean Architecture**: concentric layers isolating business rules from infrastructure details:

* **Domain**: Pure business models (`Student`, `Content`, `AssessmentResult`).
* **Use Cases**: Application logic like `GenerateAdaptivePath`, `ScoreAssessment`, `RecordProgress`.
* **Interface Adapters**: REST/gRPC controllers, WebSocket handlers, Kafka/DB repositories.
* **Frameworks & Drivers**: HTTP server (Chi), Kafka, PostgreSQL, Redis, Angular MFE, Kubernetes/Terraform.

Dependencies always point inward, so infrastructure or UI changes don’t affect core business rules.

---

#### 🎯 Prerequisites

* **Go 1.21+**
* **Node.js 20+ & Angular CLI**
* **Docker & Docker Compose**
* **Terraform 1.8+**
* **Kafka 3.6+**
* **PostgreSQL**

See installation guides in `docs/` and each service README.

---

#### 📦 Repository Structure

```plaintext
ai-adaptive-learning/
├── go.work
├── pkg/
│   ├── config/
│   ├── http/
│   ├── auth/
│   └── logging/
├── services/
│   ├── auth-service/
│   ├── content-engine/
│   ├── assessment-api/
│   └── progress-tracker/
├── ui/
│   ├── web-app/
│   └── mfe-dashboard/
├── infrastructure/
│   ├── terraform/
│   └── k8s-manifests/
├── docs/
└── .github/workflows/
```

---

#### 🔗 Quick Links

* **Services**

  * [Auth Service](services/auth-service/README.md)
  * [Content Engine](services/content-engine/README.md)
  * [Assessment API](services/assessment-api/README.md)
  * [Progress Tracker](services/progress-tracker/README.md)
* **UI**

  * [Web-App Shell](ui/web-app/README.md)
  * [MFE Dashboard](ui/mfe-dashboard/README.md)
* **Infrastructure**

  * [Terraform](infrastructure/terraform/README.md)
  * [K8s Manifests](infrastructure/k8s-manifests/README.md)
* **Docs**

  * [Detailed Architecture](docs/ARCHITECTURE.md)
  * [Study Plan](docs/STUDY-PLAN.md)

---

#### 📚 Study Materials

Public Trello board with supplementary resources:
[https://trello.com/b/SEU\_BOARD\_ID](https://trello.com/b/SEU_BOARD_ID)

---

#### 🤝 Contributing

See [CONTRIBUTING.md](/.github/CONTRIBUTING.md) for PR guidelines and coding standards.

---

#### 📄 License

MIT License — see [LICENSE](LICENSE).

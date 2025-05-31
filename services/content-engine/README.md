# Content Engine

Este microserviço é o coração do módulo de conteúdo adaptativo da plataforma. Aqui, oferecemos lições personalizadas e dinâmicas aos estudantes, baseado em algoritmos de AI (RAG e LLM). A seguir, descrevemos de forma didática os **próximos passos de desenvolvimento**, explicando **o que** será implementado e **por que** utilizamos cada padrão de arquitetura e design.

## 🚀 Objetivos de Desenvolvimento

1. **Modelagem de Domínio (Entities)**

    * *O que será feito:* Definir modelos puros como `Student`, `ContentItem`, `LearningPath` e `AssessmentResult`.
    * *Por que:* Isolar regras de negócio fundamentais sem acoplá-las a frameworks ou infraestrutura, seguindo o princípio de manter o núcleo da aplicação limpo e testável (Clean Architecture).

2. **Casos de Uso (Use Cases)**

    * *O que será feito:* Implementar métodos como `GenerateAdaptivePath(studentID)` e `GetContentByID` que orquestram a lógica central.
    * *Por que:* Concentrar a lógica de aplicação em uma camada específica torna o código mais organizado, facilita testes e respeita o Single Responsibility Principle.

3. **Interfaces de Repositório (Ports)**

    * *O que será feito:* Criar interfaces como `ContentRepository` com métodos `Save`, `FindByID`, `ListByCriteria`.
    * *Por que:* O Repository Pattern desacopla a lógica de persistência dos casos de uso, permitindo trocar o banco (PostgreSQL, mocks em testes) sem alterar o núcleo.

4. **Adaptadores de Persistência (Adapters)**

    * *O que será feito:* Implementar `PostgresContentRepository` que cumpre a interface acima, usando `database/sql` ou ORM leve.
    * *Por que:* Seguir o Ports & Adapters (Hexagonal) garante que detalhes do banco não vazem para o domínio.

5. **Camada de Serviço (Service Layer)**

    * *O que será feito:* Criar `ContentService` que recebe as dependências via construtor e aplica a lógica de negócio (filtragem, regras adaptativas).
    * *Por que:* O Service Layer Pattern organiza casos de uso de forma coesa e facilita a injeção de dependências para testes.

6. **Integração com LLM/RAG (Adapter + Strategy)**

    * *O que será feito:* Definir interface `LLMClient` e implementar estratégias (por ex., `OpenAIClient`, `LocalRAGClient`).
    * *Por que:* Adapter Pattern desacopla a API externa, e Strategy Pattern permite alternar algoritmos de geração de conteúdo dinamicamente.

7. **Event-Driven com Kafka**

    * *O que será feito:* Configurar produtores/consumidores para eventos `ContentRequested` e `ContentDelivered`.
    * *Por que:* A arquitetura orientada a eventos desacopla componentes e melhora escalabilidade e resiliência.

8. **Exposição de API (REST + gRPC)**

    * *O que será feito:* Usar Chi e middlewares comuns (autenticação, logging, CORS) para expor endpoints HTTP e gRPC.
    * *Por que:* Oferecer múltiplas interfaces de consumo e manter consistência de cross-cutting concerns via pkg/http e pkg/auth.

9. **Resiliência (Circuit Breaker & Retry)**

    * *O que será feito:* Aplicar políticas de retry e circuit breaker nas chamadas externas (LLM, DB).
    * *Por que:* Prevenir falhas em cascata e garantir robustez em produção.

10. **Testes Automatizados**

    * *O que será feito:* Escrever testes de unidade nas camadas Domain e Service, e testes de integração para adaptadores.
    * *Por que:* Validar regras de negócio isoladamente e garantir que a aplicação funcione conforme esperado em diferentes cenários.

## 📚 Por que esta abordagem?

Toda a evolução do Content Engine segue os princípios da **Clean Architecture** e dos **Design Patterns** selecionados para:

* **Separar responsabilidades** (Single Responsibility Principle)
* **Isolar regras de negócio** de frameworks e infraestrutura
* **Facilitar testes** unitários e de integração
* **Promover reuso** de código e padrões consistentes entre microserviços
* **Garantir escalabilidade e manutenção** a longo prazo

Com este roadmap, o estudante compreenderá não apenas **o que** implementar, mas **por que** cada decisão arquitetural e de design pattern foi tomada, preparando-o para construir sistemas sofisticados e profissionais.

### 📦 Estrutura do Repositório

```plaintext
services/content-engine/
├── cmd/
│   └── content-engine/      # main.go inicialização do serviço
│
├── internal/
│   ├── domain/              # definições de entidades (structs)
│   ├── service/             # casos de uso / regras de negócio
│   └── repository/          # interface e implementação de acesso a dados
│
├── pkg/                     # helpers ou libs que podem ser exportadas
│
└── go.mod
```

## 📖 Como usar este serviço

Para testar e explorar o Content Engine:

1. **Clone o repositório e sincronize o workspace Go**

   ```bash
   git clone git@github.com:eduardosirangelo/ai-adaptive-learning.git
   cd ai-adaptive-learning
   go work sync
   ```
2. **Configure variáveis de ambiente (exemplo local)**

   ```bash
   export PORT=8070
   export DATABASE_URL="postgres://user:pass@localhost:5432/contentdb"
   export KAFKA_BROKERS=localhost:9092
   ```
3. **Ingressar no serviço**

   ```bash
   cd services/content-engine
   go run cmd/content-engine/main.go
   ```
4. **Verificar health check**

   ```bash
   curl http://localhost:8070/health
   # deve retornar "OK"
   ```

## 🔗 Documentação e Coleção Postman / gRPC & Queue Tests

* **API REST (Postman Collection)**: importe esta coleção pública no Postman para visualizar e testar todos os endpoints REST do Content Engine:
  [https://www.postman.com/eduardosirangelo/workspace/ai-adaptive-learning-engine/collection/CONTENT-ENGINE](https://www.postman.com/eduardosirangelo/workspace/ai-adaptive-learning-engine/collection/CONTENT-ENGINE)

* **gRPC & Event-Driven (Queue) Tests**:

    * Use `grpcurl` ou Postman gRPC para invocar métodos gRPC definidos em `services/content-engine/api.proto`.
    * Consuma e produza eventos Kafka usando scripts em `services/content-engine/tools/kafka-cli.sh`, verificando tópicos `content.requested` e `content.delivered`.

---

## English Version

# Content Engine

This microservice is the heart of the platform's adaptive content module. It delivers personalized, dynamic lessons to students using AI algorithms (RAG and LLM). Below we describe in a didactic way the **next development steps**, explaining **what** will be implemented and **why** each architectural and design pattern was chosen.

## 🚀 Development Goals

1. **Domain Modeling (Entities)**

    * *What will be done:* Define pure domain models such as `Student`, `ContentItem`, `LearningPath`, and `AssessmentResult`.
    * *Why:* To isolate core business rules from frameworks or infrastructure, keeping the application core clean and testable (Clean Architecture).

2. **Use Cases**

    * *What will be done:* Implement methods like `GenerateAdaptivePath(studentID)` and `GetContentByID` that orchestrate core logic.
    * *Why:* Concentrating application logic in a specific layer makes code more organized, easier to test, and aligns with the Single Responsibility Principle.

3. **Repository Interfaces (Ports)**

    * *What will be done:* Create interfaces like `ContentRepository` with methods `Save`, `FindByID`, `ListByCriteria`.
    * *Why:* The Repository Pattern decouples persistence logic from use cases, allowing swapping out the database (PostgreSQL, mock in tests) without modifying the core logic.

4. **Persistence Adapters**

    * *What will be done:* Implement `PostgresContentRepository` fulfilling the above interface, using `database/sql` or a lightweight ORM.
    * *Why:* Following Ports & Adapters (Hexagonal) ensures that database details do not leak into the domain layer.

5. **Service Layer**

    * *What will be done:* Create `ContentService` that receives dependencies via constructor and applies business logic (filtering, adaptive rules).
    * *Why:* The Service Layer Pattern organizes use cases cohesively and facilitates dependency injection for testing.

6. **LLM/RAG Integration (Adapter + Strategy)**

    * *What will be done:* Define an `LLMClient` interface and implement strategies (e.g., `OpenAIClient`, `LocalRAGClient`).
    * *Why:* Adapter Pattern decouples external API details, and Strategy Pattern allows dynamic switching between content generation algorithms.

7. **Event-Driven with Kafka**

    * *What will be done:* Configure producers/consumers for `ContentRequested` and `ContentDelivered` events.
    * *Why:* Event-driven architecture decouples components and improves scalability and resilience.

8. **API Exposure (REST + gRPC)**

    * *What will be done:* Use Chi and shared middlewares (authentication, logging, CORS) to expose HTTP and gRPC endpoints.
    * *Why:* Provide multiple consumption interfaces and maintain consistency of cross-cutting concerns via pkg/http and pkg/auth.

9. **Resilience (Circuit Breaker & Retry)**

    * *What will be done:* Apply retry and circuit breaker policies on external calls (LLM, DB).
    * *Why:* Prevent cascading failures and ensure robustness in production.

10. **Automated Testing**

    * *What will be done:* Write unit tests for Domain and Service layers and integration tests for adapters.
    * *Why:* Validate business rules in isolation and ensure the application behaves as expected in various scenarios.

## 📚 Why This Approach?

The Content Engine's evolution follows **Clean Architecture** principles and carefully selected **Design Patterns** to:

* **Separate responsibilities** (Single Responsibility Principle)
* **Isolate business rules** from frameworks and infrastructure
* **Facilitate** unit and integration testing
* **Promote code reuse** and consistent patterns across microservices
* **Ensure long-term scalability and maintainability**

This roadmap helps learners grasp not just **what** to implement, but **why** each architectural and design decision was made, preparing them to build sophisticated, professional systems.

### 📦 Repository Structure

```plaintext
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

## 📖 How to Use this Service

To test and explore the Content Engine:

1. **Clone the repository and sync the Go workspace**

   ```bash
   git clone git@github.com:eduardosirangelo/ai-adaptive-learning.git
   cd ai-adaptive-learning
   go work sync
   ```
2. **Set environment variables (local example)**

   ```bash
   export PORT=8070
   export DATABASE_URL="postgres://user:pass@localhost:5432/contentdb"
   export KAFKA_BROKERS=localhost:9092
   ```
3. **Run the service**

   ```bash
   cd services/content-engine
   go run cmd/content-engine/main.go
   ```
4. **Check health endpoint**

   ```bash
   curl http://localhost:8070/health
   # should return "OK"
   ```

## 🔗 Documentation & Test Collections

* **REST API (Postman Collection)**: Import the public collection in Postman to view and test all Content Engine REST endpoints:
  [https://www.postman.com/eduardosirangelo/workspace/ai-adaptive-learning-engine/collection/CONTENT-ENGINE](https://www.postman.com/eduardosirangelo/workspace/ai-adaptive-learning-engine/collection/CONTENT-ENGINE)

* **gRPC & Event-Driven (Queue) Tests**:

    * Use `grpcurl` or Postman gRPC to invoke methods defined in `services/content-engine/api.proto`.
    * Produce and consume Kafka events using scripts in `services/content-engine/tools/kafka-cli.sh`, monitoring topics `content.requested` and `content.delivered`.


# Content Engine

This microservice in Go is responsible for:
- Generating and delivering adaptive content to students.
- Exposing REST endpoints (e.g., `GET /content/{studentId}`) that return personalized lessons.
- Integrating with the RAG/LLM mechanism to search and compose material.



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

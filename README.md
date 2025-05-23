# 🎓 AI Adaptive Learning Engine

[![Go](https://img.shields.io/badge/Go-1.21-00ADD8?logo=go)](https://golang.org)
[![Angular](https://img.shields.io/badge/Angular-18-DD0031?logo=angular)](https://angular.io)
[![Kafka](https://img.shields.io/badge/Kafka-3.6-231F20?logo=apachekafka)](https://kafka.apache.org/)
[![Terraform](https://img.shields.io/badge/Terraform-1.8-7B42BC?logo=terraform)](https://www.terraform.io/)
[![CI/CD](https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/actions/workflows/ci.yml/badge.svg)](https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/actions)

> **A scalable, cloud-native platform for adaptive and personalized education, inspired by the best practices of EdTech leaders like McGraw-Hill.**

---

## 🚀 Overview

AI Adaptive Learning Engine is a modern, microservices-based platform designed for educational technology solutions that demand high scalability, real-time personalization, and seamless integration with learning management systems (LMS).

- **Built with:** Go (microservices), Angular (micro frontends), Kafka (event streaming), PostgreSQL (data), Terraform (IaC), Docker/Kubernetes (cloud native).
- **Key Features:**  
  - Adaptive content delivery using AI agents (RAG, LLM integration)
  - Real-time student progress tracking (WebSocket)
  - Modular, scalable architecture (microservices + micro frontends)
  - LMS compatibility (LTI 1.3, xAPI)
  - Robust DevOps pipelines (CI/CD, IaC, monitoring)

---

## 🏗️ Architecture

```
                   +-------------------+
                   |   Angular MFE     |
                   +-------------------+
                             |
                             v
+-------------------+    Kafka Events    +---------------------+
|  Content Engine   |                    |  Assessment Engine  |
|   (Go Service)    |                    |   (Go Service)      |
+-------------------+                    +---------------------+
          |                                       |
          v                                       v
         PostgreSQL DB  Progress Tracker (Go + Redis)
```
- **Microservices:** Isolated Go services for content, assessment, and progress.
- **Event-Driven:** Kafka for decoupled, scalable communication.
- **Micro Frontends:** Angular 18+ with Module Federation for independent UI modules.
- **Cloud Native:** Deployable on AWS/GCP using Kubernetes and Terraform.

---

## 📦 Repository Structure

```
ai-adaptive-learning/
│
├── services/
│   ├── content-engine/      # Adaptive content delivery (Go)
│   ├── assessment-api/      # Assessment logic & scoring (Go)
│   └── progress-tracker/    # Real-time progress (Go + Redis)
│
├── web-app/
│   └── mfe-dashboard/       # Angular micro frontend
│
├── infrastructure/
│   ├── terraform/           # IaC for AWS/GCP
│   └── k8s-manifests/       # Kubernetes deployment files
│
├── docs/
│   ├── ARCHITECTURE.md      # Detailed diagrams & flows
│   └── MCGRAW-HILL_INSIGHTS.md  # EdTech research & benchmarking
│
└── .github/
    └── workflows/           # CI/CD pipelines
```

---

## 🌟 Key Features

- **Personalized Learning:**  
  AI-driven adaptive content and assessments, leveraging LLMs and RAG for dynamic student journeys.
- **Real-Time Analytics:**  
  WebSocket-based progress tracker for instant feedback and dashboards.
- **Enterprise-Ready:**  
  Modular microservices, cloud-native deployment, and robust security (JWT, OAuth2).
- **Interoperability:**  
  LTI 1.3 and xAPI support for seamless LMS integration.

---

## 🛠️ Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org)
- [Node.js 20+ & Angular CLI](https://angular.io/cli)
- [Docker & Docker Compose](https://docs.docker.com/)
- [Terraform](https://www.terraform.io/)
- [Kafka](https://kafka.apache.org/)
- [PostgreSQL](https://www.postgresql.org/)

### Quickstart

#### Clone the repository
```bash
  git clone git@github.com:eduardosirangelo/go-angular-ai-adaptive-learning.git
```
```bash
  cd go-angular-ai-adaptive-learning
```

#### Start all services (dev mode)
```bash
  docker-compose up --build
```
#### Run Angular frontend
```bash
  cd web-app/mfe-dashboard
```
```bash
  npm install
```
```bash
  ng serve
```

---

## 🧑‍💻 Contributing

Contributions, suggestions, and feedback are welcome!  
See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

## 📚 Documentation

- [ARCHITECTURE.md](docs/ARCHITECTURE.md): System diagrams, flows, and design decisions.
- [MCGRAW-HILL_INSIGHTS.md](docs/MCGRAW-HILL_INSIGHTS.md): EdTech research and benchmarking.
- [API Reference](services/README.md): Endpoints, events, and integration details.

---

## 🏆 About the Author

**Eduardo Sirangelo**  
Senior Software Engineer | 19+ years of experience  
[LinkedIn](https://linkedin.com/in/eduardosirangelo) | [Email](mailto:eduardo.sirangelo@gmail.com)

- Led digital transformation projects for global EdTech, FinTech, and HealthTech companies.
- Expert in scalable microservices, real-time analytics, and cloud-native architectures.
- Passionate about building impactful technology for education and beyond.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

> “Empowering education through adaptive, scalable, and intelligent technology.”

---

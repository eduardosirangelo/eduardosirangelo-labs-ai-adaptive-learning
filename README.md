<a id="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<div align="center">
  <a href="https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning">
    <img src="docs/images/logo.png" alt="AI Adaptive Learning Engine Logo" width="80" height="80">
  </a>

<h3 align="center">AI Adaptive Learning Engine</h3>

  <p align="center">
    Plataforma de microserviços e micro-frontends para educação adaptativa, altamente escalável, em nuvem, com personalização em tempo real e integração com LMS.
    <br />
    <a href="#about-the-project">Visão Geral</a> ·
    <a href="#getting-started">Começando</a> ·
    <a href="#architecturearquitetura">Arquitetura</a> ·
    <a href="#readme-top">Back to top</a>
  </p>
</div>

---

## Table of Contents

<ol>
  <li><a href="#about-the-project">About The Project / Sobre o Projeto</a></li>
  <li><a href="#built-with">Built With / Tecnologias</a></li>
  <li><a href="#getting-started">Getting Started / Começando</a>
    <ul>
      <li><a href="#prerequisites">Prerequisites / Pré-requisitos</a></li>
      <li><a href="#installation">Installation / Instalação</a></li>
    </ul>
  </li>
  <li><a href="#usage">Usage / Uso</a></li>
  <li><a href="#roadmap">Roadmap</a></li>
  <li><a href="#architecturearquitetura">Architecture / Arquitetura</a></li>
  <li><a href="#contributing">Contributing / Contribuindo</a></li>
  <li><a href="#license">License / Licença</a></li>
  <li><a href="#contact">Contact / Contato</a></li>
  <li><a href="#acknowledgments">Acknowledgments / Agradecimentos</a></li>
</ol>

---

## About The Project / Sobre o Projeto

![Product Screenshot][product-screenshot]

**AI Adaptive Learning Engine** é uma plataforma monorepo composta de microserviços em Go e micro-frontends em Angular para entregar conteúdo educativo adaptativo, com rastreamento de progresso em tempo real, orquestração via Kafka e deploy cloud-native em Kubernetes.
**AI Adaptive Learning Engine** is a monorepo platform of Go microservices and Angular micro-frontends delivering adaptive educational content, real-time progress tracking, Kafka orchestration, and cloud-native Kubernetes deployment.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Built With / Tecnologias

* [Go 1.21+](https://golang.org)
* [Angular 18+](https://angular.io)
* [Kafka 3.6+](https://kafka.apache.org)
* [Terraform 1.8+](https://www.terraform.io)
* [Docker & Docker Compose](https://docs.docker.com)
* [Kubernetes](https://kubernetes.io)
* [PostgreSQL](https://www.postgresql.org)
* [Redis](https://redis.io)
* [Chi Router](https://github.com/go-chi/chi)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Getting Started / Começando

Siga os passos abaixo para rodar o projeto localmente.

### Prerequisites / Pré-requisitos

* Go 1.21+
* Node.js 20+ & Angular CLI
* Docker & Docker Compose
* Terraform 1.8+
* Kafka 3.6+
* PostgreSQL
* Redis

Veja instruções de instalação detalhadas em `docs/` e nos READMEs de cada módulo.

### Installation / Instalação

1. Clone este repositório

   ```bash
   git clone https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning.git
   cd go-angular-ai-adaptive-learning
   ```
2. Sincronize módulos Go

   ```bash
   go work sync
   ```
3. Suba todos os serviços em background

   ```bash
   make docker
   ```
4. Acesse o dashboard Angular

   ```bash
   cd ui/mfe-dashboard
   npm install
   ng serve
   ```
5. Para parar tudo

   ```bash
   make down
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Usage / Uso

* **Content Engine**

    * API REST em `http://localhost:8071/health`
* **Assessment API**

    * Endpoints em `http://localhost:8090/health`
* **Auth Service**

    * Checagem de saúde em `http://localhost:9000/health`
* **Progress Tracker**

    * WebSocket em `ws://localhost:8072/ws`

Para detalhes de cada serviço e coleção Postman, veja os READMEs em `services/*` e o link público na seção Acknowledgments.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Roadmap

* [x] Estrutura básica do monorepo
* [x] Health checks para todos serviços
* [x] Configuração de Docker Compose + Makefile
* [ ] Implementar micro-frontends adicionais
* [ ] Pipeline de CI/CD completo
* [ ] Documentação de gRPC e integração Kafka

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

---

## Architecture / Arquitetura

Nossa plataforma segue os princípios de **Clean Architecture**, com camadas concêntricas onde as dependências fluem sempre para dentro:

1. **Domain (Entidades)**
2. **Use Cases (Casos de Uso)**
3. **Interface Adapters (Adaptadores)**
4. **Frameworks & Drivers (Infraestrutura)**

Para uma visão detalhada, com diagramas UML (Diagrama Geral e complementares) gerados via PlantUML, consulte **[docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)**.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Contributing / Contribuindo

Contribuições são bem-vindas! Veja [`.github/CONTRIBUTING.md`](.github/CONTRIBUTING.md) para:

1. Fork do projeto
2. Crie feature branch (`git checkout -b feature/MinhaFuncionalidade`)
3. Commit e PR
4. Code style: GoFmt, ESLint para Angular
5. Pipeline roda checks automáticos

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## License / Licença

Distributed under the MIT License. See [`LICENSE`](LICENSE) for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Contact / Contato

**Eduardo Sirangelo**

* LinkedIn: [eduardosirangelo](https://linkedin.com/in/eduardosirangelo)
* Email: [eduardo.sirangelo@gmail.com](mailto:eduardo.sirangelo@gmail.com)
* Project: [github.com/eduardosirangelo/go-angular-ai-adaptive-learning](https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Acknowledgments / Agradecimentos

* Quadro público de Trello com material de estudo:
  [https://trello.com/b/SEU\_BOARD\_ID](https://trello.com/b/SEU_BOARD_ID)
* [Best-README-Template](https://github.com/othneildrew/Best-README-Template)
* [Docker Compose](https://docs.docker.com/compose/)
* [Clean Architecture – Uncle Bob](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/eduardosirangelo/go-angular-ai-adaptive-learning.svg?style=for-the-badge
[contributors-url]: https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/eduardosirangelo/go-angular-ai-adaptive-learning.svg?style=for-the-badge
[forks-url]: https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/network/members
[stars-shield]: https://img.shields.io/github/stars/eduardosirangelo/go-angular-ai-adaptive-learning.svg?style=for-the-badge
[stars-url]: https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/stargazers
[issues-shield]: https://img.shields.io/github/issues/eduardosirangelo/go-angular-ai-adaptive-learning.svg?style=for-the-badge
[issues-url]: https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/issues
[license-shield]: https://img.shields.io/github/license/eduardosirangelo/go-angular-ai-adaptive-learning.svg?style=for-the-badge
[license-url]: https://github.com/eduardosirangelo/go-angular-ai-adaptive-learning/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/LinkedIn-EduardoSirangelo-blue?style=for-the-badge&logo=linkedin
[linkedin-url]: https://www.linkedin.com/in/eduardosirangelo/
[product-screenshot]: images/screenshot.png

## Plano de Estudos em 8 Semanas

### Semana 1: Fundamentos e Preparação do Ambiente

* **Objetivos**

    * Configurar IDEs (GoLand, VS Code/WebStorm), Docker, Kafka local, PostgreSQL.
    * Revisar sintaxe e padrões de Go 1.21 e Angular 18.
* **Tarefas (Trello cards)**

    1. Criar card: “Instalar e configurar Go 1.21 e módulos”
    2. Criar card: “Setup Docker Compose com Kafka + Zookeeper + PostgreSQL”
    3. Criar card: “Configurar Angular CLI e projeto base de MFE”
* **Entregáveis**

    * Repositório esqueleto com `docker-compose.yml` e primeiros “Hello World” em Go e Angular.
* **Revisão ativa**

    * No final da semana, resuma em 5 flashcards (Anki) os comandos Docker e as principais diretivas de `go.mod`.

---

### Semana 2: Microserviço Content Engine (Go)

* **Objetivos**

    * Construir serviço Go básico para entrega de conteúdo adaptativo.
    * Introduzir padrões de projeto (Repository, Clean Architecture).
* **Tarefas**

    * Card: “Definir modelo de dados e migrations em PostgreSQL”
    * Card: “Implementar endpoint `/content` com Gin ou Fiber”
    * Card: “Adicionar testes unitários para camada de negócio”
* **Técnicas de fixação**

    * **Pair programming**: reservar 2h para par com colega ou mentor.
    * **Feynman**: escrever um mini-blog explicando o fluxo de requisição no serviço.

---

### Semana 3: Assessment Engine e Eventos Kafka

* **Objetivos**

    * Desenvolver serviço de avaliação, integração com Kafka (produção/consumo de eventos).
    * Entender RAG e LLM na prática.
* **Tarefas**

    * Card: “Implementar producer de evento ‘assessment.requested’”
    * Card: “Criar consumer que calcula nota e emite ‘assessment.completed’”
    * Card: “Mock de LLM para geração de perguntas (RAG)”
* **Revisão ativa**

    * Mapear diagrama de sequência em Miro ou pen paper e anexar ao repositório.

---

### Semana 4: Progress Tracker (Go + Redis + WebSocket)

* **Objetivos**

    * Implementar rastreamento em tempo real via WebSocket e Redis Streams.
* **Tarefas**

    * Card: “Configurar Redis Streams + consumer group”
    * Card: “Criar servidor WebSocket em Go para push de progresso”
    * Card: “Frontend: consumir WebSocket e exibir barras de progresso”
* **Entrega didática**

    * Gravar um screencast de 5 min mostrando o fluxo “Kafka → Go → Redis → WebSocket → Angular”.

---

### Semana 5: Micro Frontends com Angular 18

* **Objetivos**

    * Usar Module Federation para compor dashboards independentes.
* **Tarefas**

    * Card: “Configurar host MFE e remote ‘mfe-dashboard’”
    * Card: “Publicar um componente de gráfico de progresso usando Recharts”
    * Card: “Implementar lazy loading e roteamento dinâmico”
* **Técnicas de fixação**

    * Criar mini-desafio: “Trocar o tema do MFE via input dinâmico” e compartilhar no Slack/Discord.

---

### Semana 6: Infraestrutura como Código (Terraform + Kubernetes)

* **Objetivos**

    * Criar módulos Terraform para provisionar cluster Kubernetes e serviços gerenciados (EKS/GKE).
    * Gerar manifests de Deployment, Service e Ingress.
* **Tarefas**

    * Card: “Escrever módulo Terraform para VPC, EKS/GKE”
    * Card: “Criar Helm Chart ou Kustomize para deployment dos serviços”
    * Card: “Automatizar rollout no pipeline CI/CD”
* **Revisão ativa**

    * Checklist de boas práticas de segurança (RBAC, Secrets) em documento MD.

---

### Semana 7: CI/CD e Monitoramento

* **Objetivos**

    * Configurar GitHub Actions/Azure Pipelines para build, testes e deploy.
    * Integrar Prometheus + Grafana.
* **Tarefas**

    * Card: “Pipeline: build Go, testar, containerizar e push”
    * Card: “Pipeline: build Angular, testes e deploy em ambiente dev”
    * Card: “Configurar alertas básicos no Grafana”
* **Didática**

    * Documentar pipeline em `docs/CI_CD.md` com diagramas.

---

### Semana 8: Integração LMS e Polimento de Portfólio

* **Objetivos**

    * Adicionar suporte LTI 1.3 e xAPI; revisar segurança (OAuth2/JWT).
    * Refatorar README e criar guia “Como ensinar este projeto”.
* **Tarefas**

    * Card: “Implementar endpoint LTI 1.3 de inicialização”
    * Card: “Gerar exemplos de payload xAPI”
    * Card: “Escrever tutorial passo-a-passo no `docs/TUTORIAL.md`”
* **Portfólio**

    * Criar vídeo de walkthrough de 10 min; hospedar no YouTube e linkar no README.

---

## Técnicas de Fixação de Aprendizado

1. **Spaced Repetition** com Anki
2. **Feynman Technique**: escrever/expor cada módulo a um “aluno fictício”
3. **Code Reviews Entre Pares**: buscar feedback externo
4. **Katas e Mini-Desafios**: exercícios de 30 min para reforçar conceitos isolados
5. **Screencasts Rápidos**: gravações de até 5 min explicando fluxos

---

## Abordagens e Ferramentas Alternativas

| Ferramenta              | Como usar                                     | Vantagem                                   |
| ----------------------- | --------------------------------------------- | ------------------------------------------ |
| **Notion/Obsidian**     | Mapear notas, diagramas, Zettelkasten         | Conexão de ideias e links bidirecionais    |
| **GitHub Projects**     | Quadro Kanban integrado ao PR                 | Automatiza status via commits/labels       |
| **Miro/Whimsical**      | Diagramas de arquitetura colaborativos        | Visualização clara em time distribuído     |
| **Anki**                | Revisão espaçada dos termos (Kafka, LTI, RAG) | Retenção de longo prazo                    |
| **Live Coding Streams** | Transmitir sessões ao vivo no Twitch/YouTube  | Feedback em tempo real e portfólio público |

---

Essa estrutura balanceia hands-on, revisão ativa e produção de artefatos didáticos, garantindo que você fixe o conteúdo e construa um portfólio sólido. Pode ajustar ritmo, duração das fases e ferramentas conforme preferência.

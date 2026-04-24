# go-gateway

## Requisitos Funcionais

| ID       | Requisito                                | Descrição                                                                                                                                                                                     |
| -------- | ---------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **RF01** | **Reverse Proxy**                        | O gateway deve interceptar requisições HTTP e encaminhá-las para os microserviços de destino correspondentes.                                                                                 |
| **RF02** | **Roteamento Baseado em Configuração**   | O sistema deve carregar um mapeamento de rotas (ex: `/api/users` → `http://user-service:8080`) a partir de um arquivo de configuração (YAML ou JSON) na inicialização.                        |
| **RF03** | **Validação de JWT (Autenticação)**      | Para rotas marcadas como protegidas, o gateway deve extrair o token do header `Authorization`, validar a assinatura (simétrica ou assimétrica) e checar a expiração (`exp`).                  |
| **RF04** | **Injeção de Headers (Header Mutation)** | Após validar o JWT, o gateway deve extrair claims específicos (ex: `user_id`, `role`) e injetá-los como novos headers na requisição antes de repassá-la ao microserviço interno.              |
| **RF05** | **Rate Limiting**                        | O sistema deve limitar o número de requisições por janela de tempo, identificando o cliente pelo IP (para rotas públicas) ou pelo ID do token (para rotas autenticadas).                      |
| **RF06** | **Distributed Tracing**                  | O gateway deve gerar um identificador único (`X-Request-ID` usando UUID/ULID) caso a requisição não possua um, e repassá-lo ao serviço de destino.                                            |
| **RF07** | **Logging Centralizado**                 | Toda requisição processada deve gerar um log estruturado (JSON) contendo: `X-Request-ID`, método HTTP, path original, target de destino, status code de resposta e latência de processamento. |

# Go Backend Lab

Laboratório de backend em Go, com exemplos pequenos, executáveis e comentados em português. Aprenda o problema que cada técnica resolve e o custo de adotá-la.

## Fundamentos de Go

**[Comece pela trilha de fundamentos](fundamentals/README.md)**: 12 exemplos com explicações `//` no código, quando usar, saída de referência e exercícios. Inclui ponteiros, slices, maps, interfaces, erros, defer, goroutines, channels, select, context.Context, mutex e uma integração concorrente de cotações.

```sh
go run ./fundamentals/07-goroutines
go run ./fundamentals/10-context
go run -race ./fundamentals/11-concurrency
```

## Backend e preparação para entrevistas

**[Guia de backend](docs/backend-interview/README.md)**: investigação de latência, saúde da aplicação, API em camadas, contratos, graceful shutdown, resiliência, mensageria, índices/EXPLAIN, CI/CD e relato de projeto.

O [consumidor idempotente](messaging/idempotency/README.md) demonstra duplicatas concorrentes e traz um roteiro de deduplicação transacional em PostgreSQL.

```sh
go run ./messaging/idempotency
```

## Comece aqui

Requisito: Go 1.22 ou superior. Só biblioteca padrão; não exige banco, Docker ou credenciais.

```sh
go test -race ./...
go run ./cmd/strategy
go run ./cmd/resilience
go run ./cmd/workers
go run ./cmd/api
```

Com a API rodando, em outro terminal: `curl http://localhost:8080/healthz`. Encerre com Ctrl+C para observar o shutdown.

## Trilha de estudos

| Ordem | Exemplo | O que observar | Código |
| --- | --- | --- | --- |
| 1 | Strategy | Trocar Pix por cartão usando uma interface | [payment.go](patterns/strategy/payment.go) |
| 2 | Decorator | Adicionar logs sem mudar o pagamento | [payment.go](patterns/decorator/payment.go) |
| 3 | Retry | Falha transitória, backoff limitado e jitter | [retry.go](resilience/retry.go) |
| 4 | Circuit breaker | Estados fechado, aberto e uma sonda após cooldown | [breaker.go](resilience/breaker.go) |
| 5 | Worker pool | Limitar goroutines e aplicar backpressure | [pool.go](scalability/workers/pool.go) |
| 6 | API HTTP | Timeouts, logs estruturados e graceful shutdown | [main.go](cmd/api/main.go) |

Cada exemplo pode ser estudado isoladamente. Rode os testes e altere um limite para observar o comportamento.

## Boas práticas demonstradas

- Interfaces pequenas no consumidor; dependências explícitas no construtor.
- Dinheiro representado em centavos, sem ponto flutuante.
- Erros com contexto usando `%w` e comparação por `errors.Is`.
- `context.Context` para prazo total e cancelamento cooperativo.
- Concorrência limitada; canais fechados pelo produtor; workers aguardados antes do retorno.
- Logs estruturados sem dados sensíveis de pagamento.
- Testes de cancelamento, concorrência, recuperação do circuito e preservação de erros.
- CI com formatação, análise estática, testes com detector de race e build.

## Quando usar e quais são os custos?

| Técnica | Resolve | Cuidado |
| --- | --- | --- |
| Strategy | Variações de uma regra | Uma interface para cada detalhe torna o código mais difícil de ler |
| Decorator | Comportamento transversal | Muitos wrappers dificultam rastrear a chamada |
| Retry | Falhas transitórias | Amplifica tráfego e pode duplicar efeitos |
| Circuit breaker | Dependência falhando continuamente | Pode rejeitar chamadas que teriam sucesso |
| Worker pool | Recursos locais finitos | Uma fila grande pode esconder sobrecarga |
| API stateless | Permite múltiplas réplicas sem sessão local | Balanceamento e estado compartilhado exigem infraestrutura |

## Limites intencionais

Este é um laboratório, não um SDK de produção. Pix e cartão são simulações.

Retry só deve envolver operações idempotentes, com classificação explícita de erros e prazo total. O contexto não mata uma função que o ignora: o cliente HTTP ou de banco também precisa respeitá-lo. Repetir uma cobrança exige chave de idempotência e persistência da decisão.

O breaker é local ao processo e conta todas as falhas retornadas; em produção, classifique falhas da dependência, separando validação de entrada e cancelamento do cliente. A função passada a `Do` deve ser não nula e retornar normalmente. O decorator exige logger e estratégia não nulos. O construtor de checkout rejeita interface nula; evite passar ponteiros tipados nulos.

Worker pool melhora controle de recursos dentro de uma instância. Escalar horizontalmente exige fila externa, consumidores coordenados, reconhecimento de mensagens e política de reprocessamento. A ordem de conclusão dos jobs é variável.

A API tem somente liveness (`/healthz`), sem dependências. Readiness real deve refletir a capacidade de receber tráfego. Limites do servidor não substituem prazos das operações de negócio.

## Próximos módulos

Itens abaixo são roteiro, ainda não implementados:

- Repository e injeção de dependências com testes de serviço.
- Adapter para provedores externos.
- Rate limiting, bulkhead e cache com TTL.
- Idempotência persistente e transactional outbox.
- Mensageria, retry de consumidor e DLQ.
- Paginação por cursor, métricas e tracing.
- Docker Compose com PostgreSQL e Redis; teste de carga comparando resultados.

## Publicar no GitHub

Crie um repositório vazio chamado `go-backend-lab` na sua conta. Não inicialize com README para evitar históricos diferentes. Depois, nesta pasta:

```sh
git init -b main
git add .
git commit -m "feat: add Go backend learning examples"
git remote add origin https://github.com/SouzaNataliaDev/go-backend-lab.git
git push -u origin main
```

Se você já recebeu a pasta com histórico Git, pule `git init`, `git add` e `git commit`.

## Referências

- [Effective Go](https://go.dev/doc/effective_go)
- [Context](https://pkg.go.dev/context)
- [Testing](https://pkg.go.dev/testing)
- [Race detector](https://go.dev/doc/articles/race_detector)
- [HTTP server](https://pkg.go.dev/net/http#Server)

# Backend: prática e respostas de entrevista

Guia baseado nas anotações da Natália. As respostas abaixo são roteiros de raciocínio: adapte aos fatos que você viveu e às ferramentas disponíveis.

| Tema | Material |
| --- | --- |
| Latência, saúde e bugs em produção | [Observabilidade e diagnóstico](observability.md) |
| API em Go, contratos e encerramento | [Arquitetura e evolução](architecture.md) |
| Falha de dependência, retry, fallback e circuit breaker | [Resiliência](resilience.md) |
| Idempotência, DLQ e offset | [Mensageria](messaging.md) |
| Índice e EXPLAIN | [Banco de dados](database.md) |
| Integração contínua e entrega | [CI/CD](ci-cd.md) |
| Projeto de sucesso | [Relato profissional](success-story.md) |

## Ajustes de vocabulário

| Anotação | Termo e significado |
| --- | --- |
| “ide potência” | **Idempotência**: repetir a mesma operação não repete seu efeito |
| “callback/fallback” | **Fallback** é alternativa controlada; **callback** é função chamada em outro ponto do fluxo |
| “circuit break” | **Circuit breaker**: bloqueia temporariamente chamadas a uma dependência em falha |
| “quando para de aceitar e termina as requisições” | **Graceful shutdown**: tenta concluir trabalho em andamento dentro de um prazo |
| “offset: posição; paginação” | No Kafka, posição dentro de uma **partição**; em SQL, quantidade de linhas ignoradas |
| “compatibilidade retroativa (fallback)” | **Compatibilidade retroativa** mantém consumidores antigos funcionando; fallback atende falhas com alternativa |
| “explain índice” | **Índice** é uma estrutura de acesso; **EXPLAIN** mostra o plano de uma consulta |

Leia também a [trilha de fundamentos de Go](../../fundamentals/README.md). A teoria aqui se conecta aos exemplos executáveis existentes no projeto.

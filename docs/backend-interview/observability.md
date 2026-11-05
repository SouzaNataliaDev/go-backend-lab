# Latência, saúde e diagnóstico de produção

## “Uma API ficou lenta. Como você investiga?”

> Primeiro delimito o impacto: começou quando, afeta quais endpoints, clientes e percentis, e houve deploy ou mudança de configuração? Comparo a janela afetada com o comportamento normal. Correlaciono latência, volume, erros e saturação. Depois sigo o tempo gasto por etapa com traces, métricas e logs. Se a situação estiver causando impacto, avalio uma mitigação segura em paralelo à investigação. Confirmo a hipótese com evidências e verifico se a mitigação recuperou o serviço.

| Sinal | Hipótese | Próxima verificação |
| --- | --- | --- |
| p95/p99 sobem após deploy | Regressão em caminho específico | Comparar traces e versão; avaliar rollback compatível com migrações |
| RPS e filas sobem | Demanda maior que capacidade | Saturação, throughput, fila e limites; avaliar escala e backpressure |
| CPU alta | Trabalho intensivo ou loop | Perfil de CPU, quando autorizado; diferença por endpoint |
| Heap/GC e goroutines crescem | Retenção de memória ou tarefas sem término | Perfis, pausas de GC, bloqueios e cancelamento |
| CPU baixa e resposta lenta | Espera de I/O, lock ou pool | Tempo nas dependências, pool de conexões e waits |
| Banco lento | Query cara, lock ou plano ruim | Slow queries, EXPLAIN, índices e estatísticas |
| Um provedor concentra o tempo | Dependência degradada | Tempo de conexão, timeout, erros e circuit breaker |

P95 de 800 ms significa que cerca de 95% das observações daquele conjunto ficaram em até 800 ms. A média pode esconder uma cauda lenta. Separe sucesso e erro quando isso mudar a interpretação.

Aumento de tráfego é hipótese, não conclusão. Mesmo sem aumento, pode haver mistura de requisições mais cara, dado específico, contenção ou dependência instável. Logs ajudam a contextualizar; métricas mostram dimensão e traces ajudam a localizar o tempo.

## “Como você avalia a saúde de uma aplicação?”

> Acompanho latência, tráfego, erros e saturação, considerando os objetivos do serviço. Observo também dependências, filas e resultados de negócio. Um processo responder health check não garante que consiga atender uma operação real.

- **Liveness:** o processo está vivo e deve continuar rodando?
- **Readiness:** esta instância pode receber tráfego agora?
- **SLO:** o objetivo de disponibilidade ou latência está sendo cumprido na janela definida?

A [API de exemplo](../../cmd/api/main.go) expõe apenas `/healthz` de liveness. Não há readiness de banco implementada neste laboratório.

## “Um bug acontece só em produção. E agora?”

> Procuro padrões por versão, endpoint, payload, horário e dependência. Correlaciono logs e traces usando IDs de requisição, sem expor dados sensíveis. Comparo configuração, volume, concorrência e características dos dados com os ambientes de teste. Tento reproduzir com um caso mínimo e dados anonimizados. Após corrigir, adiciono uma regressão útil e valido a entrega gradualmente.

Não reproduza escritas reais ou cargas perigosas em produção sem um plano específico. Flaky tests e bugs de concorrência podem exigir `go test -race`, repetição direcionada e observação de agendamento; o detector só alcança os caminhos executados.

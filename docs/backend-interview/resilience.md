# Dependência falhou: como reagir?

## Resposta de entrevista

> Defino timeouts e propago cancelamento para não ocupar recursos indefinidamente. Classifico a falha: para erros transitórios e operações seguras de repetir, uso retry limitado com backoff e jitter. Se a dependência falha continuamente, um circuit breaker interrompe chamadas temporariamente e testa recuperação. Uso fallback somente quando existe uma alternativa válida para o negócio. Acompanho erros, latência e saturação para verificar o efeito.

| Técnica | Papel | Exemplo do projeto |
| --- | --- | --- |
| Timeout/context | Limitar tempo e sinalizar cancelamento | [Context](../../fundamentals/10-context/main.go) |
| Retry | Tentar novamente uma falha transitória | [Retry](../../resilience/retry.go) |
| Backoff + jitter | Espaçar tentativas e reduzir sincronização entre clientes | [Retry](../../resilience/retry.go) |
| Circuit breaker | Evitar insistência contra dependência em falha | [Breaker](../../resilience/breaker.go) |
| Fallback | Fornecer alternativa aceitável | Exemplo conceitual: dado de cache com idade visível |
| Bulkhead | Limitar recursos consumidos por uma dependência | Próximo módulo; ainda não implementado |

Circuit breaker lembra um disjuntor: **fechado** permite chamadas; **aberto** rejeita; **half-open** testa recuperação com poucas chamadas. O exemplo permite uma sonda após cooldown.

Fallback não é callback. Callback é uma função executada em determinado ponto do programa; pode existir com ou sem falha.

Para uma cotação financeira, não use silenciosamente um preço desatualizado só para esconder erro. A alternativa precisa respeitar validade, risco e regras do produto. Em alguns casos, retornar indisponibilidade explícita é a decisão correta.

Retry não deve repetir falhas de validação e pode duplicar uma cobrança se não houver idempotência. Retentativas em várias camadas também multiplicam a carga. Cada chamada e o conjunto das tentativas precisam respeitar o orçamento total de tempo.

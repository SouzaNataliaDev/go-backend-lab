# API Go, contratos e graceful shutdown

## “Como você implementaria uma API em Go?”

> Começo pelo contrato, pelas regras de negócio e pelos requisitos de operação. Posso separar handler/controller, service e repository: o handler valida o transporte HTTP, o service aplica as regras, e o repository encapsula persistência. Defino interfaces pequenas no consumidor, propago context e erros, configuro timeouts e adiciono testes nas regras e integrações relevantes. Completo com logs, métricas e shutdown.

| Camada | Responsabilidade | Evitar |
| --- | --- | --- |
| Handler/controller | Interpretar HTTP, validar formato, mapear status e resposta | SQL e regra de negócio misturados ao transporte |
| Service | Orquestrar regras, autorização de negócio e transações | Depender de http.ResponseWriter |
| Repository | Acesso a dados e erros de persistência | Decidir status HTTP |

Exemplo de contrato de serviço (trecho ilustrativo, não uma API completa):

```go
// O serviço depende do comportamento necessário, não de um banco específico.
type OrderRepository interface {
    Find(ctx context.Context, id string) (Order, error)
}

// O handler repassa r.Context() para que cancelamento alcance o banco.
// O repository usa QueryContext, e erros são preservados com %w.
```

Camadas são uma escolha para organizar responsabilidades. Um endpoint muito simples pode dispensar uma camada que apenas repassa chamadas. A [API executável](../../cmd/api/main.go) ilustra operação HTTP; o [Strategy](../../patterns/strategy/payment.go) e [interfaces](../../fundamentals/04-interfaces/main.go) ilustram dependências substituíveis.

## “Como evitar quebrar contratos existentes?”

> Faço mudanças compatíveis com consumidores atuais, verifico contratos e coordeno alterações incompatíveis. Evito remover campos, mudar tipos ou tornar campos novos obrigatórios sem migração. Uso testes de contrato, implantação gradual e observação dos consumidores. Para mudanças incompatíveis, versiono ou mantenho os contratos durante a transição.

Adicionar campo opcional pode ser compatível se os consumidores tolerarem campos desconhecidos. Mudar de `amount_cents` inteiro para `amount` decimal, mudar unidade ou significado pode quebrar mesmo quando o JSON ainda é válido.

Para banco, prefira uma migração expand/contract: adicionar a nova estrutura, entregar código compatível com ambas, migrar dados/consumidores e remover a antiga depois. O rollback do código precisa continuar funcionando com o esquema implantado.

**Fallback não é compatibilidade retroativa:** é uma resposta alternativa a uma falha. Feature flags ajudam a controlar exposição, mas não corrigem um contrato incompatível por si só.

## “O que é graceful shutdown?”

> É encerrar de forma controlada: retirar a instância do recebimento de trabalho, parar de aceitar novas requisições e dar prazo para as que estão em andamento terminarem. Depois fecho recursos. Em consumidores, coordeno parar novas leituras, finalizar ou cancelar o trabalho e confirmar apenas o que foi concluído.

No Go, `http.Server.Shutdown(ctx)` fecha listeners e aguarda conexões HTTP ativas ficarem ociosas dentro do contexto. Se exceder o prazo, é preciso decidir como interromper. WebSockets/conexões hijacked e tarefas em background precisam de coordenação própria. Em ambiente com load balancer, readiness e drenagem de tráfego também fazem parte do fluxo.

Veja o [servidor com sinal SIGTERM e prazo de shutdown](../../cmd/api/main.go).

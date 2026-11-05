# Fundamentos de Go — prioridade de estudo

Exemplos executáveis com código comentado em português, aplicação prática, saída de referência, armadilhas e exercício. Os exemplos usam somente a biblioteca padrão e são compatíveis com o Go 1.22 declarado pelo projeto.

Em Go, comentários começam com `//` ou usam `/* ... */`. `\\` não é sintaxe de comentário; usamos `//` para que você possa copiar e executar os códigos.

## Trilha

| Ordem | Tema | Uso |
| --- | --- | --- |
| 1 | [Ponteiros](01-pointers/README.md) | Alterar um objeto recebido por uma função e entender cópia de valores. |
| 2 | [Slices](02-slices/README.md) | Trabalhar com listas e evitar alterações acidentais em dados compartilhados. |
| 3 | [Maps](03-maps/README.md) | Buscar valores por chave, distinguir ausência de zero e ordenar a exibição. |
| 4 | [Interfaces](04-interfaces/README.md) | Trocar uma dependência e criar um fake simples para testes. |
| 5 | [Tratamento de erros](05-errors/README.md) | Preservar a causa de uma falha e escolher uma resposta sem comparar textos. |
| 6 | [defer](06-defer/README.md) | Agendar limpeza e entender ordem de execução e captura de argumentos. |
| 7 | [Goroutines](07-goroutines/README.md) | Executar tarefas independentes e aguardar sua conclusão com WaitGroup. |
| 8 | [Channels](08-channels/README.md) | Comunicar resultados entre goroutines com controle de quem envia e fecha. |
| 9 | [select](09-select/README.md) | Escolher entre comunicação e cancelamento, evitando bloqueio indefinido. |
| 10 | [context.Context](10-context/README.md) | Propagar prazo e cancelamento de uma requisição às suas dependências. |
| 11 | [Concorrência e mutex](11-concurrency/README.md) | Proteger um estado compartilhado e entender por que WaitGroup sozinho não basta. |
| 12 | [Integração: fan-out com cancelamento](12-integration/README.md) | Combinar goroutines, channels, select, context, erros e defer em uma consulta de preços. |

Estude primeiro os fundamentos de valores e erros (1–6); depois siga a sequência de concorrência (7–12). Se estiver revisando para entrevista, priorize goroutines → channels → select → context → concorrência → integração, retomando os anteriores conforme precisar.

## Como estudar cada exemplo

1. Leia o README do tema e execute o comando indicado na raiz do projeto.
2. Abra o `main.go`: os comentários explicam o motivo de cada operação.
3. Compare a saída com a referência. Onde houver concorrência, confira os avisos sobre ordem variável.
4. Faça o exercício e rode os testes para observar os efeitos.

## Executar e verificar

```sh
go run ./fundamentals/07-goroutines
go run ./fundamentals/08-channels
go run ./fundamentals/09-select
go run ./fundamentals/10-context
go run -race ./fundamentals/11-concurrency
go run ./fundamentals/12-integration
go test -race ./fundamentals/...
```

`make fundamentals` executa os 12 exemplos; `make test` verifica o projeto inteiro. Os exemplos não precisam de APIs externas, credenciais ou banco de dados.

## Perguntas de revisão

- Quem cria, envia, recebe e fecha cada canal? O que acontece se ele for nil ou estiver fechado?
- Como a goroutine termina se o cliente cancelar a requisição?
- Qual memória é compartilhada ao copiar ponteiros, slices ou maps?
- O WaitGroup protege o contador ou só espera as tarefas?
- Como `%w`, `errors.Is` e `errors.As` preservam informações sobre falhas?
- Quando cada defer executa e quando os argumentos são avaliados?
- O que muda ao passar um valor concreto ou um ponteiro para uma interface?

## Fontes oficiais

- [Especificação de Go](https://go.dev/ref/spec)
- [Tour of Go](https://go.dev/tour/)
- [context](https://pkg.go.dev/context)
- [sync](https://pkg.go.dev/sync)
- [errors](https://pkg.go.dev/errors)
- [Detector de data race](https://go.dev/doc/articles/race_detector)

## Próximo passo

Depois desta trilha, retome os exemplos de [padrões de projeto](../patterns/strategy/payment.go), [resiliência](../resilience/) e [worker pool](../scalability/workers/). A integração de cotações ajuda a conectar os fundamentos com os desafios de backend.

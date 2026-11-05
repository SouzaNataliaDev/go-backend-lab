# context.Context

## Quando usar

Propagar prazo e cancelamento de uma requisição às suas dependências.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/10-context
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Prazo: context deadline exceeded
Cancelamento herdado: context canceled
Sucesso: cotação encontrada Erro: <nil>
```

A saída corresponde à simulação apresentada no código.

## Erro comum

cancel() deve ser chamado. Prazos precisam chegar ao cliente HTTP/banco, e cada tarefa deve observar o cancelamento.

## Pratique

Crie um contexto de 10 ms e chame lookup com uma operação de 1 hora; espere DeadlineExceeded.

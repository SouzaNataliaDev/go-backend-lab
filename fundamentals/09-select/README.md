# select

## Quando usar

Escolher entre comunicação e cancelamento, evitando bloqueio indefinido.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/09-select
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Resultado: cotação recebida Erro: <nil>
Cancelado: context canceled
Nenhum resultado disponível agora
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Receber de canal fechado está sempre pronto. Em um select repetido, trate ok=false para não girar sem parar.

## Pratique

Use context.WithTimeout para interromper a espera de um canal sem produtor.

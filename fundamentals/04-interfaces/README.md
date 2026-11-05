# Interfaces

## Quando usar

Trocar uma dependência e criar um fake simples para testes.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/04-interfaces
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Notificação: Pedido aprovado
Fake registrou: [Pedido aprovado]
Ponteiro nil: true
Interface nil: false
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Uma interface com ponteiro tipado nil não é igual a nil. Não crie interfaces grandes antes de ter uma necessidade concreta.

## Pratique

Implemente outro Notifier e use notifyOrder sem alterar seu código.

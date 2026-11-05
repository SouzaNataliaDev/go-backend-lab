# Slices

## Quando usar

Trabalhar com listas e evitar alterações acidentais em dados compartilhados.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/02-slices
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Compartilhado: [99 20 30]
Append com capacidade: [99 20 40]
Cópia: [1 20 40] Original: [99 20 40]
Capacidade limitada: [99 20 50] Original: [99 20 40]
Lista: [Go]
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Passar um slice copia seu descritor, não todos os elementos. Alterar elementos pode afetar quem chamou.

## Pratique

Crie uma função que retorne uma cópia de um slice sem compartilhar seu array.

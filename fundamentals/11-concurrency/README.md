# Concorrência e mutex

## Quando usar

Proteger um estado compartilhado e entender por que WaitGroup sozinho não basta.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/11-concurrency
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Total: 10000
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Deadlock e data race são problemas diferentes. O detector de race ajuda nas execuções observadas, mas não prova ausência de bugs.

## Pratique

Implemente Decrement e teste incrementos/decrementos concorrentes.

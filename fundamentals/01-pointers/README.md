# Ponteiros

## Quando usar

Alterar um objeto recebido por uma função e entender cópia de valores.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/01-pointers
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Após cópia: 10000
Após ponteiro: 15000
Dono: Natália
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Um ponteiro nil não aponta para um objeto; acessá-lo causa panic. Usar ponteiros em tudo não garante melhor desempenho.

## Pratique

Adicione um método com receiver *Account para sacar centavos.

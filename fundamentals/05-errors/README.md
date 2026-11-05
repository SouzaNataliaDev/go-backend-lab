# Tratamento de erros

## Quando usar

Preservar a causa de uma falha e escolher uma resposta sem comparar textos.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/05-errors
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
OK: pedido 42
Resposta 404: pedido não encontrado
Resposta 400: corrigir id
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Use %w para preservar a causa. Comparar err.Error() com uma string quebra quando o contexto da mensagem muda.

## Pratique

Adicione um erro de conflito e trate-o com errors.Is.

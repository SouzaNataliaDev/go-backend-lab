# Integração: fan-out com cancelamento

## Quando usar

Combinar goroutines, channels, select, context, erros e defer em uma consulta de preços.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/12-integration
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Melhor cotação: B, 9900 centavos
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Um buffer não resolve todos os vazamentos. É preciso definir como cada tarefa termina e quem fecha os canais.

## Pratique

Adicione um provedor que retorne erro e confira que as outras cotações continuam chegando.

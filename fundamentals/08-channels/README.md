# Channels

## Quando usar

Comunicar resultados entre goroutines com controle de quem envia e fecha.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/08-channels
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Soma recebida: 6
Primeiro: pedido A
Segundo: pedido B
Fechado e vazio: valor="" ok=false
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Quem envia coordena o fechamento. Em múltiplos produtores, um coordenador deve esperar todos antes de fechar.

## Pratique

Transforme os IDs em strings em uma etapa intermediária do pipeline.

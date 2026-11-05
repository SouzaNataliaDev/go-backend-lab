# Maps

## Quando usar

Buscar valores por chave, distinguir ausência de zero e ordenar a exibição.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/03-maps
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
PETG: 0 existe: true
ABS: 0 existe: false
PLA=10
TPU=5
Original após alias: 9
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Maps comuns não permitem leitura e escrita concorrentes sem sincronização. Veja o exemplo de concorrência.

## Pratique

Conte quantas vezes cada palavra aparece em um slice de strings.
